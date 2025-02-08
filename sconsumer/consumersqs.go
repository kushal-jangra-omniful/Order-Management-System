// package consumer

// import (
// 	"context"
// 	// "encoding/json"
// 	"fmt"
// 	"log"

// 	// "net/http"
// 	"oms/csvparse"
// 	"oms/models"

// 	"github.com/aws/aws-sdk-go-v2/aws"
// 	"github.com/aws/aws-sdk-go-v2/config"
// 	"github.com/aws/aws-sdk-go-v2/service/sqs"
// )

// const (
// 	queueURL = "https://sqs.eu-north-1.amazonaws.com/539247490249/Myqueue.fifo"
// )

// func StartConsumer() {
// 	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-north-1"))
// 	if err != nil {
// 		log.Fatalf("failed to load AWS SDK config: %v", err)
// 	}

// 	sqsClient := sqs.NewFromConfig(cfg)

// 	for {
// 		output, err := sqsClient.ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
// 			QueueUrl:            aws.String(queueURL),
// 			MaxNumberOfMessages: 1,
// 			WaitTimeSeconds:     1,
// 		})
// 		if err != nil {
// 			log.Printf("failed to receive messages: %v", err)
// 			continue
// 		}

// 		for _, msg := range output.Messages {
// 			messageOutput := models.MessageOutput{Message: *msg.Body}
// 			fmt.Printf("Received message: %s\n", messageOutput.Message)
// 			filePath := messageOutput.Message
// 			if err := csvparse.Csvinit(filePath); err != nil {
// 				fmt.Println("error in creating order", err)
// 			}

// 			// Delete message after processing
// 			_, err = sqsClient.DeleteMessage(context.TODO(), &sqs.DeleteMessageInput{
// 				QueueUrl:      aws.String(queueURL),
// 				ReceiptHandle: msg.ReceiptHandle,
// 			})
// 			if err != nil {
// 				log.Printf("failed to delete message: %v", err)
// 			} else {
// 				fmt.Println("Message deleted successfully")
// 			}

// 		}
// 	}
// }

// import (
// 	"context"
// 	"fmt"
// 	config "oms/configs"

// 	"github.com/omniful/go_commons/log"
// 	"github.com/omniful/go_commons/sqs"
// )

// type MyHandler struct{}

// func (h *MyHandler) Handle(msg sqs.Message) error {
// 	fmt.Println("Processing message:", string(msg.Value))
// 	return nil
// }

// func (h *MyHandler) Process(ctx context.Context, msgs *[]sqs.Message) error {
// 	for _, msg := range *msgs {
// 		if err := h.Handle(msg); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func InitSQS() {
// 	queueName := "OMSQueue.fifo"
// 	queue, err := sqs.NewFifoQueue(context.Background(), queueName, &sqs.Config{
// 		Account:  config.SQS_Config.Account,
// 		Endpoint: config.SQS_Config.Endpoint,
// 		Region:   config.SQS_Config.Region,
// 	})
// 	if err != nil || queue == nil {
// 		log.Errorf("initialization error. queue: %v, err : %v, publisher: %+v", queueName, err, queue)
// 		return
// 	}
// 	Handler := &MyHandler{}
// 	log.Debugf("creating queue: %v", queueName)
// 	consumer, err := sqs.NewConsumer(
// 		queue,
// 		uint64(1),
// 		4,
// 		Handler,
// 		10,
// 		30,
// 		true,
// 		false,
// 	)
// 	if err != nil || consumer == nil {
// 		log.Errorf("initialization error. queue: %v, err : %v, publisher: %+v", queueName, err, consumer)
// 		return
// 	}
// 	consumer.Start(context.Background())
// 	fmt.Println("queue created successfully")

// 	publisher := sqs.NewPublisher(queue)
// 	message := &sqs.Message{
// 		GroupId:         "group-1",
// 		Value:           []byte("Hello SQS!"),
// 		ReceiptHandle:   "group-1",
// 		DeduplicationId: "gp-1",
// 	}
// 	ctx := context.Background()
// 	if err := publisher.Publish(ctx, message); err != nil {
// 		log.Errorf("Failed to publish message: %v", err)
// 	}

// }


//	-----------------------------------------------------



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
