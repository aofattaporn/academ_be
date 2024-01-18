package handlers

import (
	"academ_be/models"
	"academ_be/respones"
	"academ_be/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	// validate the request body
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		response := respones.CustomResponse{
			Status:      http.StatusBadRequest,
			Message:     ERROR,
			Description: EMAIL_PASSWORD_NULL,
			Data:        err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// map newUser on save on database
	userID := c.MustGet("userID").(string)
	newUser := models.User{
		Id:       userID,
		Email:    user.Email,
		FullName: user.FullName,
	}
	services.CreateUser(c, &newUser)

	// Map response succss and sending client
	response := respones.CustomResponse{
		Status:      http.StatusCreated,
		Message:     SUCCESS,
		Description: USER_SIGNUP_SUCCESS,
	}
	c.JSON(http.StatusCreated, response)
}

func GetUser(c *gin.Context) {}

func CreateUserByGoogle(c *gin.Context) {}
