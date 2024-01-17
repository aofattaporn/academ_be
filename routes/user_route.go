package routes

import (
	"academ_be/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/api/resource", handlers.CreateResource)
}
