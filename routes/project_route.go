package routes

import (
	"academ_be/handlers"

	"github.com/gin-gonic/gin"
)

func ProjectRoute(router *gin.Engine) {
	router.POST("/api/v1/project", handlers.CreateProject)
}
