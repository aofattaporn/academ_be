package services

import (
	"academ_be/configs"
	"academ_be/models"
	"context"
	"fmt"
	"time"

	"firebase.google.com/go/messaging"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindFCMByMember(userId string) (fcm *models.FCM, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": userId}

	err = configs.GetCollection(mongoClient, USER_COLLECTION).FindOne(ctx, filter).Decode(&fcm)
	if err != nil {
		return nil, err
	}

	return fcm, nil

}

func AddNotification(fcmToken string, noti models.Notification) (err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = configs.GetCollection(mongoClient, NOTIFICATION_COLLECTION).InsertOne(ctx, noti)
	if err != nil {
		return err
	}

	PushNotification(fcmToken, noti)

	return nil

}

func PushNotificationCron(fcmToken string, noti models.Notification) {

	client := configs.CreateMessage()

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
				Link: "https://academ-fe.onrender.com/",
			},
		},
	}

	_, err := client.Send(context.Background(), message)
	if err != nil {
		return
	}

}

func PushNotification(fcmToken string, noti models.Notification) {

	client := configs.CreateMessage()

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

	_, err := client.Send(context.Background(), message)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}

func GetAllNotifications(c *gin.Context, userId string) (notifications []models.NotificationRes, err error) {

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
		return []models.NotificationRes{}, nil
	}

	return notifications, nil

}

func UpdateClearNotiById(c *gin.Context, notiId string, isClearNoti bool) (err error) {

	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	notiID, err := primitive.ObjectIDFromHex(notiId)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": notiID}
	update := bson.M{"$set": bson.M{"isClear": isClearNoti}}

	_, err = configs.GetCollection(mongoClient, NOTIFICATION_COLLECTION).UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil

}

func DeleteNotiById(c *gin.Context, notiId string) (err error) {

	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	notiID, err := primitive.ObjectIDFromHex(notiId)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": notiID}
	_, err = configs.GetCollection(mongoClient, NOTIFICATION_COLLECTION).DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil

}
