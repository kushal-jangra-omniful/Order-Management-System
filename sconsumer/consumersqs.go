package consumer

import (
	"context"
	// "fmt"
	"log"
	"oms/csvparse"
	"oms/utils"

	"github.com/omniful/go_commons/sqs"
)

// OrderMessageHandler processes incoming SQS messages.
type OrderMessageHandler struct{}

// Process handles the messages from SQS.
func (h *OrderMessageHandler) Process(ctx context.Context, messages *[]sqs.Message) error {
	for _, message := range *messages {
		messageString := string(message.Value)
		log.Printf("[DEBUG] Received message: %s", messageString)
		filepath:=messageString

		if err := csvparse.Csvinit(filepath); err != nil {
			log.Printf("[ERROR] Error processing order: %v", err)
			return err // Return error to possibly retry processing
		}
		log.Println("[SUCCESS] Order processed successfully.")
	}
	return nil
}

// StartConsumer initializes and runs the SQS consumer.
func StartConsumer() {
	if utils.SQSqueue == nil {
		log.Fatal("[FATAL] SQS queue is not initialized. Ensure SQSInitialization() is called first.")
	}

	consumer, err := sqs.NewConsumer(utils.SQSqueue, 1, 2, &OrderMessageHandler{}, 1, 10, true, false)
	if err != nil {
		log.Fatalf("[FATAL] Failed to create SQS consumer: %v", err)
	}

	log.Println("[INFO] Starting SQS Consumer for MyOrders.fifo...")
	consumer.Start(context.Background())

	// Keep the consumer running indefinitely
	select {}
}
