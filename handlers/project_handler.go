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
	var projectReq models.ProjectReq
	if err := c.BindJSON(&projectReq); err != nil {
		handleBussinessError(c, err.Error())
		return
	}

	// Get the user ID from the context
	userID := c.MustGet("userID").(string)
	userName := c.GetHeader("userName")

	// Set up roles
	ownerID := primitive.NewObjectID()
	memberID := primitive.NewObjectID()
	roles := []models.Role{
		{RoleID: ownerID, RoleName: "Owner"},
		{RoleID: memberID, RoleName: "Member"},
	}

	// Set up processes
	processes := setUpProcesses()

	// Set up members
	members := []models.Member{
		{UserID: userID, UserName: userName, RoleID: ownerID},
	}

	// set up invite request
	var invite = []models.Invite{}

	for _, v := range projectReq.InviteRequests {
		var roleId primitive.ObjectID
		if v.InviteRole == "Owner" {
			roleId = ownerID
		} else {
			roleId = memberID
		}
		invite = append(invite, models.Invite{
			InviteRoleID: roleId,
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
	processStr := []string{"Todo", "Inprogress", "Done"}
	processes := make([]models.Process, len(processStr))
	for i, v := range processStr {
		processes[i] = models.Process{
			ProcessID:   primitive.NewObjectID(),
			ProcessName: v,
		}
	}
	return processes
}
