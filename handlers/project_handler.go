package handlers

import (
	"academ_be/models"
	"academ_be/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllMyProjects(c *gin.Context) {
	// Extract the user_id from the request parameters
	userID := c.MustGet("userID").(string)

	// Call your business logic function to get projects by user ID
	projects, err := services.GetProjectsByMemberUserID(c, userID)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	handleSuccess(c, http.StatusCreated, SUCCESS, USER_SIGNUP_SUCCESS, projects)
}

func CreateProject(c *gin.Context) {

	// Mapping request project body
	var projectReq models.CreateProject
	if err := c.BindJSON(&projectReq); err != nil {
		handleBussinessError(c, err.Error())
		return
	}

	// Get the user ID from the context
	userId := c.MustGet("userID").(string)
	userName := c.GetHeader("userName")

	// Set up roles
	ownerId := primitive.NewObjectID()
	memberId := primitive.NewObjectID()
	roles := []models.Role{
		{RoleId: ownerId, RoleName: "Owner"},
		{RoleId: memberId, RoleName: "Member"},
	}

	// Set up processes
	processes := setUpProcesses()

	// Set up members
	members := []models.Member{
		{UserId: userId, UserName: userName, RoleId: ownerId},
	}

	// set up invite request
	var invite = []models.Invite{}

	for _, v := range projectReq.InviteRequests {
		var roleId primitive.ObjectID
		if v.InviteRole == "Owner" {
			roleId = ownerId
		} else {
			roleId = memberId
		}
		invite = append(invite, models.Invite{
			InviteRoleId: roleId,
			InviteEmail:  v.InviteEmail,
			InviteRole:   v.InviteRole,
			InviteDate:   time.Now(),
		})
	}

	// Create a new project instance
	newProject := models.Project{
		ProjectProfile:   projectReq.ProjectProfile,
		ProjectStartDate: projectReq.ProjectStartDate,
		ProjectEndDate:   projectReq.ProjectEndDate,
		Invite:           invite,
		Views:            projectReq.Views,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
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

	handleSuccess(c, http.StatusOK, SUCCESS, CREATE_PROJECT_SUCCESS, nil)
}

func setUpProcesses() []models.Process {
	processStr := []string{PROCESS_DEFAULT_TODO, PROCESS_DEFAULT_IN_PROGRESS, PROCESS_DEFAULT_DONE}
	processes := make([]models.Process, len(processStr))
	for i, v := range processStr {
		processes[i] = models.Process{
			ProcessId:   primitive.NewObjectID(),
			ProcessName: v,
		}
	}
	return processes
}
