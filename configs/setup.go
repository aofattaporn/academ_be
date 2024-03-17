package configs

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB creates a new MongoDB client and connects to the database
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

// GetCollection returns a MongoDB collection from the given client and collection name
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database(EnvMongoDatabase()).Collection(collectionName)
}
