package handlers

import (
	"academ_be/models"
	"academ_be/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	var createTasks models.CreateTasks
	if err := c.BindJSON(&createTasks); err != nil {
		handleBussinessError(c, err.Error())
		return
	}

	createTasks.TasksId = primitive.NewObjectID()

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
		handleBussinessError(c, "Can't to find your Tasks ID")
	}

	var updateTasks models.UpdateTasks
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

	err := services.UpdateTasksByTaskId(c, tasksId, updateTasks)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	// Return success response
	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_TASKS_SUCCESS, &updateTasks)

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
