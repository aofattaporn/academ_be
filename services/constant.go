package services

import (
	"academ_be/respones"
	"log"

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

func handleError(c *gin.Context, err error, statusCode int, message string) {
	log.Printf("Error: %v", err)
	response := respones.CustomResponse{
		Status:      statusCode,
		Message:     message,
		Description: err.Error(),
	}
	c.AbortWithStatusJSON(statusCode, response)
}
