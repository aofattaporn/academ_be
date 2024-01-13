package middlewares

import (
	"academ_be/respones"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			// Map response succss and sending client
			response := respones.UserResponse{
				Status:      http.StatusInternalServerError,
				Message:     "Error",
				Description: "Internal Error",
				Data:        nil,
			}
			c.JSON(http.StatusInternalServerError, response)
		}

	}

}
