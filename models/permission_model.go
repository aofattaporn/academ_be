package models

// *********** For Permissions Collections ***************

// -------------------- Permissions -------------------------
// ----------------------------------------------------------

type Permission struct {
	Members struct {
		AddRole bool `bson:"addRole"`
		Invite  bool `bson:"invite"`
		Remove  bool `bson:"remove"`
	} `bson:"members"`
	Project struct {
		EditProfile bool `bson:"editProfile"`
		ManageView  bool `bson:"manageView"`
	} `bson:"project"`
	Task struct {
		AddNew        bool `bson:"addNew"`
		Delete        bool `bson:"delete"`
		Edit          bool `bson:"edit"`
		ManageProcess bool `bson:"manageProcess"`
	} `bson:"task"`
	Role struct {
		AddNew bool `bson:"addNew"`
		Edit   bool `bson:"edit"`
		Delete bool `bson:"delete"`
	} `bson:"role"`
}
