package routes

import (
	"academ_be/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/api/v1/sign-up", controllers.SignUpUser())
	router.POST("/api/v1/sign-in", controllers.SignInUser())
	router.POST("/api/v1/sign-in/google", controllers.SignInWithGoogle())
}
