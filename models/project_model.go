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
	ProjectStartDate *time.Time         `bson:"projectStartDate"`
	ProjectEndDate   *time.Time         `bson:"projectEndDate"`
	CreatedAt        *time.Time         `bson:"createdAt"`
	UpdatedAt        *time.Time         `bson:"updatedAt"`
	Process          []Process          `bson:"process"`
	Members          []Member           `bson:"members"`
	Views            []string           `bson:"views"`
	Roles            []Role             `bson:"roles"`
	Invite           []Invite           `bson:"invite"`
}

type ProjectInfoPermission struct {
	ProjectInfo    ProjectInfo    `bson:"projectInfo" json:"projectInfo"`
	TaskPermission TaskPermission `bson:"taskPermission" json:"taskPermission"`
}

type ProjectInfo struct {
	ProjectId      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ProjectProfile ProjectProfile     `bson:"projectProfile"  json:"projectProfile"`
	Process        []Process          `bson:"process"  json:"process"`
	Members        []Member           `bson:"members"  json:"members"`
	Roles          []Role             `bson:"roles"  json:"roles"`
	Views          []string           `bson:"views"  json:"views"`
}

type ProjectDetails struct {
	ProjectId        primitive.ObjectID `bson:"_id,omitempty" json:"projectId"`
	ProjectProfile   ProjectProfile     `bson:"projectProfile" json:"projectProfile"`
	Views            []string           `bson:"views" json:"views"`
	ProjectStartDate *time.Time         `bson:"projectStartDate" json:"startDate"`
	ProjectEndDate   *time.Time         `bson:"projectEndDate" json:"dueDate"`
}

type ProjectProfile struct {
	ProjectName string `bson:"projectName" json:"projectName"`
	AvatarColor string `bson:"avatarColor" json:"avatarColor"`
}

type Process struct {
	ProcessId    primitive.ObjectID `bson:"processId" json:"processId"`
	ProcessName  string             `bson:"processName" json:"processName"`
	ProcessColor string             `bson:"processColor" json:"processColor"`
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
	ProjectName    string     `bson:"projectName"`
	ProjectEndDate *time.Time `bson:"projectEndDate"`
	Views          []string   `bson:"views"`
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
	ProjectStartDate *time.Time         `bson:"projectStartDate" json:"projectStartDate"`
	ProjectEndDate   *time.Time         `bson:"projectEndDate" json:"projectEndDate"`
}
