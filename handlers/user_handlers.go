package handlers

import (
	"academ_be/models"
	"academ_be/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		handleBussinessError(c, ERROR, err.Error())
		return
	}

	userID := c.MustGet("userID").(string)
	newUser := models.User{
		Id:       userID,
		Email:    user.Email,
		FullName: user.FullName,
	}
	services.CreateUser(c, &newUser)

	handleSuccess(c, http.StatusCreated, SUCCESS, USER_SIGNUP_SUCCESS, nil)
}

func GetUser(c *gin.Context) {

	var user *models.UserResponse
	userID := c.MustGet("userID").(string)
	user, err := services.FindUserOneById(c, userID)
	if err != nil {
		handleBussinessError(c, ERROR, MONGO_ERROR)
		return
	}

	handleSuccess(c, http.StatusCreated, SUCCESS, USER_SIGNUP_SUCCESS, user)
}

func CreateUserByGoogle(c *gin.Context) {

	userID := c.MustGet("userID").(string)
	count, err := services.FindUserAndCount(c, userID)
	if err != nil {
		handleBussinessError(c, ERROR, MONGO_ERROR)
	}

	if count < 0 {

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			handleBussinessError(c, ERROR, err.Error())
			return
		}

		newUser := models.User{
			Id:       userID,
			Email:    user.Email,
			FullName: user.FullName,
		}
		services.CreateUser(c, &newUser)

		handleSuccess(c, http.StatusCreated, SUCCESS, USER_SIGNUP_SUCCESS, nil)
	} else {
		handleSuccess(c, http.StatusCreated, SUCCESS, USER_SIGNUP_SUCCESS, nil)
	}
}
