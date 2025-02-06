package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func InitMongoDB() error {
	uri := "mongodb://localhost:27017" 

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("MongoDB connection error: %v", err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("MongoDB ping error: %v", err)
		return err
	}

	fmt.Println("Connected to MongoDB successfully!")
	MongoClient = client
	return nil
}


func GetCollection(collectionName string) *mongo.Collection {
	if MongoClient == nil {
		log.Println("MongoClient is not initialized!")
		return nil
	}

	return MongoClient.Database("oms_db").Collection(collectionName)
}
