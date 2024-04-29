package handlers

import (
	"academ_be/models"
	"academ_be/services"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @summary Health Check
// @description Health checking for the service
// @id CreateUser
// @tags users
// @accept json
// @produce json
// @response 200 {string} string "OK"
// @router /api/v1/sign-up [post]
func CreateUser(c *gin.Context) {

	// mapping request body
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		handleBussinessError(c, err.Error())
		return
	}

	// mapping save data on database
	userID := c.MustGet(USER_ID).(string)
	newUser := models.User{
		Id:       userID,
		Email:    user.Email,
		FullName: user.FullName,
	}

	err := services.CreateUser(c, &newUser)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	handleSuccess(c, http.StatusCreated, SUCCESS, USER_SIGNUP_SUCCESS, nil)
}

// GetUser godoc
// @summary Health Check
// @description Health checking for the service
// @id GetUser
// @tags users
// @accept json
// @produce json
// @response 200 {string} string "OK"
// @router /api/v1/users [get]
func GetUser(c *gin.Context) {

	// getting userID
	var user *models.UserResponse
	userID := c.MustGet(USER_ID).(string)

	// find user in database from header
	user, err := services.FindUserOneById(c, userID)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}
	// update token device
	fcmToken := c.GetHeader("FCM_TOEKN")
	err = services.SaveFCMToken(c, userID, fcmToken)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	handleSuccess(c, http.StatusOK, SUCCESS, USER_SIGNUP_SUCCESS, user)
}

// CreateUserByGoogle godoc
// @summary Health Check
// @description Health checking for the service
// @id CreateUserByGoogle
// @tags users
// @accept json
// @produce json
// @response 200 {string} string "OK"
// @router /api/v1/sign-in [post]
func CreateUserByGoogle(c *gin.Context) {

	// getting userID
	userID := c.MustGet(USER_ID).(string)

	// find user and count
	count, err := services.FindUserAndCount(c, userID)
	if err != nil {
		handleTechnicalError(c, err.Error())
	}

	fmt.Println("****")

	if count < 1 {

		// no userId on database
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			handleBussinessError(c, err.Error())
			return
		}

		newUser := models.User{
			Id:          userID,
			Email:       user.Email,
			FullName:    user.FullName,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			AvatarColor: getRandomColor(),
		}
		services.CreateUser(c, &newUser)

		handleSuccess(c, http.StatusCreated, SUCCESS, USER_SIGNUP_SUCCESS, nil)
	} else {

		// update token device
		fcmToken := c.GetHeader("FCM_TOEKN")
		err = services.SaveFCMToken(c, userID, fcmToken)
		if err != nil {
			handleTechnicalError(c, err.Error())
			return
		}

		// already existing user in database
		handleSuccess(c, http.StatusCreated, SUCCESS, USER_SIGNUP_SUCCESS, nil)
	}
}
