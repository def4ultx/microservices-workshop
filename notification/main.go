package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	// make a new reader that consumes from topic-A
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"kafka:9092"},
		GroupID:  "consumer-group-1",
		Topic:    "notification",
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

type KafkaMessage struct {
	Channel string `json:"channel"`
	Address string `json:"address"`
	Message string `json:"message"`
}

func handleMessage(str []byte) error {

	var msg KafkaMessage
	err := json.Unmarshal(str, &msg)
	if err != nil {
		return err
	}

	switch msg.Channel {
	case "email":
		return sendEmail(msg.Address, msg.Message)
	case "sms":
		return errors.New("sms not implemented")
	}
	return nil
}

func sendEmail(addr, text string) error {
	log.Println("send email to: ", addr)
	return nil
}
