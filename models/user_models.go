package models

type User struct {
	Id       string `bson:"_id"`
	Email    string `json:"email,omitempty" validate:"required"`
	FullName string `json:"fullName,omitempty" validate:"required"`
}
