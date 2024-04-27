package models

import "time"

type Notification struct {
	ProjectProfile ProjectProfile `bson:"projectProfile" json:"projectProfile"`
	UserId         string         `bson:"userId" json:"userId"`
	Title          string         `bson:"title" json:"title"`
	Body           string         `bson:"Body" json:"Body"`
	Date           *time.Time     `bson:"date" json:"date"`
	IsClear        bool           `bson:"isClear" json:"isClear"`
}

type NotificationRes struct {
	Id             string         `bson:"_id" json:"id,omitempty"`
	ProjectProfile ProjectProfile `bson:"projectProfile" json:"projectProfile"`
	UserId         string         `bson:"userId" json:"userId"`
	Title          string         `bson:"title" json:"title"`
	Body           string         `bson:"Body" json:"Body"`
	Date           *time.Time     `bson:"date" json:"date"`
	IsClear        bool           `bson:"isClear" json:"isClear"`
}
