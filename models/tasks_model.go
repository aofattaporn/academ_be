package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// tasksId: string;
// tasksName: string;
// processId: string;
// assignee: string;
// startDate: Moment;
// dueDate: Moment;

type Tasks struct {
	TasksId   primitive.ObjectID `bson:"_id,omitempty" json:"tasksId"`
	ProjectId string             `bson:"projectId" json:"projectId"`
	TasksName string             `bson:"tasksName" json:"tasksName"`
	ProcessId string             `bson:"processId" json:"processId"`
	Assignee  string             `bson:"assignee" json:"assignee"`
	StartDate *time.Time         `bson:"startDate" json:"startDate"`
	DueDate   *time.Time         `bson:"dueDate" json:"dueDate"`
}

type UpdateTasks struct {
	TasksId   primitive.ObjectID `bson:"_id,omitempty" json:"tasksId"`
	TasksName string             `bson:"tasksName" json:"tasksName"`
	ProcessId string             `bson:"processId" json:"processId"`
	Assignee  string             `bson:"assignee" json:"assignee"`
	StartDate *time.Time         `bson:"startDate" json:"startDate"`
	DueDate   *time.Time         `bson:"dueDate" json:"dueDate"`
}

type CreateTasks struct {
	TasksId   primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	ProjectId string             `bson:"projectId" json:"projectId"`
	TasksName string             `bson:"tasksName" json:"tasksName"`
	ProcessId string             `bson:"processId" json:"processId"`
}

type AllMyTasks struct {
	Projects []AllTasksMyProject `bson:"projects" json:"projects"`
	Tasks    []Tasks             `bson:"tasks" json:"tasks"`
}
