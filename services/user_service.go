package services

import (
	"academ_be/configs"
	"academ_be/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var mongoClient *mongo.Client

func init() {
	mongoClient = configs.ConnectDB()
}

func CreateUser(c *gin.Context, newUser *models.User) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	_, err := configs.GetCollection(mongoClient, "User").InsertOne(ctx, newUser)
	if err != nil {
		handleError(c, err, http.StatusUnauthorized, ERROR)
	}
}

func FindUserOneById(c *gin.Context, userID string) *models.UserResponse {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	var result models.UserResponse
	err := configs.GetCollection(mongoClient, "User").FindOne(ctx, bson.M{"_id": userID}).Decode(&result)
	if err != nil {
		handleError(c, err, http.StatusUnauthorized, ERROR)
	}

	return &result
}

func FindUserAndCount(c *gin.Context, userID string) int64 {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	count, err := configs.GetCollection(mongoClient, "User").CountDocuments(ctx, bson.M{"_id": userID})
	if err != nil {
		handleError(c, err, http.StatusUnauthorized, ERROR)
	}

	return count
}
