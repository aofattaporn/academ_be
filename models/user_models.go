package models

import "time"

// *********** For User Collections ***************

type User struct {
	Id          string    `bson:"_id" json:"id,omitempty"`
	Email       string    `json:"email,omitempty" validate:"required"`
	FullName    string    `json:"fullName,omitempty" validate:"required"`
	AvatarColor string    `bson:"avatarColor" json:"avatarColor"`
	CreatedAt   time.Time `bson:"createdAt"`
	UpdatedAt   time.Time `bson:"updatedAt"`
}

type UserResponse struct {
	Email       string `json:"email,omitempty" validate:"required"`
	FullName    string `json:"fullName,omitempty" validate:"required"`
	AvatarColor string `bson:"avatarColor" json:"avatarColor"`
}

type FCM struct {
	FCM_TOKEN string `bson:"fcm_token" json:"fcm_token"`
}

type UserFullInfo struct {
	Id          string `bson:"_id" json:"id,omitempty" json:id`
	Email       string `json:"email,omitempty" validate:"required"`
	FullName    string `json:"fullName,omitempty" validate:"required"`
	AvatarColor string `bson:"avatarColor" json:"avatarColor"`
}
