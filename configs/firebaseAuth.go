package configs

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func ConnectFirebase() *auth.Client {

	// Connect to Firebase
	opt := option.WithCredentialsFile("./academprojex-firebase-adminsdk.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	// Initialize Firebase Auth client
	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Error creating Firebase Auth client: %v\n", err)
	}

	return authClient
}
