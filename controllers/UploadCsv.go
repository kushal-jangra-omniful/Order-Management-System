package controllers

import (
	"context"
	"fmt"
	"log"
	"oms/utils"
	"oms/models"
	"time"

	"github.com/omniful/go_commons/csv"
)

// Csvinit reads a CSV file, parses it, and inserts orders into MongoDB.
func Csvinit() error {
	// Path to your local CSV file
	localFilePath := "./controllers/csvfile.csv"

	// Initialize CommonCSV with options
	commonCSV, err := csv.NewCommonCSV(
		csv.WithBatchSize(1),
		csv.WithSource(csv.Local),
		csv.WithLocalFileInfo(localFilePath),
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
	for _, record := range records {
		if len(record) < 4 { // Ensure there are enough fields
			fmt.Printf("Skipping invalid record: %v\n", record)
			continue
		}

		order := models.Order{
			// TenantID:  record[0],
			// SellerID:  record[1],
			// HubID:     record[2],
			ID:        record[0],
			SKU:       record[1],
			Quantity:  1,         // Default quantity (modify if needed)
			Status:    "on_hold", // Default status
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		orders = append(orders, order)
	}

	fmt.Println("Parsed Orders:", orders)

	// Insert orders into MongoDB

	orderCollection := utils.GetCollection("orders")
	if orderCollection == nil {
		log.Fatal("MongoDB collection retrieval failed! Check MongoDB connection.")
	}
	
	_, err = orderCollection.InsertMany(context.Background(), orders)
	if err != nil {
		log.Printf("Error inserting orders into MongoDB: %v\n", err)
		return err
	}

	fmt.Println("Orders successfully inserted into MongoDB.")
	return nil
}
