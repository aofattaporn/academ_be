package routes

import (
	"academ_be/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/api/v1/sign-in", controllers.SignUpUser())
}
