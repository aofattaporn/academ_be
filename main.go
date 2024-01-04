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

	//routes
	routes.UserRoute(router)

	router.Run("localhost:6000")
}
