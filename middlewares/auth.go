package middlewares

import (
	"academ_be/configs"
	"academ_be/respones"
	"log"
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

var firebaseClient *auth.Client = configs.ConnectFirebase()

func AuthRequire() gin.HandlerFunc {
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
			return
		}

		// Verify the ID token
		tokenString := strings.Replace(idToken, "Bearer ", "", 1)

		credential, err := firebaseClient.VerifyIDToken(c, tokenString)
		if err != nil {
			log.Printf("Failed to verify ID token: %v", err)
			response := respones.UserResponse{
				Status:      http.StatusUnauthorized,
				Message:     "ERROR",
				Description: "INVALID_TOKEN",
				Data:        nil,
			}
			c.JSON(http.StatusUnauthorized, response)
			return
		}

		// Set example variable
		c.Set("userID", credential.UID)
		c.Next()
	}
}
