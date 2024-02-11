package routes

import (
	"academ_be/handlers"

	"github.com/gin-gonic/gin"
)

func ProjectRoute(router *gin.Engine) {
	router.GET("/api/v1/projects/users/id", handlers.GetAllMyProjects)
	router.POST("/api/v1/projects/users/id", handlers.CreateProject)
}
