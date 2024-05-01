package handlers

import (
	"academ_be/models"
	"academ_be/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetProjectRoleAndPermissions(c *gin.Context) {

	userID := c.MustGet(USER_ID).(string)
	projectId := c.Param("projectId")
	if projectId == "" {
		handleBussinessError(c, "Can't to find your Tasks ID")
	}

	roleAndRolePermission := getRolesAndPermissionIdByProject(c, projectId, userID)

	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_PROJECT_SUCCESS, roleAndRolePermission)

}

func CreateProjectRoleAndPermissions(c *gin.Context) {

	userID := c.MustGet(USER_ID).(string)
	projectId := c.Param("projectId")
	if projectId == "" {
		handleBussinessError(c, "Can't to find your Tasks ID")
	}

	var createProject models.CreateRole
	if err := c.BindJSON(&createProject); err != nil {
		handleBussinessError(c, err.Error())
		return
	}

	newId := primitive.NewObjectID()
	memberPermissionId := setUpOwnerPermission(c, FLAG_DEFAULT_MEMBER)

	newRoles := models.Role{
		RoleId: newId, RoleName: createProject.NewRole, PermissionId: memberPermissionId,
	}

	err := services.CreateNewRole(c, projectId, newRoles)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	roleAndRolePermission := getRolesAndPermissionIdByProject(c, projectId, userID)

	handleSuccess(c, http.StatusCreated, SUCCESS, GET_MY_PROJECT_SUCCESS, roleAndRolePermission)

}

func UpdateRoleName(c *gin.Context) {

	userID := c.MustGet(USER_ID).(string)
	projectId := c.Param("projectId")
	if projectId == "" {
		handleBussinessError(c, "Can't to find your Tasks ID")
	}

	var role models.CreateRole
	if err := c.BindJSON(&role); err != nil {
		handleBussinessError(c, err.Error())
		return
	}

	roleId := c.Param("roleId")
	if projectId == "" {
		handleBussinessError(c, "Can't to find your Role ID")
	}

	err := services.UpdateRoleName(c, projectId, roleId, role.NewRole)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	roleAndRolePermission := getRolesAndPermissionIdByProject(c, projectId, userID)

	handleSuccess(c, http.StatusCreated, SUCCESS, GET_MY_PROJECT_SUCCESS, roleAndRolePermission)

}

func DeleteRole(c *gin.Context) {

	userID := c.MustGet(USER_ID).(string)
	projectId := c.Param("projectId")
	if projectId == "" {
		handleBussinessError(c, "Can't to find your Tasks ID")
	}

	roleId := c.Param("roleId")
	if projectId == "" {
		handleBussinessError(c, "Can't to find your Role ID")
	}

	err := services.DeleteRole(c, projectId, roleId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	// TODO : Delete Each Permission id within roles ?

	roleAndRolePermission := getRolesAndPermissionIdByProject(c, projectId, userID)

	handleSuccess(c, http.StatusCreated, SUCCESS, GET_MY_PROJECT_SUCCESS, roleAndRolePermission)

}

func UpdatePermission(c *gin.Context) {

	userID := c.MustGet(USER_ID).(string)
	var updatePermission models.Permission
	if err := c.BindJSON(&updatePermission); err != nil {
		handleBussinessError(c, err.Error())
		return
	}

	projectId := c.Param("projectId")
	if projectId == "" {
		handleBussinessError(c, "Can't to find your Tasks ID")
	}

	permissionId := c.Param("permissionId")
	if permissionId == "" {
		handleBussinessError(c, "Can't to find your Permission ID")
	}

	err := services.UpdatePermission(c, permissionId, updatePermission)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	roleAndRolePermission := getRolesAndPermissionIdByProject(c, projectId, userID)

	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_PROJECT_SUCCESS, roleAndRolePermission)

}

func getRolesAndPermissionIdByProject(c *gin.Context, projectId string, userID string) (roleAndRolePermission models.RoleAndRolePermission) {

	project, err := services.GetProjectById(c, projectId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	var roleAndPermissions []models.RoleAndFullPermission

	for _, role := range project.Roles {
		temp, err := services.GetPermission(c, role.PermissionId)
		if err != nil {
			handleTechnicalError(c, err.Error())
			return
		}

		roleAndPermission := models.RoleAndFullPermission{
			RoleId:     role.RoleId,
			RoleName:   role.RoleName,
			Permission: *temp,
		}

		roleAndPermissions = append(roleAndPermissions, roleAndPermission)

	}

	permission, err := getPermissionIdByUser(c, project, userID)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	return models.RoleAndRolePermission{
		RolesAndFullPermission: roleAndPermissions,
		RolePermission:         permission.Role,
	}
}
