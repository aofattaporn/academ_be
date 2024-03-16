package services

import (
	"academ_be/configs"
	"academ_be/models"
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreatePermission(c *gin.Context, permission *models.Permission, permissionId primitive.ObjectID) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	permission.Id = permissionId
	_, err = configs.GetCollection(mongoClient, PERMISSION_COLLECTION).InsertOne(ctx, permission)
	if err != nil {
		return err
	}

	return nil
}

func GetPermission(c *gin.Context, permissionId primitive.ObjectID) (permission *models.Permission, err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	err = configs.GetCollection(mongoClient, PERMISSION_COLLECTION).FindOne(ctx, bson.M{"_id": permissionId}).Decode(&permission)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("permission not found")
		}
		return nil, fmt.Errorf("error decoding permission: %v", err)
	}

	return permission, nil
}
