package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `bson:"_id"`
	Email    string             `json:"email,omitempty" validate:"required"`
	FullName string             `json:"fullName,omitempty" validate:"required"`
}
