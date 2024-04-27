package configs

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"firebase.google.com/go/messaging"
	"github.com/gin-gonic/gin"
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

func SendPushNotification(c *gin.Context) {
	// Connect to Firebase
	opt := option.WithCredentialsFile("./academprojex-firebase-adminsdk.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Notification Test",
			Body:  "Hello React!!",
		},
		Token: "fWL1hQobmZBg0DzrlABGe7:APA91bH_t6Vv7JkZFWhvfgx4ccH_RlzcYNYzPWv6UJ1z8dgP9f1Kcdi3MXywcY1ocqCpN3TjTOjuX52cPSgPTqilo6wI3LXLnPQQ9vp2rHkZqeUqMm2PNy7PondvkVpI6DgX_I6Dlvvm",
	}

	response, err := client.Send(ctx, message)
	if err != nil {
		// log.Fatalf("Error sending message: %v", err)
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Successfully sent message:", response)
	fmt.Println("message:", message)

}

func CreateMessage(c *gin.Context) (client *messaging.Client) {

	// Connect to Firebase
	opt := option.WithCredentialsFile("./academprojex-firebase-adminsdk.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	ctx := context.Background()
	client, err = app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	return client

}
