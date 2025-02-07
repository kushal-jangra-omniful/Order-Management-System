package producer

import (
	"context"
	"fmt"
	"log"

	"oms/utils"

	"github.com/google/uuid"
	"github.com/omniful/go_commons/sqs"
)

func PublishOrderMessage(orderpath string) error {
	// utils.Init()
	if utils.SQSpublisher == nil {
		return fmt.Errorf("SQS publisher is not initialized")
	}
	fmt.Println("Publishing order message")
    // filepath:="./controllers/csvfile.csv"
	message := &sqs.Message{
		Value:           []byte(orderpath),
		GroupId:         "order-processing-group",
		DeduplicationId: uuid.New().String(), 
	}

	err := utils.SQSpublisher.Publish(context.Background(), message)
	if err != nil {
		log.Printf("Failed to publish order message: %v", err)
		return err
	}

	fmt.Println("Successfully published message:", orderpath)
	return nil
}