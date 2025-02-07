package main

import (
    "context"
    "github.com/omniful/go_commons/kafka"
    "github.com/omniful/go_commons/pubsub"
)

func main() {
    // Initialize producer with configuration
    producer := kafka.NewProducer(
        kafka.WithBrokers([]string{"localhost:9092"}),
        kafka.WithClientID("my-producer"),
        kafka.WithKafkaVersion("2.8.1"),
    )
    defer producer.Close()

    // Create message with key for FIFO ordering
    msg := &pubsub.Message{
        Topic: "my-topic",
        // Key is crucial for maintaining FIFO ordering
        // Messages with the same key will be delivered to the same partition in order
        Key: "customer-123",  
        Value: []byte("Hello Kafka!"),
        Headers: map[string]string{
            "custom-header": "value",
            // Note: HeaderXOmnifulRequestID will be automatically added
            // from context if present
        },
    }

    // Context with request ID
    ctx := context.WithValue(context.Background(), "request_id", "req-123")
    
    // Synchronous publish - HeaderXOmnifulRequestID will be automatically added
    err := producer.Publish(ctx, msg)
    if err != nil {
        panic(err)
    }

    // Batch publish with consistent keys for ordering
    messages := []*pubsub.Message{
        {
            Topic: "my-topic",
            Key: "customer-123",  // Same key maintains ordering
            Value: []byte("Message 1"),
        },
        {
            Topic: "my-topic",
            Key: "customer-123",  // Same key maintains ordering
            Value: []byte("Message 2"),
        },
    }
    err = producer.PublishBatch(ctx, messages)
    if err != nil {
        panic(err)
    }
}