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

	ERROR                 string = "ERROR"
	ACCESS_FORBIDEN_ERROR string = "ACCESS FORBIDEN ERROR"
	BUSSINESS_ERROR       string = "BUSSINESS ERROR"
	TECHNICAL_ERROR       string = "TECHNICAL ERROR"

	CODE_STATUS_1000 int16 = 1000
	CODE_STATUS_1799 int16 = 1799
	CODE_STATUS_1899 int16 = 1899
	CODE_STATUS_1999 int16 = 1999

	EMAIL_PASSWORD_NULL string = "Email or Password is null"
	INPUT_INVALID       string = "Input is invalid"
	MONGO_ERROR         string = "Something wrong on mongoDB"
	INVALID_TOKEN       string = "Invalid ID token"
	MISSING_AUTH_HEADER string = "Missing Authorization header"

	Authorization string = "Authorization"

	// validation constant
	TOKEN string = "token"
)

func handleBussinessError(c *gin.Context, message, description string) {
	response := respones.CustomResponse{
		Status:      CODE_STATUS_1899,
		Message:     BUSSINESS_ERROR,
		Description: message,
		Data:        description,
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, response)
}

func handleTechnicalError(c *gin.Context, message, description string) {
	response := respones.CustomResponse{
		Status:      CODE_STATUS_1999,
		Message:     TECHNICAL_ERROR,
		Description: message,
		Data:        description,
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, response)
}

func handleSuccess(c *gin.Context, statusCode int, message, description string, data interface{}) {
	response := respones.CustomResponse{
		Status:      CODE_STATUS_1000,
		Message:     message,
		Description: description,
		Data:        data,
	}
	c.JSON(statusCode, response)
}
