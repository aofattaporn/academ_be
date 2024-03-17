package services

import (
	"academ_be/configs"
	"academ_be/models"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllTasksByProjectId(c *gin.Context, projectId string) (tasks []models.Tasks, err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return nil, err
	}

	cursor, err := configs.GetCollection(mongoClient, TASKS_COLLECTION).Find(ctx, bson.M{"projectId": objID})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx) // Close the cursor when done

	// Initialize tasks as a pointer to a slice of models.Tasks
	tasks = make([]models.Tasks, 0)

	// Pass the pointer to tasks to the All method
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}

	if len(tasks) == 0 {
		return []models.Tasks{}, nil
	}

	return tasks, err
}
