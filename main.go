package main

import (
	"academ_be/configs"
	"academ_be/middlewares"
	"academ_be/routes"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

var firebaseClient *auth.Client = configs.ConnectFirebase()

func main() {
	router := gin.Default()

	// Run database
	configs.ConnectDB()
	admin := configs.ConnectFirebase()

	// Middlewares
	router.Use(middlewares.CORSMiddleware())
	router.Use(middlewares.AuthRequire(admin))
	router.Use(middlewares.ErrorHandler())

	// Routes
	routes.UserRoute(router)

	router.Run("127.0.0.1:8080")
}
