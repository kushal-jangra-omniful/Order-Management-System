package csvparse

import (
	"context"
	"encoding/json"
	"fmt"

	// "log"
	// "net/http"
	// "oms/utils"

	// "oms/consumer"
	"oms/interservice"
	"oms/kafka"
	"oms/models"
	"oms/repo"
	"time"

	"github.com/omniful/go_commons/csv"
)

// Csvinit reads a CSV file, parses it, and inserts orders into MongoDB.
func Csvinit(filepath string) error {

	ctx := context.Background()
	// Initialize CommonCSV with options
	commonCSV, err := csv.NewCommonCSV(
		csv.WithBatchSize(1),
		csv.WithSource(csv.Local),
		csv.WithLocalFileInfo(filepath),
	)
	if err != nil {
		fmt.Printf("Error initializing CommonCSV: %v\n", err)
		return err
	}

	// Initialize the reader
	err = commonCSV.InitializeReader(context.Background())
	if err != nil {
		fmt.Printf("Error initializing reader: %v\n", err)
		return err
	}

	// Parse the CSV headers
	headers, err := commonCSV.GetHeaders()
	if err != nil {
		fmt.Printf("Error parsing headers: %v\n", err)
		return err
	}
	fmt.Println("Headers:", headers)

	// Read and process the CSV records
	records, _, err := commonCSV.ProcessSheet()
	if err != nil {
		fmt.Printf("Error processing sheet: %v\n", err)
		return err
	}

	// Convert CSV records into Order structs
	var orders []interface{} // Using interface{} for bulk insert
	for _, record := range records[1:] {
		if len(record) < 4 { // Ensure there are enough fields
			fmt.Printf("Skipping invalid record: %v\n", record)
			continue
		}
		sku := record[4] // Assuming SKU is in the 4th column

		// Verify SKU in WMS
		exists, err := interservice.VerifySKU(ctx, sku)
		if err != nil {
			fmt.Printf("Error verifying SKU %s: %v\n", sku, err)
			continue
		}
		if !exists {
			fmt.Printf("SKU do not exist")
			continue
		}

		order := models.Order{
			ID:        record[0],
			TenantID:  record[1],
			SellerID:  record[2],
			HubID:     record[3],
			SKU:       record[4],
			Quantity:  1,         // Default quantity (modify if needed)
			Status:    "on_hold", // Default status
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		bytesOrderItem, err := json.Marshal(order.Quantity)
		if err != nil {
			fmt.Println("Error marshaling order quantity:", err)

		}
		kafka.PublishMessageToKafka(bytesOrderItem, order.ID)

		orders = append(orders, order)
	}

	fmt.Println("Parsed Orders:", orders)

	// Insert orders into MongoDB
	repo.InsertOrdersIntoMongo(orders)

	// fmt.Println("Orders successfully inserted into MongoDB.")
	return nil
}
