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

	// Create a new project instance
	newProject := models.Project{
		ProjectProfile: models.ProjectProfile{
			ProjectName: createProject.ProjectName,
			AvatarColor: getRandomColor(),
		},
		ProjectStartDate: time.Now(),
		ProjectEndDate:   createProject.ProjectEndDate,
		Views:            createProject.Views,
		Invite:           []models.Invite{},
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

	// Create a new project instance
	newProjectRes := models.MyProject{
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
			ProcessId:   primitive.NewObjectID(),
			ProcessName: v,
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
	services.CreatePermission(c, &permission)

	return ownerPermissionID
}

func getRandomColor() string {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	randomIndex := random.Intn(len(DEFULT_COLORS))
	return DEFULT_COLORS[randomIndex]
}
