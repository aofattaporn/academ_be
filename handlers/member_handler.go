package handlers

import (
	"academ_be/models"
	"academ_be/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProjectMembers(c *gin.Context) {

	projectId := c.Param("projectsId")
	if projectId == "" {
		handleBussinessError(c, "Can't to find your Tasks ID")
	}

	project, err := services.GetProjectById(c, projectId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	// modeling response
	var memberSetting models.AllMemberProject
	memberSetting.Members = project.Members
	memberSetting.Roles = project.Roles
	memberSetting.Invite = []models.Invite{}

	handleSuccess(c, http.StatusCreated, SUCCESS, GET_MY_PROJECT_SUCCESS, memberSetting)

}

func ChangeRoleMember(c *gin.Context) {

	projectId := c.Param("projectsId")
	if projectId == "" {
		handleBussinessError(c, "Can't to find your Tasks ID")
	}

	memberId := c.Param("memberId")
	if memberId == "" {
		handleBussinessError(c, "Can't to find your Member ID")
	}

	roleId := c.Param("roleId")
	if memberId == "" {
		handleBussinessError(c, "Can't to find your Member ID")
	}

	err := services.UpdateRoleByMemberID(c, projectId, memberId, roleId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	project, err := services.GetProjectById(c, projectId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	// modeling response
	var memberSetting models.AllMemberProject
	memberSetting.Members = project.Members
	memberSetting.Roles = project.Roles
	memberSetting.Invite = []models.Invite{}

	handleSuccess(c, http.StatusCreated, SUCCESS, GET_MY_PROJECT_SUCCESS, memberSetting)
}
