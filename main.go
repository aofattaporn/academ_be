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

	//routes
	routes.UserRoute(router) //add this

	router.Run("localhost:6000")
}
