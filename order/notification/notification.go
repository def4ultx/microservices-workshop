package notification

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

type Client struct {
	kafkaWriter *kafka.Writer
}

func NewClient() *Client {
	config := kafka.WriterConfig{
		Brokers:   []string{"kafka:9092"},
		Topic:     "notification",
		Balancer:  &kafka.Hash{},
		BatchSize: 1,
		Async:     true,
	}
	writer := kafka.NewWriter(config)

	return &Client{
		kafkaWriter: writer,
	}
}

type EmailMessage struct {
	Channel string `json:"channel"`
	Address string `json:"address"`
	Message string `json:"message"`
}

func (c *Client) SendEmail(addr, message string) error {
	ctx := context.Background()

	email := EmailMessage{
		Channel: "email",
		Address: addr,
		Message: message,
	}
	text, err := json.Marshal(&email)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Key:   nil,
		Value: text,
	}
	err = c.kafkaWriter.WriteMessages(ctx, msg)
	if err != nil {
		return err
	}

	return nil
}
