package main

import (
	"fmt"
	"oms/controllers"
	"oms/routes"
	"time"

	// "github.com/omniful/go_commons/csv"
	// "fmt"
	// "context"
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
	utils.InitRedis()
	controllers.Csvinit()
	routes.RegisterRoutes(server)

	// Start the server
	if err := server.StartServer("MyService"); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}

	fmt.Println("Server started successfully")
}
