package services

import (
	"academ_be/configs"
	"academ_be/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "User")

type UserService interface {
	CreateUser(user *models.User) error
}

type ConcreteUserService struct{}

func (c *ConcreteUserService) CreateUser(newUser *models.User) error {
	return nil
}

// CreateResource creates a new resource
func CreateUser(newUser *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		fmt.Println(err)

		return err
	}

	fmt.Println(newUser)

	return nil
}
