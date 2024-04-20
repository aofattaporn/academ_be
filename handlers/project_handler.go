package handlers

import (
	"academ_be/models"
	"academ_be/services"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetProject godoc
// @summary Health Check
// @description Health checking for the service
// @id GetProject
// @tags projects
// @accept json
// @produce json
// @response 200 {string} string "OK"
// @router /api/v1/projects/:projectsId [get]
func GetProjectById(c *gin.Context) {
	// Extract the user ID from the request context
	userID := c.MustGet(USER_ID).(string)
	projectID := c.Param("projectsId")

	// Retrieve the project by ID
	project, err := services.GetProjectById(c, projectID)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	// Find the role ID corresponding to the user
	var roleID primitive.ObjectID
	for _, member := range project.Members {
		if member.UserId == userID {
			roleID = member.RoleId
			break
		}
	}

	// Find the permission ID corresponding to the role
	var permissionID primitive.ObjectID
	for _, role := range project.Roles {
		if role.RoleId == roleID {
			permissionID = role.PermissionId
			break
		}
	}

	// Retrieve the permission by ID
	permission, err := services.GetPermission(c, permissionID)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	projectInfo := models.ProjectInfoPermission{
		ProjectInfo:    *project,
		TaskPermission: permission.Task,
	}

	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_PROJECT_SUCCESS, projectInfo)
}

// GetAllMyProjects godoc
// @summary Health Check
// @description Health checking for the service
// @id GetAllMyProjects
// @tags projects
// @accept json
// @produce json
// @response 200 {string} string "OK"
// @router /api/v1/projects/users/id [get]
func GetAllMyProjects(c *gin.Context) {
	// Extract the user_id from the request parameters
	userID := c.MustGet(USER_ID).(string)

	// Call your business logic function to get projects by user ID
	projects, err := services.GetProjectsByMemberUserID(c, userID)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	handleSuccess(c, http.StatusCreated, SUCCESS, GET_MY_PROJECT_SUCCESS, projects)
}

// CreateProject godoc
// @summary Health Check
// @description Health checking for the service
// @id CreateProject
// @tags projects
// @accept json
// @produce json
// @response 200 {string} string "OK"
// @router /api/v1/projects/users/id [post]
func CreateProject(c *gin.Context) {

	// Mapping request project body
	var createProject models.CreateProject
	if err := c.BindJSON(&createProject); err != nil {
		handleBussinessError(c, err.Error())
		return
	}

	// Get the user ID from the context
	userId := c.MustGet("userID").(string)
	userName := c.GetHeader(USER_NAME)

	// Set up roles and permission
	ownerId := primitive.NewObjectID()
	memberId := primitive.NewObjectID()
	ownerPermissionId := setUpOwnerPermission(c, FLAG_DEFAULT_OWNER)
	memberPermissionId := setUpOwnerPermission(c, FLAG_DEFAULT_MEMBER)
	roles := []models.Role{
		{RoleId: ownerId, RoleName: ROLE_DEFAULT_OWNER, PermissionId: ownerPermissionId},
		{RoleId: memberId, RoleName: ROLE_DEFAULT_MEMBER, PermissionId: memberPermissionId},
	}

	// Set up processes
	processes := setUpProcesses()

	// Set up members
	members := []models.Member{
		{UserId: userId, UserName: userName, RoleId: ownerId},
	}
	now := time.Now()

	// Create a new project instance
	projectId := primitive.NewObjectID()
	newProject := models.Project{
		ProjectId: projectId,
		ProjectProfile: models.ProjectProfile{
			ProjectName: createProject.ProjectName,
			AvatarColor: getRandomColor(),
		},
		ProjectStartDate: &now,
		ProjectEndDate:   createProject.ProjectEndDate,
		Views:            createProject.Views,
		Invite:           []models.Invite{},
		CreatedAt:        &now,
		UpdatedAt:        &now,
		Process:          processes,
		Members:          members,
		Roles:            roles,
	}

	// Create the project in the database
	err := services.CreateProject(c, &newProject)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	// Create a new project instance
	newProjectRes := models.MyProject{
		ProjectId:      projectId,
		ProjectProfile: newProject.ProjectProfile,
		ProjectEndDate: newProject.ProjectEndDate,
		Members:        newProject.Members,
	}

	handleSuccess(c, http.StatusOK, SUCCESS, CREATE_PROJECT_SUCCESS, newProjectRes)
}

func setUpProcesses() []models.Process {
	processStr := []string{PROCESS_DEFAULT_TODO, PROCESS_DEFAULT_IN_PROGRESS, PROCESS_DEFAULT_DONE}
	processes := make([]models.Process, len(processStr))
	for i, v := range processStr {
		processes[i] = models.Process{
			ProcessId:    primitive.NewObjectID(),
			ProcessName:  v,
			ProcessColor: getRandomColor(),
		}
	}
	return processes
}

func setUpOwnerPermission(c *gin.Context, flag bool) primitive.ObjectID {
	ownerPermissionID := primitive.NewObjectID()

	// Create permission service
	permission := models.Permission{
		Members: models.MembersPermission{
			AddRole: flag,
			Invite:  flag,
			Remove:  flag,
		},
		Project: models.ProjectPermission{
			EditProfile: flag,
			ManageView:  flag,
		},
		Task: models.TaskPermission{
			AddNew:        flag,
			Delete:        flag,
			Edit:          flag,
			ManageProcess: flag,
		},
		Role: models.RolePermission{
			AddNew: flag,
			Edit:   flag,
			Delete: flag,
		},
	}

	// Use permission as needed
	services.CreatePermission(c, &permission, ownerPermissionID)

	return ownerPermissionID
}

func getRandomColor() string {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	randomIndex := random.Intn(len(DEFULT_COLORS))
	return DEFULT_COLORS[randomIndex]
}

func GetProjectDetails(c *gin.Context) {

	projectId := c.Param("projectsId")
	if projectId == "" {
		handleBussinessError(c, "Can't to find your Tasks ID")
	}

	// Retrieve the project by ID
	projectDetails, err := services.GetProjectDetails(c, projectId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	handleSuccess(c, http.StatusCreated, SUCCESS, GET_MY_PROJECT_SUCCESS, projectDetails)

}

func UpdateProjectDetails(c *gin.Context) {

	projectId := c.Param("projectsId")
	if projectId == "" {
		handleBussinessError(c, "Can't to find your Tasks ID")
	}

	// update-project

	var projectUpdate models.ProjectDetails
	if err := c.BindJSON(&projectUpdate); err != nil {
		handleBussinessError(c, err.Error())
		return
	}

	if projectUpdate.ProjectStartDate != nil && projectUpdate.ProjectStartDate.IsZero() {
		projectUpdate.ProjectStartDate = nil
	}

	if projectUpdate.ProjectEndDate != nil && projectUpdate.ProjectEndDate.IsZero() {
		projectUpdate.ProjectEndDate = nil
	}

	err := services.UpdateProjectDetails(c, projectId, projectUpdate)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	// Retrieve the project by ID
	projectDetails, err := services.GetProjectDetails(c, projectId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	handleSuccess(c, http.StatusCreated, SUCCESS, GET_MY_PROJECT_SUCCESS, projectDetails)

}
