package consumer

import (
	"context"
	"fmt"
	"log"
    "oms/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

const (
	queueURL = "https://sqs.eu-north-1.amazonaws.com/539247490249/Myqueue.fifo"
)



func StartConsumer() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-north-1"))
	if err != nil {
		log.Fatalf("failed to load AWS SDK config: %v", err)
	}

	sqsClient := sqs.NewFromConfig(cfg)

	for {
		output, err := sqsClient.ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(queueURL),
			MaxNumberOfMessages: 10,
			WaitTimeSeconds:     5,
		})
		if err != nil {
			log.Printf("failed to receive messages: %v", err)
			continue
		}

		for _, msg := range output.Messages {
			messageOutput := models.MessageOutput{Message: *msg.Body}
			fmt.Printf("Received message: %s\n", messageOutput.Message)

			// Delete message after processing
			_, err = sqsClient.DeleteMessage(context.TODO(), &sqs.DeleteMessageInput{
				QueueUrl:      aws.String(queueURL),
				ReceiptHandle: msg.ReceiptHandle,
			})
			if err != nil {
				log.Printf("failed to delete message: %v", err)
			} else {
				fmt.Println("Message deleted successfully")
			}
		}
	}
}