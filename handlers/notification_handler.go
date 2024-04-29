package handlers

import (
	"academ_be/models"
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

func UpdateClearNotiById(c *gin.Context) {

	notiId := c.Param("notiId")
	if notiId == "" {
		handleBussinessError(c, "Can't to find your Tasks ID")
	}

	var updateIsClearNoti models.UpdateIsClearNoti
	if err := c.BindJSON(&updateIsClearNoti); err != nil {
		handleBussinessError(c, err.Error())
		return
	}

	err := services.UpdateClearNotiById(c, notiId, updateIsClearNoti.IsClear)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_PROJECT_SUCCESS, nil)
}

func DeleteNotiById(c *gin.Context) {

	notiId := c.Param("notiId")
	if notiId == "" {
		handleBussinessError(c, "Can't to find your Tasks ID")
	}

	err := services.DeleteNotiById(c, notiId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_PROJECT_SUCCESS, nil)
}
