package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"oms/models"

	"github.com/omniful/go_commons/kafka"
	"github.com/omniful/go_commons/pubsub"
	"github.com/omniful/go_commons/pubsub/interceptor"
)

// MessageHandler processes incoming messages
type MessageHandler struct{}

// Process processes each Kafka message
func (h *MessageHandler) Process(ctx context.Context, msg *pubsub.Message) error {
	// Deserialize message into Order struct
	var order models.Order
	if err := json.Unmarshal(msg.Value, &order); err != nil {
		log.Printf("❌ Error deserializing order: %v\n", err)
		return err
	}

	// Process order
	err := processOrder(order)
	if err != nil {
		log.Printf("⚠️ Error processing order ID %s: %v\n", order.ID, err)
		return err
	}

	log.Printf("✅ Successfully processed order: %+v\n", order)
	return nil
}

// processOrder contains business logic for processing the order
func processOrder(order models.Order) error {
	
	fmt.Printf("📦 Order processed: %+v\n", order)
	return nil
}

// StartConsumer initializes the Kafka consumer
func StartConsumerk() {
	// Initialize Kafka consumer
	consumer := kafka.NewConsumer(
		kafka.WithBrokers([]string{"localhost:9092"}),  // Replace with actual broker list
		kafka.WithConsumerGroup("oms-order-consumer"),
		kafka.WithClientID("oms-consumer"),
		kafka.WithKafkaVersion("2.8.1"),
		kafka.WithRetryInterval(time.Second),
	)

	defer consumer.Close()

	// Set monitoring interceptor
	consumer.SetInterceptor(interceptor.NewRelicInterceptor())

	// Register message handler for "verified-orders" topic
	handler := &MessageHandler{}
	consumer.RegisterHandler("verified-orders", handler)

	// Start consuming messages
	ctx := context.Background()
	log.Println("📥 Consumer started, listening to verified-orders topic...")
	consumer.Subscribe(ctx)
}
