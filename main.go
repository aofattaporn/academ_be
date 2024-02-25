// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
package main

import (
	"academ_be/configs"
	"academ_be/middlewares"
	"academ_be/routes"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var firebaseClient *auth.Client = configs.ConnectFirebase()

func main() {
	router := gin.Default()

	// Run database
	configs.ConnectDB()
	// admin := configs.ConnectFirebase()

	// Middlewares
	router.Use(middlewares.CORSMiddleware())
	// router.Use(middlewares.AuthRequire(admin))
	router.Use(middlewares.ErrorHandler())

	// Routes
	routes.UserRoute(router)
	routes.ProjectRoute(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run("0.0.0.0:8080")
}
