package handlers

import (
	"academ_be/models"
	"academ_be/services"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetAllTasks godoc
// @summary Health Check
// @description Health checking for the service
// @id GetAllTasks
// @tags users
// @accept json
// @produce json
// @response 200 {string} string "OK"
// @router /api/v1/sign-in [post]
func GetAllTasksByProjectId(c *gin.Context) {

	projectId := c.Param("projectId")
	tasks, err := services.GetAllTasksByProjectId(c, projectId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_TASKS_SUCCESS, tasks)

}

// CreateTasks godoc
// @summary Health Check
// @description Health checking for the service
// @id CreateTasks
// @tags users
// @accept json
// @produce json
// @response 200 {string} string "OK"
// @router /api/v1/sign-in [post]
func CreateTasks(c *gin.Context) {

	var createTasks models.Tasks
	if err := c.BindJSON(&createTasks); err != nil {
		handleBussinessError(c, err.Error())
		return
	}

	if err := services.CreateTasks(c, &createTasks); err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	tasks, err := services.GetAllTasksByProjectId(c, createTasks.ProjectId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	// Return success response
	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_TASKS_SUCCESS, tasks)

}

// GetTasksById godoc
// @summary Health Check
// @description Health checking for the service
// @id GetTasksById
// @tags users
// @accept json
// @produce json
// @response 200 {string} string "OK"
// @router /api/v1/sign-in [post]
func GetTasksById(c *gin.Context) {

	tasksId := c.Param("taskId")
	if tasksId == "" {
		handleBussinessError(c, "Can't to find your Tasks ID")
	}

	tasks, err := services.GetTasksByProjectId(c, tasksId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	// Return success response
	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_TASKS_SUCCESS, &tasks)

}

// GetAllTasks godoc
// @summary Health Check
// @description Health checking for the service
// @id GetAllTasks
// @tags users
// @accept json
// @produce json
// @response 200 {string} string "OK"
// @router /api/v1/sign-in [post]
func ChangeProcesss(c *gin.Context) {

	tasksId := c.Param("taskId")
	if tasksId == "" {
		handleBussinessError(c, "Can't to find your Tasks ID")
	}

	processId := c.Param("processId")
	if processId == "" {
		handleBussinessError(c, "Can't to find your Process ID")
	}

	err := services.ChangeProcesss(c, tasksId, processId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

}

// GetAllTasks godoc
// @summary Health Check
// @description Health checking for the service
// @id GetAllTasks
// @tags users
// @accept json
// @produce json
// @response 200 {string} string "OK"
// @router /api/v1/sign-in [post]
func UpdateTasks(c *gin.Context) {
	tasksId := c.Param("taskId")
	if tasksId == "" {
		handleBussinessError(c, "Can't find your Tasks ID")
		return
	}

	var updateTasks models.Tasks
	if err := c.BindJSON(&updateTasks); err != nil {
		handleBussinessError(c, err.Error())
		return
	}

	if updateTasks.StartDate != nil && updateTasks.StartDate.IsZero() {
		updateTasks.StartDate = nil
	}

	if updateTasks.DueDate != nil && updateTasks.DueDate.IsZero() {
		updateTasks.DueDate = nil
	}

	// Update Tasks
	_, err := services.UpdateTasksByTaskId(c, tasksId, updateTasks)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	// Return success response
	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_TASKS_SUCCESS, updateTasks)

	// push noti on assignee
	if updateTasks.Assignee != nil && updateTasks.Assignee.UserId != "" {
		project, err := services.GetProjectById(c, updateTasks.ProjectId)
		if err != nil {
			handleTechnicalError(c, err.Error())
			return
		}

		// Sent Notification on browser
		sendNotificationByUserId(c, updateTasks.Assignee.UserId, &project.ProjectProfile)

		// Sent Notification on email
		invitationToken := generateInvitationToken()
		err = sendInvite(updateTasks.Assignee.Emaill, project.ProjectProfile.ProjectName, invitationToken)
		if err != nil {
			handleTechnicalError(c, err.Error())
			return
		}
	}
}

// DeleteTasksById godoc
// @summary Health Check
// @description Health checking for the service
// @id DeleteTasksById
// @tags users
// @accept json
// @produce json
// @response 200 {string} string "OK"
// @router /api/v1/sign-in [delete]
func DeleteTasksById(c *gin.Context) {

	tasksId := c.Param("taskId")
	if tasksId == "" {
		handleBussinessError(c, "Can't to find your Tasks ID")
	}

	err := services.DeleteTasksByTasksId(c, tasksId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	// Return success response
	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_TASKS_SUCCESS, nil)

}

func GetAllTasksEachProject(c *gin.Context) {

	var allMytasks models.AllMyTasks

	// getting userID
	userID := c.MustGet(USER_ID).(string)
	projects, err := services.GetProjectByUserId(c, userID)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}
	allMytasks.Projects = projects

	tasks, err := services.GetTasksByUserId(c, userID)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}
	allMytasks.Tasks = tasks

	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_TASKS_SUCCESS, allMytasks)

}

func GetAllTasksHomePage(c *gin.Context) {

	// getting userID
	userID := c.MustGet(USER_ID).(string)

	tasks, err := services.GetTasksByUserId(c, userID)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_TASKS_SUCCESS, tasks)

}

func sendNotificationByUserId(c *gin.Context, userId string, projectProfile *models.ProjectProfile) {

	fmt.Println("Send noti start")
	fcm, err := services.FindFCMByMember(c, userId)
	if err != nil {
		fmt.Println(err)

		return
	}
	fmt.Println(fcm)
	fmt.Println("Send noti end")

	now := time.Now()
	noti := models.Notification{
		UserId:         userId,
		ProjectProfile: *projectProfile,
		Title:          "Project Assignee",
		Body:           "You are assigned a tasks",
		Date:           &now,
		IsClear:        false,
	}

	err = services.AddNotification(c, fcm.FCM_TOKEN, noti)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

}
func sendNotificationByUserIdCron(userId string, projectProfile *models.ProjectProfile) {

	fmt.Println("Send noti start")
	fcm, err := services.FindFCMByMemberCron(userId)
	if err != nil {
		fmt.Println(err)

		return
	}
	fmt.Println(fcm)
	fmt.Println("Send noti end")

	now := time.Now()
	noti := models.Notification{
		UserId:         userId,
		ProjectProfile: *projectProfile,
		Title:          "Project Assignee",
		Body:           "You are assigned a tasks",
		Date:           &now,
		IsClear:        false,
	}

	err = services.AddNotificationCron(fcm.FCM_TOKEN, noti)
	if err != nil {
		fmt.Println(err)
		return
	}

}
