package main

import (
	"academ_be/configs"
	"academ_be/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//run database
	configs.ConnectDB()
	configs.ConnectFirebase()
	router.Use(configs.CORSMiddleware())

	//routes
	routes.UserRoute(router)

	router.Run("127.0.0.1:8080")
}
