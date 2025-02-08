package main

import (
	"fmt"
	
	"oms/sconsumer"
	"oms/interservice"

	"oms/routes"
	"time"

	"oms/kafka"
	// "oms/controllers"
	"oms/utils"

	"github.com/omniful/go_commons/http"
)

func main() {
   
	fmt.Println("Hello World!")

	// Initialize the server
	server := http.InitializeServer(
		":8081",        // listen address
		10*time.Second, // read timeout
		10*time.Second, // write timeout
		70*time.Second, // idle timeout
	)
	err := utils.InitMongoDB()
	if err != nil {
		fmt.Println("Error in connecting to mongo")
		return
	}
    // initialize sqs 
	utils.Initsqs()
    // iniatialize interservice
	interservice.InitInterSrvClient()
	// consumer sqs
	go consumer.StartConsumer()
    
    kafka.InitializeKafkaProducer()
	go kafka.StartConsumerk()
	// initialize redis client
	utils.InitRedis()

	// routes
	
	routes.RegisterRoutes(server)

	

	// Start the server
	if err := server.StartServer("MyService"); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}

	fmt.Println("Server started successfully")
}
