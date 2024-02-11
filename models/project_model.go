package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InviteRequests struct {
	InviteRole  string `bson:"inviteRole"`
	InviteEmail string `bson:"inviteEmail"`
}

type Invite struct {
	InviteRoleID primitive.ObjectID `bson:"inviteRoleId"`
	InviteRole   string             `bson:"inviteRole"`
	InviteDate   time.Time          `bson:"inviteDate"`
	InviteEmail  string             `bson:"inviteEmail"`
}

type ProjectProfile struct {
	ProjectName string `bson:"projectName"`
	AvatarColor string `bson:"avatarColor"`
}

type ProjectReq struct {
	ProjectProfile   ProjectProfile   `bson:"projectProfile"`
	ProjectStartDate time.Time        `bson:"projectStartDate"`
	ProjectEndDate   time.Time        `bson:"projectEndDate"`
	Views            []string         `bson:"views"`
	InviteRequests   []InviteRequests `bson:"inviteRequests"`
}

type ListMyProjectRes struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	ProjectProfile ProjectProfile     `bson:"projectProfile"`
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
	Invite           []Invite           `bson:"invite"`
}

type Process struct {
	ProcessID   primitive.ObjectID `bson:"processId"`
	ProcessName string             `bson:"processName"`
}

type Member struct {
	UserID   string             `bson:"userId"`
	UserName string             `bson:"userName"`
	RoleID   primitive.ObjectID `bson:"roleId"`
}

type Role struct {
	RoleID   primitive.ObjectID `bson:"roleId"`
	RoleName string             `bson:"roleName"`

	// TODO : add permissions
}
