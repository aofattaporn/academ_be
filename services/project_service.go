package services

import (
	"academ_be/configs"
	"academ_be/models"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateProject(c *gin.Context, newUser *models.Project) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	_, err = configs.GetCollection(mongoClient, "Project").InsertOne(ctx, newUser)
	if err != nil {
		return err
	}

	return nil
}

func GetProjectsByMemberUserID(c *gin.Context, myUserID string) (project []models.Project, err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	// Define filter to find projects where Members.UserID equals myUserID
	filter := bson.M{"members.userID": myUserID}

	// Define options to specify fields to include
	projection := bson.M{"_id": 1, "projectProfile": 1}
	opts := options.Find().SetProjection(projection)

	// Find projects matching the filter
	cursor, err := configs.GetCollection(mongoClient, "Project").Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var projects []models.Project
	if err := cursor.All(ctx, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}
