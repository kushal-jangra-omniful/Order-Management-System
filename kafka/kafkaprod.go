package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	// "time"
    "oms/models"
	"github.com/omniful/go_commons/kafka"
	"github.com/omniful/go_commons/pubsub"
)

// Order represents the order structure to be sent to Kafka


// Initialize Kafka producer (Singleton)
var producer = kafka.NewProducer(
	kafka.WithBrokers([]string{"localhost:9092"}),
	kafka.WithClientID("oms-producer"),
	kafka.WithKafkaVersion("2.8.1"),
)

// PublishVerifiedOrder sends a verified order to Kafka
func PublishVerifiedOrder(order models.Order) error {
	defer producer.Close() // Ensure producer is closed when function exits

	// Convert order struct to JSON
	orderJSON, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("failed to serialize order: %v", err)
	}

	// Context with request ID
	ctx := context.WithValue(context.Background(), "request_id", fmt.Sprintf("order-%s", order.ID))

	// Kafka message with ordering key (e.g., order ID)
	msg := &pubsub.Message{
		Topic: "verified-orders",
		Key:   order.ID, // Ensures FIFO ordering per order
		Value: orderJSON,
	}

	// Publish message
	err = producer.Publish(ctx, msg)
	if err != nil {
		return fmt.Errorf("failed to publish verified order: %v", err)
	}

	fmt.Printf("Published verified order: %+v\n", order)
	return nil
}
