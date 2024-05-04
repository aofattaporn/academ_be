package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// *********** For Permissions Collections ***************

// -------------------- Permissions -------------------------
// ----------------------------------------------------------

type Permission struct {
	Id      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Members MembersPermission  `bson:"members" json:"members"`
	Project ProjectPermission  `bson:"project" json:"project"`
	Task    TaskPermission     `bson:"task" json:"task"`
	Role    RolePermission     `bson:"role" json:"role"`
}

type MembersPermission struct {
	AddRole bool `bson:"addRole" json:"addRole"`
	Invite  bool `bson:"invite" json:"invite"`
	Remove  bool `bson:"remove" json:"remove"`
}

type ProjectPermission struct {
	EditProfile bool `bson:"editProfile" json:"editProfile"`
	Archive     bool `bson:"archive" json:"archive"`
	Delete      bool `bson:"delete" json:"delete"`
}

type TaskPermission struct {
	AddNew        bool `bson:"addNew" json:"addNew"`
	Delete        bool `bson:"delete" json:"delete"`
	Edit          bool `bson:"edit" json:"edit"`
	ManageProcess bool `bson:"manageProcess" json:"manageProcess"`
}

type RolePermission struct {
	AddNew bool `bson:"addNew" json:"addNew"`
	Edit   bool `bson:"edit" json:"edit"`
	Delete bool `bson:"delete" json:"delete"`
}

type CreateRole struct {
	NewRole string `bson:"newRole" json:"newRole"`
}

type RoleAndRolePermission struct {
	RolesAndFullPermission []RoleAndFullPermission `bson:"rolesAndFullPermission" json:"rolesAndFullPermission"`
	RolePermission         RolePermission          `bson:"rolePermission" json:"rolePermission"`
}
