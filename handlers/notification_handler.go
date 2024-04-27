package handlers

import (
	"academ_be/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllMyNotification(c *gin.Context) {

	// Extract the user ID from the request context
	userID := c.MustGet(USER_ID).(string)

	notifications, err := services.GetAllNotifications(c, userID)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_PROJECT_SUCCESS, notifications)
}
