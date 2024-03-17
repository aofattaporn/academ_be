package services

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var mongoClient *mongo.Client

const (
	// Collection Name
	USER_COLLECTION       string = "User"
	PROJECT_COLLECTION    string = "Project"
	PERMISSION_COLLECTION string = "Permission"
	TASK_COLLECTION       string = "Task"
	CLASS_COLLECTION      string = "Class"
)
