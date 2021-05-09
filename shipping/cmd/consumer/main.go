package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	// make a new reader that consumes from topic-A
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"shipping-order-kafka:9092"},
		GroupID:  "consumer-group-1",
		Topic:    "shipping-order",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println("got err while reading message: ", err)
			break
		}

		err = handleMessage(m.Value)
		if err != nil {
			log.Println("got err while handle message: ", err)
		}

	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

type Product struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type KafkaMessage struct {
	Address   string    `json:"address"`
	Recipient string    `json:"recipient"`
	Products  []Product `json:"products"`
}

func handleMessage(str []byte) error {

	var msg KafkaMessage
	err := json.Unmarshal(str, &msg)
	if err != nil {
		return err
	}

	log.Println("shipping order for", msg.Recipient, "to", msg.Address)
	return nil
}
