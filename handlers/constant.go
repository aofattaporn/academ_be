package handlers

import (
	"academ_be/respones"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (

	// Response Status constant
	SUCCESS         string = "SUCCESS"
	REQUEST_SUCCESS string = "REQUEST SUCCESS"
	CREATE_SUCCESS  string = "CREATE SUCCESS"

	USER_SIGNIN_SUCCESS    string = "User Sign In Sccuess"
	USER_SIGNUP_SUCCESS    string = "User Sign Up Sccuess"
	CREATE_PROJECT_SUCCESS string = "CREATE PROJECT SUCCESS"
	GET_MY_PROJECT_SUCCESS string = "GET MY PROJECT SUCCESS"
	GET_MY_TASKS_SUCCESS   string = "GET MY TASKS SUCCESS"

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

	PROCESS_DEFAULT_TODO        string = "Todo"
	PROCESS_DEFAULT_IN_PROGRESS string = "In Progress"
	PROCESS_DEFAULT_DONE        string = "Done"

	ROLE_DEFAULT_OWNER  string = "Owner"
	ROLE_DEFAULT_MEMBER string = "Member"
	FLAG_DEFAULT_OWNER  bool   = true
	FLAG_DEFAULT_MEMBER bool   = false

	// validation constant
	TOKEN     string = "token"
	USER_ID   string = "userID"
	USER_NAME string = "userName"

	// Notification
	NOTI_HEADER_TASKS_UPDATED    string = "Tasks Assignee"
	NOTI_BODY_TASKS_UPDATED      string = "The tasks you assigned have been updated"
	NOTI_HEADER_PROJECT_DEADLINE string = "Project Deadline Alert"
	NOTI_BODY_PROJECT_DEADLINE   string = "This project is approaching its deadline."
)

var DEFULT_COLORS = []string{
	"#AF8AE2",
	"#6985FF",
	"#3FB1B2",
	"#64C7A2",
	"#FABE34",
	"#E78945",
	"#DD646A",
	"#BDBDBD",
}

func handleBussinessError(c *gin.Context, description string) {
	response := respones.CustomResponse{
		Status:      CODE_STATUS_1899,
		Message:     BUSSINESS_ERROR,
		Description: description,
	}
	c.AbortWithStatusJSON(http.StatusOK, response)
}

func handleTechnicalError(c *gin.Context, description string) {
	response := respones.CustomResponse{
		Status:      CODE_STATUS_1999,
		Message:     TECHNICAL_ERROR,
		Description: description,
	}
	c.AbortWithStatusJSON(http.StatusOK, response)
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
