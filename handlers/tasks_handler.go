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
	// Bind JSON data to createTasks struct
	var createTasks models.CreateTasks
	if err := c.BindJSON(&createTasks); err != nil {
		handleBussinessError(c, err.Error())
		return
	}

	// Generate a new tasks ID
	createTasks.TasksId = primitive.NewObjectID()

	// Create tasks
	if err := services.CreateTasks(c, &createTasks); err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	// Get all tasks by project ID
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

}
