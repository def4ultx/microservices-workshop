package order

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	InsertOrder(order Order) (string, error)
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
	return "", nil
}
