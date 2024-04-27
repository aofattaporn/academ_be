package services

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var mongoClient *mongo.Client

const (
	// Collection Name
	USER_COLLECTION         string = "User"
	PROJECT_COLLECTION      string = "Project"
	PERMISSION_COLLECTION   string = "Permission"
	TASKS_COLLECTION        string = "Tasks"
	CLASS_COLLECTION        string = "Class"
	NOTIFICATION_COLLECTION string = "Notification"
)
