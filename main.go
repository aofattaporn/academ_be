// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
package main

import (
	"academ_be/configs"
	"academ_be/handlers"
	"academ_be/middlewares"

	"github.com/gin-gonic/gin"

	_ "academ_be/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupRouter() *gin.Engine {

	router := gin.Default()

	// Run database
	configs.ConnectDB()
	admin := configs.ConnectFirebase()

	// Middlewares
	router.Use(middlewares.CORSMiddleware())

	v1 := router.Group("/api/v1")
	{
		v1.Use(middlewares.AuthRequire(admin))

		users := v1.Group("/")
		{
			users.GET("/users", handlers.GetUser)
			users.POST("/sign-up", handlers.CreateUser)
			users.POST("/sign-in", handlers.GetUser)
			users.POST("/sign-in/google", handlers.CreateUserByGoogle)
		}

		projects := v1.Group("/projects")
		{
			projects.GET("/users/id", handlers.GetAllMyProjects)
			projects.POST("", handlers.CreateProject)

			// Routes related to project details
			projects.GET("/:projectId", handlers.GetProjectById)
			projects.DELETE(":projectId", handlers.DeleteProjectById)
			projects.GET("/:projectId/details", handlers.GetProjectDetails)
			projects.PUT("/:projectId/details", handlers.UpdateProjectDetails)

			// Routes related to project invites
			projects.POST("/:projectId/invites", handlers.InviteNewMember)
			projects.DELETE("/:projectId/invites/:inviteId", handlers.DeleteInviteMember)
			projects.GET("/invites/token/:token", handlers.AcceptInviteMember)

			// Routes related to project roles and permissions
			projects.GET("/:projectId/roleAndPermission", handlers.GetProjectRoleAndPermissions)
			projects.POST("/:projectId/roleAndPermission", handlers.CreateProjectRoleAndPermissions)
			projects.PUT("/:projectId/roles/:roleId", handlers.UpdateRoleName)
			projects.DELETE("/:projectId/roles/:roleId", handlers.DeleteRole)
			projects.PUT("/:projectId/permissions/:permissionId", handlers.UpdatePermission)

			// Routes related to project members
			projects.GET("/:projectId/members", handlers.GetProjectMembers)
			projects.DELETE("/:projectId/members/:memberId", handlers.RemoveMember)
			projects.GET("/:projectId/members/:memberId/roles/:roleId", handlers.ChangeRoleMember)
		}

		tasks := v1.Group("/tasks")
		{
			tasks.POST("", handlers.CreateTasks)               // Create a new task
			tasks.GET("/:taskId", handlers.GetTasksById)       // Get task by ID
			tasks.PUT("/:taskId", handlers.UpdateTasks)        // Update task by ID
			tasks.DELETE("/:taskId", handlers.DeleteTasksById) // Delete task by ID

			// Routes related to tasks and projects
			tasks.GET("projects/:projectId", handlers.GetAllTasksByProjectId) // Get all tasks by project ID

			// Routes related to task processes
			tasks.PUT("/:taskId/process/:processId", handlers.ChangeProcesss) // Change task process
		}

	}

	return router
}

func main() {
	router := setupRouter()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run("0.0.0.0:8080")
}
