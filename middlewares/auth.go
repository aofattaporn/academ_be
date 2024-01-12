package middlewares

import (
	"academ_be/configs"
	"academ_be/respones"
	"context"
	"log"
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

var firebaseClient *auth.Client = configs.ConnectFirebase()

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get ID token from the request header
		idToken := c.GetHeader("Authorization")
		if idToken == "" {
			response := respones.UserResponse{
				Status:      http.StatusUnauthorized,
				Message:     "ERROR",
				Description: "MISSING_AUTH_HEADER",
				Data:        nil,
			}
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		// Verify token by Firebase admin
		tokenString := strings.Replace(idToken, "Bearer ", "", 1)
		credential, err := firebaseClient.VerifyIDToken(context.Background(), tokenString)
		if err != nil {
			log.Printf("Failed to verify ID token: %v", err)
			response := respones.UserResponse{
				Status:      http.StatusUnauthorized,
				Message:     "ERROR",
				Description: "INVALID_TOKEN",
				Data:        nil,
			}
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		// Set the user ID from the token in the context for further use
		c.Set("userID", credential.UID)

		c.Next()
	}
}
