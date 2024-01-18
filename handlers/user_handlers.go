package handlers

import (
	"academ_be/models"
	"academ_be/respones"
	"academ_be/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userService services.UserService

func init() {
	userService = &services.ConcreteUserService{}
}

func SetUserService(us services.UserService) {
	userService = us
}

// Exported function
func CreateResource(c *gin.Context) {
	// validate the request body
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		response := respones.UserResponse{
			Status:      http.StatusBadRequest,
			Message:     "ERROR222",
			Description: "Input can't be mapped from JSON",
			Data:        err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err := userService.CreateUser(&user)
	if err != nil {
		response := respones.UserResponse{
			Status:      http.StatusBadRequest,
			Message:     "ERROR1111",
			Description: "Can't create user in service",
			Data:        err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := respones.UserResponse{
		Status:      http.StatusCreated,
		Message:     "SUCCESS",
		Description: "USER_SIGNUP_SUCCESS",
	}
	c.JSON(http.StatusCreated, response)
}
