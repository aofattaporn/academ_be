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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "User")
var firebaseClient *auth.Client = configs.ConnectFirebase()
var validate = validator.New()

func SignInWithGoogle() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 10*time.Second)
		defer cancel()

		// get user data and finding on db
		userID := c.MustGet("userID").(string)

		count, err := userCollection.CountDocuments(ctx, bson.M{"_id": userID})
		if err != nil {
			log.Printf("Failed to Count Document: %v", err)
			response := respones.CustomResponse{
				Status:      http.StatusBadRequest,
				Message:     ERROR,
				Description: "error firebase count documents",
			}
			c.JSON(http.StatusUnauthorized, response)
			return
		}

		// checking user count on db
		if count > 0 {
			response := respones.CustomResponse{
				Status:      http.StatusOK,
				Message:     SUCCESS,
				Description: USER_SIGNIN_SUCCESS,
			}
			c.JSON(http.StatusCreated, response)

		} else {

			// map user and create newUser account on db
			var user models.User
			if err := c.BindJSON(&user); err != nil {
				response := respones.CustomResponse{
					Status:      http.StatusBadRequest,
					Message:     ERROR,
					Description: EMAIL_PASSWORD_NULL,
					Data:        err.Error(),
				}
				c.JSON(http.StatusBadRequest, response)
				return
			}

			newUser := models.User{
				Id:       userID,
				Email:    user.Email,
				FullName: user.Email,
			}

			_, err := userCollection.InsertOne(ctx, newUser)
			if err != nil {
				response := respones.CustomResponse{
					Status:      http.StatusBadRequest,
					Message:     ERROR,
					Description: MONGO_ERROR,
					Data:        err.Error(),
				}
				c.JSON(http.StatusBadRequest, response)
				return
			}

			response := respones.CustomResponse{
				Status:      http.StatusCreated,
				Message:     SUCCESS,
				Description: USER_SIGNUP_SUCCESS,
			}
			c.JSON(http.StatusCreated, response)
		}
	}
}

func SignInUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Retrieves the first matching document
		userID := c.MustGet("userID").(string)

		var result models.UserResponse
		err := userCollection.FindOne(ctx, bson.M{"_id": userID}).Decode(&result)
		if err != nil {
			log.Printf("Failed to verify ID token: %v", err)
			response := respones.CustomResponse{
				Status:      http.StatusUnauthorized,
				Message:     ERROR,
				Description: "error firebase find one",
				Data:        nil,
			}
			c.JSON(http.StatusUnauthorized, response)
			return
		}

		response := respones.CustomResponse{
			Status:      http.StatusOK,
			Message:     SUCCESS,
			Description: USER_SIGNIN_SUCCESS,
		}
		c.JSON(http.StatusCreated, response)

	}
}

func SignUpUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// validate the request body
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			response := respones.CustomResponse{
				Status:      http.StatusBadRequest,
				Message:     ERROR,
				Description: EMAIL_PASSWORD_NULL,
				Data:        err.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		// map newUser on save on database
		userID := c.MustGet("userID").(string)
		newUser := models.User{
			Id:       userID,
			Email:    user.Email,
			FullName: user.FullName,
		}

		_, err := userCollection.InsertOne(ctx, newUser)
		if err != nil {
			response := respones.CustomResponse{
				Status:      http.StatusBadRequest,
				Message:     ERROR,
				Description: MONGO_ERROR,
				Data:        err.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}

		// Map response succss and sending client
		response := respones.CustomResponse{
			Status:      http.StatusCreated,
			Message:     SUCCESS,
			Description: USER_SIGNUP_SUCCESS,
			Data:        nil,
		}
		c.JSON(http.StatusCreated, response)
	}
}
