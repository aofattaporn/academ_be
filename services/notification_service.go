package services

import (
	"academ_be/configs"
	"academ_be/models"
	"context"
	"errors"
	"fmt"
	"time"

	"firebase.google.com/go/messaging"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func FindFCMByMember(c *gin.Context, userId string) (fcm *models.FCM, err error) {

	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": userId}
	fmt.Println(userId)

	err = configs.GetCollection(mongoClient, USER_COLLECTION).FindOne(ctx, filter).Decode(&fcm)
	if err != nil {
		return nil, err
	}

	if fcm.FCM_TOKEN == "" {
		return nil, errors.New("not found fcm")
	}

	return fcm, nil

}

func AddNotification(c *gin.Context, fcmToken string, noti models.Notification) (err error) {

	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	_, err = configs.GetCollection(mongoClient, NOTIFICATION_COLLECTION).InsertOne(ctx, noti)
	if err != nil {
		return err
	}

	PushNotification(c, fcmToken, noti)

	return nil

}

func PushNotification(c *gin.Context, fcmToken string, noti models.Notification) {

	client := configs.CreateMessage(c)

	data := map[string]string{
		"ProjectName": noti.ProjectProfile.ProjectName,
		"AvatarColor": noti.ProjectProfile.AvatarColor,
		"Title":       noti.Title,
		"Body":        noti.Body,
		"Date":        noti.Date.String(),
	}

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: noti.Title,
			Body:  noti.Body,
		},
		Data:  data,
		Token: fcmToken,
		Webpush: &messaging.WebpushConfig{
			FcmOptions: &messaging.WebpushFcmOptions{
				Link: "https://localhost:5173/",
			},
		},
	}

	_, err := client.Send(c, message)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

}

func GetAllNotifications(c *gin.Context, userId string) (notifications []models.Notification, err error) {

	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	filter := bson.M{"userId": userId}

	cursor, err := configs.GetCollection(mongoClient, NOTIFICATION_COLLECTION).Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Decode projects directly into the result slice
	if err := cursor.All(ctx, &notifications); err != nil {
		return nil, err
	}

	if len(notifications) == 0 {
		return []models.Notification{}, nil
	}

	return notifications, nil

}
