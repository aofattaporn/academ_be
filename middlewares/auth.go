package middlewares

import (
	"academ_be/respones"
	"log"
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

const (
	AUTHORIZATION       string = "Authorization"
	ERROR               string = "ERROR"
	MISSING_AUTH_HEADER string = "MISSING_AUTH_HEADER"
	INVALID_TOKEN       string = "INVALID_TOKEN"
	BEARER              string = "Bearer "
	EMPTY_STRING        string = ""
)

func AuthRequire(admin *auth.Client) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get ID token from the request header
		idToken := c.GetHeader(AUTHORIZATION)
		if idToken == "" {
			response := respones.CustomResponse{
				Status:      http.StatusUnauthorized,
				Message:     ERROR,
				Description: MISSING_AUTH_HEADER,
				Data:        nil,
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Verify the ID token
		tokenString := strings.Replace(idToken, BEARER, EMPTY_STRING, 1)
		credential, err := admin.VerifyIDToken(c, tokenString)
		if err != nil {
			log.Printf("Failed to verify ID token: %v", err)
			response := respones.CustomResponse{
				Status:      http.StatusUnauthorized,
				Message:     ERROR,
				Description: INVALID_TOKEN,
				Data:        nil,
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Set example variable
		c.Set("userID", credential.UID)
		c.Next()
	}
}
