package routes

import (
	"academ_be/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/api/v1/users", handlers.GetUser)
	router.POST("/api/v1/sign-up", handlers.CreateUser)
	router.POST("/api/v1/sign-in", handlers.GetUser)
	router.POST("/api/v1/sign-in/google", handlers.CreateUserByGoogle)
}
