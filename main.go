package main

import (
	"academ_be/configs"
	"academ_be/middlewares"
	"academ_be/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//run database
	configs.ConnectDB()
	configs.ConnectFirebase()
	router.Use(configs.CORSMiddleware())

	// create auth middleware
	middleware := router.Group("/")

	//routes
	middleware.Use(middlewares.AuthRequired())
	{
		routes.UserRoute(router)
	}

	router.Run("127.0.0.1:8080")
}
