package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjectReq struct {
	ProjectProfile   ProjectProfile  `bson:"projectProfile"`
	ProjectStartDate time.Time       `bson:"projectStartDate"`
	ProjectEndDate   time.Time       `bson:"projectEndDate"`
	Views            []string        `bson:"views"`
	InviteRequests   []InviteRequest `bson:"inviteRequests"`
}

type Project struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	ProjectProfile   ProjectProfile     `bson:"projectProfile"`
	ProjectStartDate time.Time          `bson:"projectStartDate"`
	ProjectEndDate   time.Time          `bson:"projectEndDate"`
	CreatedAt        time.Time          `bson:"createdAt"`
	UpdatedAt        time.Time          `bson:"updatedAt"`
	Process          []Process          `bson:"process"`
	Members          []Member           `bson:"members"`
	Views            []string           `bson:"views"`
	Roles            []Role             `bson:"roles"`
	InviteRequests   []InviteRequest    `bson:"inviteRequests"`
}

type Process struct {
	ProcessID   primitive.ObjectID `bson:"process_id"`
	ProcessName string             `bson:"process_name"`
}

type Member struct {
	UserID   string             `bson:"user_id"`
	UserName string             `bson:"user_name"`
	RoleID   primitive.ObjectID `bson:"role_id"`
}

type Role struct {
	RoleID   primitive.ObjectID `bson:"role_id"`
	RoleName string             `bson:"role_name"`

	// TODO : add permissions
}

type InviteRequest struct {
	InviteID     primitive.ObjectID `bson:"invite_id"`
	InviteRoleID string             `bson:"invite_role_id"`
	InviteRole   string             `bson:"invite_role"`
	InviteDate   time.Time          `bson:"invite_date"`
	InviteEmail  string             `bson:"invite_email"`
}

type ProjectProfile struct {
	ProjectName string `bson:"projectName"`
	AvatarColor string `bson:"avatarColor"`
}
