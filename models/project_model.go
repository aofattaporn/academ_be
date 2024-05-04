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
	ClassName        string             `bson:"className"`
	ProjectStartDate *time.Time         `bson:"projectStartDate"`
	ProjectEndDate   *time.Time         `bson:"projectEndDate"`
	CreatedAt        *time.Time         `bson:"createdAt"`
	UpdatedAt        *time.Time         `bson:"updatedAt"`
	Process          []Process          `bson:"process"`
	Members          []Member           `bson:"members"`
	Views            []string           `bson:"views"`
	Roles            []Role             `bson:"roles"`
	Invites          []Invite           `bson:"invites"  json:"invites"`
	IsArchive        bool               `bson:"isArchive"  json:"isArchive"`
}

type ProjectInfoPermission struct {
	ProjectInfo       ProjectInfo       `bson:"projectInfo" json:"projectInfo"`
	TaskPermission    TaskPermission    `bson:"taskPermission" json:"taskPermission"`
	ProjectPermission ProjectPermission `bson:"projectPermission" json:"projectPermission"`
}

type ProjectInfo struct {
	ProjectId        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ProjectProfile   ProjectProfile     `bson:"projectProfile"  json:"projectProfile"`
	ClassName        string             `bson:"className"  json:"className"`
	Process          []Process          `bson:"process"  json:"process"`
	Members          []Member           `bson:"members"  json:"members"`
	Roles            []Role             `bson:"roles"  json:"roles"`
	Views            []string           `bson:"views"  json:"views"`
	Invites          []Invite           `bson:"invites"  json:"invites"`
	ProjectStartDate *time.Time         `bson:"projectStartDate" json:"projectStartDate"`
	ProjectEndDate   *time.Time         `bson:"projectEndDate"  json:"projectEndDate"`
	IsArchive        bool               `bson:"isArchive"  json:"isArchive"`
}

type ProjectDetailsPermission struct {
	ProjectDetails    ProjectDetails    `bson:"projectDetails" json:"projectDetails"`
	ProjectPermission ProjectPermission `bson:"projectPermission" json:"projectPermission"`
}

type ProjectArchive struct {
	IsArchive bool `bson:"isArchive"  json:"isArchive"`
}

type ProjectDetails struct {
	ProjectId        primitive.ObjectID `bson:"_id,omitempty" json:"projectId"`
	ProjectProfile   ProjectProfile     `bson:"projectProfile" json:"projectProfile"`
	Views            []string           `bson:"views" json:"views"`
	ClassName        string             `bson:"className" json:"className"`
	ProjectStartDate *time.Time         `bson:"projectStartDate,omitempty" json:"projectStartDate,omitempty"`
	ProjectEndDate   *time.Time         `bson:"projectEndDate,omitempty" json:"projectEndDate,omitempty"`
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
	UserId      string             `bson:"userId" json:"userId"`
	UserName    string             `bson:"userName" json:"userName"`
	Emaill      string             `bson:"email" json:"email"`
	RoleId      primitive.ObjectID `bson:"roleId" json:"roleId"`
	AvatarColor string             `bson:"avatarColor" json:"avatarColor"`
}

type Role struct {
	RoleId       primitive.ObjectID `bson:"roleId" json:"roleId"`
	RoleName     string             `bson:"roleName" json:"roleName"`
	PermissionId primitive.ObjectID `bson:"permissionsId" json:"permissionId"`
}

type RoleAndFullPermission struct {
	RoleId     primitive.ObjectID `bson:"roleId" json:"roleId"`
	RoleName   string             `bson:"roleName" json:"roleName"`
	Permission Permission         `bson:"permissionsId" json:"permission"`
}

type Invite struct {
	InviteId     primitive.ObjectID `bson:"inviteId" json:"inviteId"`
	InviteRoleId string             `bson:"inviteRoleId" json:"inviteRoleId"`
	InviteDate   time.Time          `bson:"inviteDate" json:"inviteDate"`
	InviteEmail  string             `bson:"inviteEmail" json:"inviteEmail"`
	Token        string             `bson:"token" json:"token" `
}

type InviteReq struct {
	InviteRoleId string    `bson:"inviteRoleId" json:"inviteRoleId"`
	InviteDate   time.Time `bson:"inviteDate" json:"inviteDate"`
	InviteEmail  string    `bson:"inviteEmail" json:"inviteEmail"`
}

// --------------- Create Project Models --------------------
// ----------------------------------------------------------

type CreateProject struct {
	ProjectName      string     `bson:"projectName"`
	ProjectStartDate *time.Time `bson:"projectStartDate" json:"projectStartDate"`
	ProjectEndDate   *time.Time `bson:"projectEndDate"`
	ClassName        string     `bson:"className"`
	Views            []string   `bson:"views"`
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
	ClassName        string             `bson:"className" json:"className"`
	ProjectStartDate *time.Time         `bson:"projectStartDate" json:"projectStartDate"`
	ProjectEndDate   *time.Time         `bson:"projectEndDate" json:"projectEndDate"`
}

type AllMemberProject struct {
	Members []Member `bson:"members" json:"members"`
	Roles   []Role   `bson:"roles" json:"roles"`
	Invites []Invite `bson:"invites"  json:"invites"`
}

type AllMemberAndPermission struct {
	AllMemberProject  AllMemberProject  `bson:"allMemberProject" json:"allMemberProject"`
	MembersPermission MembersPermission `bson:"membersPermission" json:"membersPermission"`
}

type AllTasksMyProject struct {
	ProjectId        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ProjectProfile   ProjectProfile     `bson:"projectProfile" json:"projectProfile"`
	ProjectStartDate *time.Time         `bson:"projectStartDate" json:"projectStartDate"`
	ProjectEndDate   *time.Time         `bson:"projectEndDate" json:"projectEndDate"`
	ClassName        string             `bson:"className" json:"className"`
	CreatedAt        *time.Time         `bson:"createdAt" json:"createdAt"`
	UpdatedAt        *time.Time         `bson:"updatedAt" json:"updatedAt"`
	Process          []Process          `bson:"process" json:"process"`
}
