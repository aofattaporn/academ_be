package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB creates a new MongoDB client and connects to the database
func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatalf("Error creating MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	if err := pingMongoDB(ctx, client); err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB")
	return client
}

// GetCollection returns a MongoDB collection from the given client and collection name
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database(EnvMongoDatabase()).Collection(collectionName)
}

// pingMongoDB pings MongoDB to check the connection
func pingMongoDB(ctx context.Context, client *mongo.Client) error {
	return client.Ping(ctx, nil)
}
