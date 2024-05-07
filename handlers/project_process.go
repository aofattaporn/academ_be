package handlers

import (
	"academ_be/models"
	"academ_be/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateProcess(c *gin.Context) {

	userID := c.MustGet(USER_ID).(string)

	projectId := c.Param("projectId")
	if projectId == "" {
		handleBussinessError(c, "Can't to find your Project ID")
	}

	processId := c.Param("processId")
	if processId == "" {
		handleBussinessError(c, "Can't to find your Process ID")
	}

	viewName := c.Param("viewName")
	if viewName == "" {
		handleBussinessError(c, "Can't to find your Tasks ID")
	}

	var updaateProcess models.Process
	if err := c.BindJSON(&updaateProcess); err != nil {
		handleBussinessError(c, err.Error())
		return
	}

	err := services.UpdateProcessByID(c, projectId, processId, updaateProcess)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	project, err := services.GetProjectById(c, projectId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	permission, err := getPermissionIdByUser(c, project, userID)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	projectInfo := models.ProjectInfoPermission{
		NavigateView:      viewName,
		ProjectInfo:       *project,
		TaskPermission:    permission.Task,
		ProjectPermission: permission.Project,
	}

	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_PROJECT_SUCCESS, projectInfo)

}

func DeleteProcess(c *gin.Context) {

	userID := c.MustGet(USER_ID).(string)

	projectId := c.Param("projectId")
	if projectId == "" {
		handleBussinessError(c, "Can't to find your Tasks ID")
	}

	processId := c.Param("processId")
	if processId == "" {
		handleBussinessError(c, "Can't to find your Tasks ID")
	}

	viewName := c.Param("viewName")
	if viewName == "" {
		handleBussinessError(c, "Can't to find your Tasks ID")
	}

	err := services.DeleteProcessbyId(c, projectId, processId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	project, err := services.GetProjectById(c, projectId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	permission, err := getPermissionIdByUser(c, project, userID)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	projectInfo := models.ProjectInfoPermission{
		NavigateView:      viewName,
		ProjectInfo:       *project,
		TaskPermission:    permission.Task,
		ProjectPermission: permission.Project,
	}

	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_PROJECT_SUCCESS, projectInfo)

}

func CreateNewProcess(c *gin.Context) {

	userID := c.MustGet(USER_ID).(string)

	projectId := c.Param("projectId")
	if projectId == "" {
		handleBussinessError(c, "Can't to find your Project ID")
	}

	var updaateProcess models.Process
	if err := c.BindJSON(&updaateProcess); err != nil {
		handleBussinessError(c, err.Error())
		return
	}

	viewName := c.Param("viewName")
	if viewName == "" {
		handleBussinessError(c, "Can't to find your Tasks ID")
	}

	err := services.CreateProcessByID(c, projectId, updaateProcess)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	project, err := services.GetProjectById(c, projectId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	permission, err := getPermissionIdByUser(c, project, userID)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	projectInfo := models.ProjectInfoPermission{
		NavigateView:      viewName,
		ProjectInfo:       *project,
		TaskPermission:    permission.Task,
		ProjectPermission: permission.Project,
	}

	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_PROJECT_SUCCESS, projectInfo)

}
