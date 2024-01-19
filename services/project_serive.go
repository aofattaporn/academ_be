package services

import (
	"academ_be/configs"
	"academ_be/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	mongoClient = configs.ConnectDB()
}

func CreateProject(c *gin.Context, newProject *models.Project) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	_, err := configs.GetCollection(mongoClient, "Project").InsertOne(ctx, newProject)
	if err != nil {
		handleError(c, err, http.StatusUnauthorized, ERROR)
	}
}
