package models

type User struct {
	Id       string `bson:"_id" json:"id,omitempty"`
	Email    string `json:"email,omitempty" validate:"required"`
	FullName string `json:"fullName,omitempty" validate:"required"`
}

type UserResponse struct {
	Email    string `json:"email,omitempty" validate:"required"`
	FullName string `json:"fullName,omitempty" validate:"required"`
}
