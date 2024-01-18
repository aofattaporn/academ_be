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
			response := respones.CustomResponse{
				Status:      http.StatusInternalServerError,
				Message:     ERROR,
				Description: c.Errors.String(),
				Data:        nil,
			}
			c.JSON(http.StatusInternalServerError, response)
		}

	}

}
