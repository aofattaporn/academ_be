package handlers

import (
	"academ_be/models"
	"academ_be/respones"
	"academ_be/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateResource(c *gin.Context) {
	// Implement your logic to create a resource in MongoDB

	// validate the request body
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		response := respones.UserResponse{
			Status:      http.StatusBadRequest,
			Message:     "ERROR222",
			Description: "EMAIL_PASSWORD_NULL",
			Data:        err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	user.Id = "1"

	err := services.CreateUser(&user)
	if err != nil {
		response := respones.UserResponse{
			Status:      http.StatusBadRequest,
			Message:     "ERROR1111",
			Description: "EMAIL_PASSWORD_NULL",
			Data:        err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	response := respones.UserResponse{
		Status:      http.StatusCreated,
		Message:     "SUCCESS",
		Description: "USER_SIGNUP_SUCCESS",
	}
	c.JSON(http.StatusCreated, response)
}
