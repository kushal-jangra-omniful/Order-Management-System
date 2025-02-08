package utils

import (
	"context"
	"fmt"

	"github.com/omniful/go_commons/sqs"
)

var(
	SQSqueue     *sqs.Queue
	SQSpublisher *sqs.Publisher
	SQSconsumer  *sqs.Consumer
)
func Initsqs() {
	config := sqs.Config{
		Account:  "539247490249",
		Endpoint: "https://sqs.eu-north-1.amazonaws.com/539247490249/Myqueue.fifo",
		Region:   "eu-north-1",
	}

	queueName := "Myqueue.fifo"
	queue, err := sqs.NewFifoQueue(context.Background(), queueName, &sqs.Config{
		Account:  config.Account,
		Endpoint: config.Endpoint,
		Region:   config.Region,
	})
	if err != nil || queue == nil {
		fmt.Printf("initialization error. queue: %v, err : %v, publisher: %+v", queueName, err, queue)
		return
	}
	SQSqueue = queue
	SQSpublisher=sqs.NewPublisher(queue)
	consumer, err := sqs.NewConsumer(queue, 1, 1, nil, 1, 30, true, false)
	if err != nil {
		fmt.Printf("consumer initialization error: %v", err)
		return
	}
	SQSconsumer = consumer

	fmt.Println("queue created successfully")

}
