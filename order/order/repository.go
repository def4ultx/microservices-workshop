package order

import (
	"context"
	"order/inventory"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type Repository interface {
	InsertOrder(Order) (string, error)
	GetOrderByID(string) (*Order, error)
	GetOrdersByUserID(int) ([]Order, error)
}

type repository struct {
	mongoClient *mongo.Client
}

func NewRepository(client *mongo.Client) *repository {
	return &repository{
		mongoClient: client,
	}
}

func (r *repository) InsertOrder(order Order) (string, error) {
	ctx := context.Background()

	products := make([]bson.M, len(order.Products))
	for i, v := range order.Products {
		products[i] = bson.M{
			"id":     v.ID,
			"name":   v.Name,
			"price":  v.Price,
			"amount": v.Amount,
		}
	}

	orderID := uuid.New().String()
	doc := bson.M{
		"orderId":     orderID,
		"status":      order.Status,
		"userId":      order.UserID,
		"paymentId":   order.PaymentID,
		"totalAmount": order.TotalAmount,
		"products":    products,
	}

	_, err := r.mongoClient.Database("workshop").Collection("orders").InsertOne(ctx, doc)
	if err != nil {
		return "", err
	}

	return orderID, nil
}

func (r *repository) GetOrderByID(id string) (*Order, error) {
	ctx := context.Background()

	filter := bson.M{"orderId": id}
	result := r.mongoClient.Database("workshop").Collection("orders").FindOne(ctx, filter)
	order := struct {
		OrderID     string `bson:"orderId"`
		Status      string `bson:"status"`
		UserID      int    `bson:"userId"`
		PaymentID   int    `bson:"paymentId"`
		TotalAmount int    `bson:"totalAmount"`
		Products    []struct {
			ID     int64  `bson:"id"`
			Name   string `bson:"name"`
			Price  int    `bson:"price"`
			Amount int    `bson:"amount"`
		} `bson:"products"`
	}{}
	err := result.Decode(&order)
	if err != nil {
		return nil, err
	}

	products := make([]inventory.Product, len(order.Products))
	for i, v := range order.Products {
		products[i] = inventory.Product{
			ID:     v.ID,
			Name:   v.Name,
			Price:  v.Price,
			Amount: v.Amount,
		}
	}

	o := Order{
		OrderID:     order.OrderID,
		Status:      order.Status,
		UserID:      order.UserID,
		PaymentID:   order.PaymentID,
		TotalAmount: order.TotalAmount,
		Products:    products,
	}
	return &o, nil
}

func (r *repository) GetOrdersByUserID(id int) ([]Order, error) {
	ctx := context.Background()

	filter := bson.M{"userId": id}
	cursor, err := r.mongoClient.Database("workshop").Collection("orders").Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	orders := make([]Order, 0)
	for cursor.Next(ctx) {
		order := struct {
			OrderID     string `bson:"orderId"`
			Status      string `bson:"status"`
			UserID      int    `bson:"userId"`
			PaymentID   int    `bson:"paymentId"`
			TotalAmount int    `bson:"totalAmount"`
			Products    []struct {
				ID     int64  `bson:"id"`
				Name   string `bson:"name"`
				Price  int    `bson:"price"`
				Amount int    `bson:"amount"`
			} `bson:"products"`
		}{}
		if err = cursor.Decode(&order); err != nil {
			return nil, err
		}

		products := make([]inventory.Product, len(order.Products))
		for i, v := range order.Products {
			products[i] = inventory.Product{
				ID:     v.ID,
				Name:   v.Name,
				Price:  v.Price,
				Amount: v.Amount,
			}
		}

		o := Order{
			OrderID:     order.OrderID,
			Status:      order.Status,
			UserID:      order.UserID,
			PaymentID:   order.PaymentID,
			TotalAmount: order.TotalAmount,
			Products:    products,
		}
		orders = append(orders, o)
	}
	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}
