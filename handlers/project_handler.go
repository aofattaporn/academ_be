package handlers

import (
	"academ_be/models"
	"academ_be/services"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ?Routes related to project details

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
	projectId := c.Param("projectId")
	if projectId == "" {
		handleBussinessError(c, "Can't to find your Tasks ID")
	}

	// TODO: GET project type *models.ProjectInfo if not found get from service and set project Id
	var project *models.ProjectInfo

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
		NavigateView:      project.Views[0],
		ProjectInfo:       *project,
		TaskPermission:    permission.Task,
		ProjectPermission: permission.Project,
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

func GetAllMyProjectsHomePage(c *gin.Context) {
	// Extract the user_id from the request parameters
	userID := c.MustGet(USER_ID).(string)

	// Call your business logic function to get projects by user ID
	projects, err := services.GetProjectsByMemberUserID(c, userID)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_PROJECT_SUCCESS, projects)
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

	// mapping save data on database
	userID := c.MustGet(USER_ID).(string)
	// find user in database from header
	user, err := services.FindUserOneById(c, userID)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

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
		{UserId: userID, UserName: user.FullName, Emaill: user.Email, RoleId: ownerId, AvatarColor: user.AvatarColor},
	}
	now := time.Now()

	base_view := []string{"List", "Board", "Calendar", "TimeLine"}
	filteredViews := make([]string, 0)

	// Add views from projectViews in the order they appear in views
	for _, view := range base_view {
		for _, pv := range createProject.Views {
			if pv == view {
				filteredViews = append(filteredViews, view)
				break
			}
		}
	}

	// Create a new project instance
	projectId := primitive.NewObjectID()
	newProject := models.Project{
		ProjectId: projectId,
		ProjectProfile: models.ProjectProfile{
			ProjectName: createProject.ProjectName,
			AvatarColor: getRandomColor(),
		},
		ClassName:        createProject.ClassName,
		ProjectStartDate: createProject.ProjectStartDate,
		ProjectEndDate:   createProject.ProjectEndDate,
		Views:            filteredViews,
		Invites:          []models.Invite{},
		CreatedAt:        &now,
		UpdatedAt:        &now,
		Process:          processes,
		Members:          members,
		Roles:            roles,
		IsArchive:        false,
	}

	// Create the project in the database
	err = services.CreateProject(c, &newProject)
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
	processColor := []string{"#C2C2C2", "#F9E116", "#72C554"}
	processes := make([]models.Process, len(processStr))
	for i, v := range processStr {
		processes[i] = models.Process{
			ProcessId:    primitive.NewObjectID(),
			ProcessName:  v,
			ProcessColor: processColor[i],
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
			Archive:     flag,
			Delete:      flag,
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

	userID := c.MustGet(USER_ID).(string)

	projectId := c.Param("projectId")
	if projectId == "" {
		handleBussinessError(c, "Can't to find your Tasks ID")
	}

	projectCacheKey := "projectDetails:" + projectId
	cachedProject, err := redisClient.Get(projectCacheKey).Result()

	var projectDetailsPermission models.ProjectDetailsPermission
	if err == redis.Nil {
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

		projectDetails := models.ProjectDetails{
			ProjectId:        project.ProjectId,
			ProjectProfile:   project.ProjectProfile,
			ClassName:        project.ClassName,
			Views:            project.Views,
			ProjectStartDate: project.ProjectStartDate,
			ProjectEndDate:   project.ProjectEndDate,
		}

		projectDetailsPermission = models.ProjectDetailsPermission{
			ProjectDetails:    projectDetails,
			ProjectPermission: permission.Project,
		}

	} else if err != nil {
		handleTechnicalError(c, err.Error())
		return
	} else {
		err = json.Unmarshal([]byte(cachedProject), &projectDetailsPermission)
		if err != nil {
			handleTechnicalError(c, err.Error())
			return
		}
	}

	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_PROJECT_SUCCESS, projectDetailsPermission)

}

func UpdateProjectDetails(c *gin.Context) {

	userID := c.MustGet(USER_ID).(string)
	projectId := c.Param("projectId")
	if projectId == "" {
		handleBussinessError(c, "Can't to find your Tasks ID")
	}

	var projectUpdate models.ProjectDetails
	if err := c.BindJSON(&projectUpdate); err != nil {
		handleBussinessError(c, err.Error())
		return
	}

	err := services.UpdateProjectDetails(c, projectId, projectUpdate)
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
		NavigateView:      project.Views[0],
		ProjectInfo:       *project,
		TaskPermission:    permission.Task,
		ProjectPermission: permission.Project,
	}

	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_PROJECT_SUCCESS, projectInfo)

	projectCacheKey := "projectDetails:" + projectId
	_, err = redisClient.Del(projectCacheKey).Result()
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

}

func DeleteProjectById(c *gin.Context) {

	projectId := c.Param("projectId")
	if projectId == "" {
		handleBussinessError(c, "Can't to find your Tasks ID")
	}

	project, err := services.DeleteProjectById(c, projectId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	err = services.DeleteTasksByProjectId(c, projectId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	for _, r := range project.Roles {
		err = services.DeletePermissionBy(c, r.PermissionId)
		if err != nil {
			handleTechnicalError(c, err.Error())
			return
		}
	}

	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_PROJECT_SUCCESS, nil)

}

func getPermissionIdByUser(c *gin.Context, project *models.ProjectInfo, userID string) (permission *models.Permission, err error) {

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
	permission, err = services.GetPermission(c, permissionID)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	return permission, nil

}

func ArchiveProjectById(c *gin.Context) {

	userID := c.MustGet(USER_ID).(string)
	projectId := c.Param("projectId")
	if projectId == "" {
		handleBussinessError(c, "Can't to find your Tasks ID")
	}

	var projectArchive models.ProjectArchive
	if err := c.BindJSON(&projectArchive); err != nil {
		handleBussinessError(c, err.Error())
		return
	}

	err := services.UpdateProjecArchive(c, projectId, projectArchive)
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
		NavigateView:      project.Views[0],
		ProjectInfo:       *project,
		TaskPermission:    permission.Task,
		ProjectPermission: permission.Project,
	}

	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_PROJECT_SUCCESS, projectInfo)

}
