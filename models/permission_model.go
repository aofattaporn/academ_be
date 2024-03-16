package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// *********** For Permissions Collections ***************

// -------------------- Permissions -------------------------
// ----------------------------------------------------------

type Permission struct {
	Id      primitive.ObjectID `bson:"_id,omitempty"`
	Members MembersPermission  `bson:"members"`
	Project ProjectPermission  `bson:"project"`
	Task    TaskPermission     `bson:"task"`
	Role    RolePermission     `bson:"role"`
}

type MembersPermission struct {
	AddRole bool `bson:"addRole"`
	Invite  bool `bson:"invite"`
	Remove  bool `bson:"remove"`
}

type ProjectPermission struct {
	EditProfile bool `bson:"editProfile"`
	ManageView  bool `bson:"manageView"`
}

type TaskPermission struct {
	AddNew        bool `bson:"addNew"`
	Delete        bool `bson:"delete"`
	Edit          bool `bson:"edit"`
	ManageProcess bool `bson:"manageProcess"`
}

type RolePermission struct {
	AddNew bool `bson:"addNew"`
	Edit   bool `bson:"edit"`
	Delete bool `bson:"delete"`
}
