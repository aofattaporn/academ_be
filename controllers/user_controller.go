package controllers

import (
	"academ_be/configs"
	"academ_be/models"
	"academ_be/respones"
	"context"
	"log"

	"net/http"
	"time"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "User")
var firebaseClient *auth.Client = configs.ConnectFirebase()
var validate = validator.New()

func SignUpUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&user); err != nil {
			response := respones.UserResponse{
				Status:      http.StatusBadRequest,
				Message:     "Error",
				Description: "Email or Password is null",
				Data:        err.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			response := respones.UserResponse{
				Status:      http.StatusBadRequest,
				Message:     "Error",
				Description: "Email or Password is null",
				Data:        validationErr.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		newUser := models.User{
			Id:       primitive.NewObjectID(),
			Email:    user.Email,
			Password: user.Password,
		}

		result, err := userCollection.InsertOne(ctx, newUser)
		if err != nil {
			response := respones.UserResponse{
				Status:      http.StatusBadRequest,
				Message:     "Error",
				Description: "Email or Password is null",
				Data:        err.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		// Generate a UUID for the new user
		uid := uuid.New().String()

		// Create a custom token for the user using the Firebase Admin SDK
		customToken, err := firebaseClient.CustomToken(context.Background(), uid)
		if err != nil {
			log.Printf("failed to create custom token for user: %v", err)
		}

		c.SetCookie("token", customToken, 0, "/", "", false, true)

		response := respones.UserResponse{
			Status:      http.StatusCreated,
			Message:     "Success",
			Description: "Create User Sccuess",
			Data:        result,
		}
		c.JSON(http.StatusCreated, response)

	}
}
