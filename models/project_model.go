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
	ProjectName string `bson:"projectName"`
	AvatarColor string `bson:"avatarColor"`
}

type Process struct {
	ProcessId   primitive.ObjectID `bson:"processId"`
	ProcessName string             `bson:"processName"`
}

type Member struct {
	UserId   string             `bson:"userId"`
	UserName string             `bson:"userName"`
	RoleId   primitive.ObjectID `bson:"roleId"`
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
	ProjectProfile   ProjectProfile `bson:"projectProfile"`
	ProjectStartDate time.Time      `bson:"projectStartDate"`
	ProjectEndDate   time.Time      `bson:"projectEndDate"`
	Views            []string       `bson:"views"`
	InviteRequests   []CreateInvite `bson:"inviteRequests"`
}

type CreateInvite struct {
	InviteRole  string `bson:"inviteRole"`
	InviteEmail string `bson:"inviteEmail"`
}

// --------------- GET Project Models --------------------
// ----------------------------------------------------------

type MyProject struct {
	ProjectId        primitive.ObjectID `bson:"_id,omitempty"`
	ProjectProfile   ProjectProfile     `bson:"projectProfile"`
	Members          []Member           `bson:"members"`
	ProjectStartDate time.Time          `bson:"projectStartDate"`
	ProjectEndDate   time.Time          `bson:"projectEndDate"`
}
