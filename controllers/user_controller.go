package controllers

import (
	"academ_be/configs"
	"academ_be/models"
	"academ_be/respones"
	"context"
	"log"
	"strings"

	"net/http"
	"time"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "User")
var firebaseClient *auth.Client = configs.ConnectFirebase()
var validate = validator.New()

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
		tokenString := strings.Replace(idToken, "Bearer ", "", 1)

		credential, err := firebaseClient.VerifyIDToken(ctx, tokenString)
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

		// Retrieves the first matching document
		var result models.UserResponse
		err = userCollection.FindOne(ctx, bson.M{"_id": credential.UID}).Decode(&result)
		if err != nil {
			log.Printf("Failed to verify ID token: %v", err)
			response := respones.UserResponse{
				Status:      http.StatusUnauthorized,
				Message:     ERROR,
				Description: "error firebase find one",
				Data:        nil,
			}
			c.JSON(http.StatusUnauthorized, response)
			return
		}

		response := respones.UserResponse{
			Status:      http.StatusOK,
			Message:     SUCCESS,
			Description: USER_SIGNIN_SUCCESS,
			Data:        result,
		}
		c.JSON(http.StatusCreated, response)

	}
}

func SignUpUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()

		// validate the request body
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

		// Verify token by Firebase admin
		tokenString := strings.Replace(idToken, "Bearer ", "", 1)
		credential, err := firebaseClient.VerifyIDToken(ctx, tokenString)
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

		// map newUser on save on database
		newUser := models.User{
			Id:       credential.UID,
			Email:    user.Email,
			FullName: user.FullName,
		}

		_, err = userCollection.InsertOne(ctx, newUser)
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

		// Map response succss and sending client
		response := respones.UserResponse{
			Status:      http.StatusCreated,
			Message:     SUCCESS,
			Description: USER_SIGNUP_SUCCESS,
			Data:        nil,
		}
		c.JSON(http.StatusCreated, response)
	}
}
