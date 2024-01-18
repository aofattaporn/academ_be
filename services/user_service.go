package services

import (
	"academ_be/configs"
	"academ_be/models"
	"academ_be/respones"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type UserService interface {
	CreateUser(user *models.User) error
	FindUserOneByID(c *gin.Context, userID string) (models.UserResponse, error)
}

type ConcreteUserService struct{}

func (c *ConcreteUserService) CreateUser(newUser *models.User) error {
	return nil
}

func CreateUser(c *gin.Context, newUser *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := configs.GetCollection(configs.DB, "User").InsertOne(ctx, newUser)
	if err != nil {
		response := respones.CustomResponse{
			Status:      http.StatusBadRequest,
			Message:     ERROR,
			Description: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return err
	}

	return nil
}

func FindUserOneById(c *gin.Context, userID string) error {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	var result models.UserResponse
	err := configs.GetCollection(configs.DB, "User").FindOne(ctx, bson.M{"_id": userID}).Decode(&result)
	if err != nil {
		log.Printf("Failed to verify ID token: %v", err)
		response := respones.CustomResponse{
			Status:      http.StatusUnauthorized,
			Message:     ERROR,
			Description: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
	}

	return nil
}
