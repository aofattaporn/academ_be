package services

import (
	"academ_be/configs"
	"academ_be/models"
	"context"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllTasksByProjectId(c *gin.Context, projectId string) (tasks []models.Tasks, err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	cursor, err := configs.GetCollection(mongoClient, TASKS_COLLECTION).Find(ctx, bson.M{"projectId": projectId})
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

func GetTasksByProjectId(c *gin.Context, projectId string) (tasks *models.Tasks, err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return nil, err
	}

	err = configs.GetCollection(mongoClient, TASKS_COLLECTION).FindOne(ctx, bson.M{"_id": objID}).Decode(&tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil

}

func CreateTasks(c *gin.Context, newTasks *models.CreateTasks) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	_, err = configs.GetCollection(mongoClient, TASKS_COLLECTION).InsertOne(ctx, &newTasks)
	if err != nil {
		return err
	}

	return nil

}

func DeleteTasksByTasksId(c *gin.Context, tasksId string) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(tasksId)
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "_id", Value: objID}}
	_, err = configs.GetCollection(mongoClient, TASKS_COLLECTION).DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil

}

func DeleteTasksByProjectId(c *gin.Context, projectId string) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	filter := bson.D{{Key: "projectId", Value: projectId}}
	_, err = configs.GetCollection(mongoClient, TASKS_COLLECTION).DeleteMany(ctx, filter)
	if err != nil {
		return err
	}

	return nil

}

func ChangeProcesss(c *gin.Context, tasksId string, processId string) (err error) {

	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	id, _ := primitive.ObjectIDFromHex(tasksId)
	filter := bson.D{{Key: "_id", Value: id}}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "processId", Value: processId}}}}

	_, err = configs.GetCollection(mongoClient, TASKS_COLLECTION).UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil

}

func UpdateTasksByTaskId(c *gin.Context, tasksId string, tasks models.UpdateTasks) error {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	// Convert the string task ID to an ObjectID
	id, err := primitive.ObjectIDFromHex(tasksId)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": tasks}

	result, err := configs.GetCollection(mongoClient, TASKS_COLLECTION).UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return errors.New("task not found")
	}

	return nil
}

func GetTasksByUserId(c *gin.Context, userId string) (projects []models.Tasks, err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	filter := bson.M{"assignee": bson.M{"$elemMatch": bson.M{"userId": userId}}}
	cursor, err := configs.GetCollection(mongoClient, PROJECT_COLLECTION).Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Decode projects directly into the result slice
	if err := cursor.All(ctx, &projects); err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return []models.Tasks{}, nil
	}

	return projects, nil
}
