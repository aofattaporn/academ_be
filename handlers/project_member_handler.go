package handlers

import (
	"academ_be/models"
	"academ_be/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getProjectMembers(c *gin.Context, projectId string) (*models.AllMemberProject, error) {
	project, err := services.GetProjectById(c, projectId)
	if err != nil {
		return nil, err
	}

	return &models.AllMemberProject{
		Members: project.Members,
		Roles:   project.Roles,
		Invites: project.Invites,
	}, nil
}

func GetProjectMembers(c *gin.Context) {
	projectId := c.Param("projectId")
	if projectId == "" {
		handleBussinessError(c, "Can't find your Project ID")
		return
	}

	memberSetting, err := getProjectMembers(c, projectId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_PROJECT_SUCCESS, memberSetting)
}

func ChangeRoleMember(c *gin.Context) {
	projectId := c.Param("projectId")
	if projectId == "" {
		handleBussinessError(c, "Can't find your Project ID")
		return
	}

	memberId := c.Param("memberId")
	if memberId == "" {
		handleBussinessError(c, "Can't find your Member ID")
		return
	}

	roleId := c.Param("roleId")
	if roleId == "" {
		handleBussinessError(c, "Can't find your Role ID")
		return
	}

	err := services.UpdateRoleByMemberID(c, projectId, memberId, roleId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	memberSetting, err := getProjectMembers(c, projectId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_PROJECT_SUCCESS, memberSetting)
}
