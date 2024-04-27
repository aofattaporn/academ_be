package models

import "time"

type Notification struct {
	ProjectProfile ProjectProfile `bson:"projectProfile" json:"projectProfile"`
	Title          string         `bson:"title" json:"title"`
	Body           string         `bson:"Body" json:"Body"`
	Date           *time.Time     `bson:"date" json:"date"`
}
