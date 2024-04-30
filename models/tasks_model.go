package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tasks struct {
	TasksId   primitive.ObjectID `bson:"_id,omitempty" json:"tasksId,omitempty"`
	ProjectId string             `bson:"projectId" json:"projectId"`
	ProcessId string             `bson:"processId" json:"processId"`
	TasksName string             `bson:"tasksName,omitempty" json:"tasksName"`
	Assignee  *Member            `bson:"assignee,omitempty" json:"assignee,omitempty"`
	StartDate *time.Time         `bson:"startDate,omitempty" json:"startDate,omitempty"`
	DueDate   *time.Time         `bson:"dueDate,omitempty" json:"dueDate,omitempty"`
}

type UpdateTasks struct {
	TasksId   primitive.ObjectID `bson:"_id,omitempty" json:"tasksId"`
	TasksName string             `bson:"tasksName" json:"tasksName"`
	ProcessId string             `bson:"processId" json:"processId"`
	Assignee  *Member            `bson:"assignee,omitempty" json:"assignee,omitempty"`
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
