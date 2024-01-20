package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Project represents a project entity.
type Project struct {
	ID                 primitive.ObjectID  `bson:"_id" json:"id,omitempty"`
	ProjectName        string              `json:"projectName,omitempty" validate:"required"`
	CreateDate         time.Time           `json:"createeDate,omitempty" validate:"required"`
	ProjectStartDate   time.Time           `json:"projectStartDate,omitempty" validate:"required"`
	ProjectEndDate     time.Time           `json:"projectEndDate,omitempty" validate:"required"`
	Views              []View              `json:"views,omitempty" validate:"required"`
	Members            []Member            `json:"members,omitempty" validate:"required"`
	Roles              []Role              `json:"roles,omitempty"`
	InvitationRequests []InvitationRequest `json:"invitationRequest,omitempty"`
}

type View string

type Member struct {
	ID    string `bson:"_id" json:"id,omitempty"`
	Email string `bson:"email" json:"email,omitempty" validate:"required,email"`
	Roles Role   `bson:"roles" json:"roles,omitempty" validate:"required"`
}

type Role string

type InvitationRequest struct {
	Email      string    `json:"email"`
	Roles      Role      `json:"roles"`
	InviteDate time.Time `json:"inviteDate,omitempty" validate:"required"`
}

type Permossoions struct {
}
