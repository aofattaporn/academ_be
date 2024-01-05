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

func Verify() gin.HandlerFunc {
	return func(c *gin.Context) {
		return
	}
}

func SignInUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Get ID token from the request header
		idToken := c.GetHeader(Authorization)
		if idToken == "" {
			response := respones.UserResponse{
				Status:      http.StatusUnauthorized,
				Message:     ERROR,
				Description: MISSING_AUTH_HEADER,
				Data:        nil,
			}
			c.JSON(http.StatusUnauthorized, response)
			return
		}

		// Verify the ID token
		_, err := firebaseClient.VerifyIDToken(ctx, idToken)
		if err != nil {
			log.Printf("Failed to verify ID token: %v", err)
			response := respones.UserResponse{
				Status:      http.StatusUnauthorized,
				Message:     ERROR,
				Description: INVALID_TOKEN,
				Data:        nil,
			}
			c.JSON(http.StatusUnauthorized, response)
			return
		}

		// Generate a UUID for the new user
		uid := uuid.New().String()

		// Create a custom token for the user using the Firebase Admin SDK
		customToken, err := firebaseClient.CustomToken(ctx, uid)
		if err != nil {
			log.Printf("Failed to create custom token for user: %v", err)
		}

		c.SetCookie(TOKEN, customToken, 0, "/", "", false, true)

		response := respones.UserResponse{
			Status:      http.StatusOK,
			Message:     SUCCESS,
			Description: USER_SIGNIN_SUCCESS,
			Data:        nil,
		}
		c.JSON(http.StatusCreated, response)

	}
}

func SignUpUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&user); err != nil {
			response := respones.UserResponse{
				Status:      http.StatusBadRequest,
				Message:     ERROR,
				Description: EMAIL_PASSWORD_NULL,
				Data:        err.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			response := respones.UserResponse{
				Status:      http.StatusBadRequest,
				Message:     ERROR,
				Description: INPUT_INVALID,
				Data:        validationErr.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		newUser := models.User{
			Id:       primitive.NewObjectID(),
			Email:    user.Email,
			FullName: user.FullName,
		}

		_, err := userCollection.InsertOne(ctx, newUser)
		if err != nil {
			response := respones.UserResponse{
				Status:      http.StatusBadRequest,
				Message:     ERROR,
				Description: MONGO_ERROR,
				Data:        err.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		response := respones.UserResponse{
			Status:      http.StatusCreated,
			Message:     SUCCESS,
			Description: USER_SIGNUP_SUCCESS,
			Data:        nil,
		}
		c.JSON(http.StatusCreated, response)

	}
}
