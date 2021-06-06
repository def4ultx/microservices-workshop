package main

import (
	"context"
	"net/http"
	"order/handler"
	"order/inventory"
	"order/middleware"
	"order/notification"
	"order/payment"
	"order/shipping"
	"os"
	"os/signal"
	"time"

	fluent "github.com/evalphobia/logrus_fluent"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	log.Info("Starting the service")
	setupLogger()

	mongoClient := connectDB()
	defer disconnectMongo(mongoClient)

	inventoryClient := inventory.NewClient()
	paymentClient := payment.NewClient()
	shippingClient := shipping.NewClient()
	notificationClient := notification.NewClient()

	r := mux.NewRouter()
	r.Use(middleware.Logging, middleware.Metric, middleware.Recover)
	r.Handle("/prometheus", promhttp.Handler()).Methods(http.MethodGet)
	r.Handle("/healthz", http.HandlerFunc(handler.HealthCheck))

	o := handler.NewOrderHandler(inventoryClient, paymentClient, shippingClient, notificationClient, mongoClient)
	r.HandleFunc("/order", o.CreateOrder).Methods(http.MethodPost)
	r.HandleFunc("/order/{id}", o.GetOrderByID).Methods(http.MethodGet)
	r.HandleFunc("/orders/{userId}", o.GetUserOrders).Methods(http.MethodGet)

	StartServer(r)
}

func setupLogger() {
	log.SetFormatter(&log.TextFormatter{})
	hook, err := fluent.NewWithConfig(fluent.Config{
		Host: "fluentd",
		Port: 24224,
	})
	if err != nil {
		log.Error(err)
		return
	}
	hook.SetTag("original.tag")
	hook.SetMessageField("message")
	log.AddHook(hook)
}

func connectDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://order-mongodb:27017"))
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func disconnectMongo(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func StartServer(r http.Handler) {

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	log.Println("The service is ready to listen and serve.")
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// Wait for an interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	log.Info("Shutting down server")

	err := shutdownServer(srv)
	if err != nil {
		log.WithError(err).Error("failed to shutdown server")
	}
}

func shutdownServer(srv *http.Server) error {
	waitTime := 30 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), waitTime)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}
