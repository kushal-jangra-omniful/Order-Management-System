package repo

import (
    "context"
    "fmt"
    "log"
    "oms/utils"
)


func InsertOrdersIntoMongo(orders []interface{}) error {
    orderCollection := utils.GetCollection("orders")
    if orderCollection == nil {
        log.Fatal("MongoDB collection retrieval failed! Check MongoDB connection.")
    }

    _, err := orderCollection.InsertMany(context.Background(), orders)
    if err != nil {
        log.Printf("Error inserting orders into MongoDB: %v\n", err)
        return err
    }

    fmt.Println("Orders successfully inserted into MongoDB.")
    return nil
}