package handlers

import (
	"academ_be/models"
	"academ_be/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProject(c *gin.Context) {

	// map request body
	var project models.Project
	if err := c.BindJSON(&project); err != nil {
		handleBadRequest(c, ERROR, err.Error())
		return
	}

	userID := c.MustGet("userID").(string)
	data := *services.FindUserOneById(c, userID)

	var owner models.Member
	var members []models.Member
	owner.ID = userID
	owner.FullName = data.FullName
	owner.Email = data.Email
	owner.Roles = "Owner"
	members = append(members, owner)

	// Create a new project
	project.ProjectStartDate, _ = time.Parse(time.RFC3339, project.ProjectStartDate.Format(time.RFC3339))
	project.ProjectEndDate, _ = time.Parse(time.RFC3339, project.ProjectEndDate.Format(time.RFC3339))

	newProject := models.Project{
		ID:                 primitive.NewObjectID(),
		ProjectName:        project.ProjectName,
		ProjectStartDate:   project.ProjectStartDate,
		ProjectEndDate:     project.ProjectEndDate,
		Views:              project.Views,
		Members:            members,
		InvitationRequests: project.InvitationRequests,
	}

	// save new project on mongodb
	services.CreateProject(c, &newProject)

	handleSuccess(c, http.StatusCreated, SUCCESS, USER_SIGNUP_SUCCESS)
}

func UpdayeProjectName(c *gin.Context) {}

func UpdayeProjectDuration(c *gin.Context) {}

func UpdateProjectView(c *gin.Context) {}

func GetProjectRequest(c *gin.Context) {}
