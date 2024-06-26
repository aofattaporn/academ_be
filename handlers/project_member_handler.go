package handlers

import (
	"academ_be/models"
	"academ_be/services"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func getMemberAndMemberPermission(c *gin.Context, projectId string, userId string) *models.AllMemberAndPermission {

	memberSetting, err := getProjectMembers(c, projectId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return nil
	}

	project, err := services.GetProjectById(c, projectId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return nil
	}

	permission, err := getPermissionIdByUser(c, project, userId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return nil
	}

	return &models.AllMemberAndPermission{
		AllMemberProject:  *memberSetting,
		MembersPermission: permission.Members,
	}

}

func GetProjectMembers(c *gin.Context) {

	userID := c.MustGet(USER_ID).(string)
	projectId := c.Param("projectId")
	if projectId == "" {
		handleBussinessError(c, "Can't find your Project ID")
		return
	}

	membersAndPermission := getMemberAndMemberPermission(c, projectId, userID)

	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_PROJECT_SUCCESS, membersAndPermission)
}

func ChangeRoleMember(c *gin.Context) {

	userID := c.MustGet(USER_ID).(string)
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

	err := checkHaveOneOwner(c, roleId, projectId, userID)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	err = services.UpdateRoleByMemberID(c, projectId, memberId, roleId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	membersAndPermission := getMemberAndMemberPermission(c, projectId, userID)

	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_PROJECT_SUCCESS, membersAndPermission)
}

func RemoveMember(c *gin.Context) {

	userID := c.MustGet(USER_ID).(string)
	projectId := c.Param("projectId")
	if projectId == "" {
		handleBussinessError(c, "Can't find your Project ID")
		return
	}

	memberId := c.Param("memberId")
	if memberId == "" {
		handleBussinessError(c, "Can't find your Project ID")
		return
	}

	err := services.RemoveMember(c, projectId, memberId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	membersAndPermission := getMemberAndMemberPermission(c, projectId, userID)

	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_PROJECT_SUCCESS, membersAndPermission)

}

func checkHaveOneOwner(c *gin.Context, roleIdChange string, projectId string, memberId string) (err error) {

	// Convert projectId string to ObjectId
	roleIDChange, err := primitive.ObjectIDFromHex(roleIdChange)
	if err != nil {
		return err
	}

	project, err := services.GetProjectById(c, projectId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	// find owner
	var ownerId primitive.ObjectID
	for _, r := range project.Roles {
		if r.RoleName == "Owner" {
			ownerId = r.RoleId
			break
		}
	}

	var myRoleId primitive.ObjectID
	var countOwner int
	for _, m := range project.Members {
		if m.UserId == memberId {
			myRoleId = m.RoleId
		}

		if m.RoleId == ownerId {
			countOwner++
		}
	}

	if ownerId == myRoleId && countOwner == 1 && roleIDChange != ownerId {
		return errors.New("Please assign the ownership to someone else first.")
	}

	return nil

}
