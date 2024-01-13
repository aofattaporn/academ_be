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

	//run database
	configs.ConnectDB()
	configs.ConnectFirebase()

	router.Use(configs.CORSMiddleware())
	router.Use(middlewares.AuthRequire())

	//routes
	routes.UserRoute(router)

	router.Run("127.0.0.1:8080")
}
