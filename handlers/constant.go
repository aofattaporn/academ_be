package handlers

import (
	"academ_be/respones"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (

	// Response Status constant
	SUCCESS             string = "Success"
	USER_SIGNIN_SUCCESS string = "User Sign In Sccuess"
	USER_SIGNUP_SUCCESS string = "User Sign Up Sccuess"

	ERROR               string = "Error"
	EMAIL_PASSWORD_NULL string = "Email or Password is null"
	INPUT_INVALID       string = "Input is invalid"
	MONGO_ERROR         string = "Something wrong on mongoDB"
	INVALID_TOKEN       string = "Invalid ID token"
	MISSING_AUTH_HEADER string = "Missing Authorization header"

	Authorization string = "Authorization"

	// validation constant
	TOKEN string = "token"
)

func handleBadRequest(c *gin.Context, message, description string) {
	response := respones.CustomResponse{
		Status:      http.StatusBadRequest,
		Message:     ERROR,
		Description: message,
		Data:        description,
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, response)
}

func handleSuccess(c *gin.Context, statusCode int, message, description string, data interface{}) {
	response := respones.CustomResponse{
		Status:      statusCode,
		Message:     message,
		Description: description,
		Data:        data,
	}
	c.JSON(statusCode, response)
}
