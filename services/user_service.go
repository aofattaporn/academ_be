package services

import (
	"academ_be/configs"
	"academ_be/models"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func init() {
	mongoClient = configs.ConnectDB()
}

func CreateUser(c *gin.Context, newUser *models.User) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	_, err = configs.GetCollection(mongoClient, "User").InsertOne(ctx, newUser)
	if err != nil {
		return err
	}

	return nil
}

func FindUserOneById(c *gin.Context, userID string) (user *models.UserResponse, err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	var result models.UserResponse
	err = configs.GetCollection(mongoClient, "User").FindOne(ctx, bson.M{"_id": userID}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func FindUserAndCount(c *gin.Context, userID string) (count int64, err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	count, err = configs.GetCollection(mongoClient, "User").CountDocuments(ctx, bson.M{"_id": userID})
	if err != nil {
		return 0, err
	}

	return count, nil
}
