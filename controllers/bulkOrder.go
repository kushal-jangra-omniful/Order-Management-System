package controllers

import (
	"fmt"
	"net/http"
	"oms/producer"
	"oms/consumer"
	"oms/csvparse"


	"github.com/gin-gonic/gin"
)

// BulkOrder handles bulk order processing via API
func BulkOrder(c *gin.Context) {
	fmt.Println("Starting bulk order processing...")

	// Call the Csvinit function
	err := csvparse.Csvinit()
	orderPath := "./controllers/csvfile.csv"
    fmt.Println("msg is calling",producer.PublishOrderMessage(orderPath))
    consumer.StartConsumer()
	

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error processing bulk orders: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bulk order processing completed successfully."})
}
