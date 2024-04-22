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

			projects.POST("/:projectsId/invites", handlers.InviteNewMember)
			projects.GET("/:projectsId/details", handlers.GetProjectDetails)
			projects.PUT("/:projectsId/details", handlers.UpdateProjectDetails)
			projects.GET("/:projectsId/roleAndPermission", handlers.GetProjectRoleAndPermissions)
			projects.POST("/:projectsId/roleAndPermission", handlers.CreateProjectRoleAndPermissions)
			projects.PUT("/:projectsId/roles/:roleId", handlers.UpdateRoleName)
			projects.DELETE("/:projectsId/roles/:roleId", handlers.DeleteRole)
			projects.PUT("/:projectsId/permissions/:permissionId", handlers.UpdatePermission)
			projects.GET("/:projectsId/members/:memberId/roles/:roleId", handlers.ChangeRoleMember)
			projects.GET("/:projectsId/members", handlers.GetProjectMembers)
			projects.POST("", handlers.CreateProject)
			projects.GET("/:projectsId", handlers.GetProjectById)
			projects.GET("/users/id", handlers.GetAllMyProjects)

		}

		tasks := v1.Group("/tasks")
		{
			tasks.POST("", handlers.CreateTasks)
			tasks.GET("projects/:projectId", handlers.GetAllTasksByProjectId)
			tasks.GET("/:tasksId", handlers.GetTasksById)
			tasks.PUT("/:tasksId/process/:processId", handlers.ChangeProcesss)
			tasks.PUT("/:tasksId", handlers.UpdateTasks)
			tasks.DELETE("/:tasksId", handlers.DeleteTasksById)
		}
	}

	return router
}

func main() {
	router := setupRouter()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run("0.0.0.0:8080")
}
