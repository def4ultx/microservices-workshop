package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Starting the service")
	StartServer(nil)
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
