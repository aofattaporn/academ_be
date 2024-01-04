package main

import (
	"academ_be/configs"
	"academ_be/routes"
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/api/option"
)

func main() {
	router := gin.Default()

	// Connect to Firebase
	opt := option.WithCredentialsFile("./academprojex-firebase-adminsdk-yew3p-4a85b7993d.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Errorf("error initializing app: %v", err)
	}

	// Initialize Firebase Auth client
	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Error creating Firebase Auth client: %v\n", err)
	}

	// Generate a UUID for the new user
	uid := uuid.New().String()

	// Create a custom token for the user using the Firebase Admin SDK
	customToken, err := authClient.CustomToken(context.Background(), uid)
	if err != nil {
		log.Printf("failed to create custom token for user: %v", err)
	}

	fmt.Println(customToken)
	fmt.Println("==================")
	//run database
	configs.ConnectDB()

	//routes
	routes.UserRoute(router) //add this

	router.Run("localhost:6000")
}
