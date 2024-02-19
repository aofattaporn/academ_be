package services

import (
	"academ_be/configs"
	"academ_be/models"
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func CreatePermission(c *gin.Context, permission *models.Permission) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	_, err = configs.GetCollection(mongoClient, PERMISSION_COLLECTION).InsertOne(ctx, permission)
	if err != nil {
		return err
	}

	return nil
}
