package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// *********** For Project Collections ***************

// --------------- Project Models --------------------
// ---------------------------------------------------

type Project struct {
	ProjectId        primitive.ObjectID `bson:"_id,omitempty"`
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

type ProjectProfile struct {
	ProjectName string `bson:"projectName" json:"projectName"`
	AvatarColor string `bson:"avatarColor" json:"avatarColor"`
}

type Process struct {
	ProcessId   primitive.ObjectID `bson:"processId"`
	ProcessName string             `bson:"processName"`
}

type Member struct {
	UserId   string             `bson:"userId" json:"userId"`
	UserName string             `bson:"userName" json:"userName"`
	RoleId   primitive.ObjectID `bson:"roleId" json:"roleId"`
}

type Role struct {
	RoleId       primitive.ObjectID `bson:"roleId"`
	RoleName     string             `bson:"roleName"`
	PermissionId primitive.ObjectID `bson:"permissionsId"`
}

type Invite struct {
	InviteRoleId primitive.ObjectID `bson:"inviteRoleId"`
	InviteRole   string             `bson:"inviteRole"`
	InviteDate   time.Time          `bson:"inviteDate"`
	InviteEmail  string             `bson:"inviteEmail"`
}

// --------------- Create Project Models --------------------
// ----------------------------------------------------------

type CreateProject struct {
	ProjectName    string    `bson:"projectName"`
	ProjectEndDate time.Time `bson:"projectEndDate"`
	Views          []string  `bson:"views"`
}

type CreateInvite struct {
	InviteRole  string `bson:"inviteRole"`
	InviteEmail string `bson:"inviteEmail"`
}

// --------------- GET Project Models --------------------
// ----------------------------------------------------------

type MyProject struct {
	ProjectId        primitive.ObjectID `bson:"_id,omitempty" json:"projectId"`
	ProjectProfile   ProjectProfile     `bson:"projectProfile" json:"projectProfile"`
	Members          []Member           `bson:"members" json:"members"`
	ProjectStartDate time.Time          `bson:"projectStartDate" json:"projectStartDate"`
	ProjectEndDate   time.Time          `bson:"projectEndDate" json:"projectEndDate"`
}
