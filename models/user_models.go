package models

import "time"

// *********** For User Collections ***************

type User struct {
	Id        string    `bson:"_id" json:"id,omitempty"`
	Email     string    `json:"email,omitempty" validate:"required"`
	FullName  string    `json:"fullName,omitempty" validate:"required"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

type UserResponse struct {
	Email    string `json:"email,omitempty" validate:"required"`
	FullName string `json:"fullName,omitempty" validate:"required"`
}
