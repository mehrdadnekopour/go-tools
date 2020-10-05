package models

import (
	"github.com/mehrdadnekopour/go-tools/mypes"
)

// "./mypes"

// User ...
type User struct {
	mypes.Model
	Title     string        `json:"title"`
	Username  string        `json:"username" validate:"required" gorm:"size:255;index:idx_uname,unique"`
	Password  string        `json:"password" validate:"required"`
	OwnerType UserOwnerType `json:"owner_type"`
	OwnerID   int           `json:"owner_id"`
	Roles     Roles         `json:"roles" gorm:"many2many:user_roles;"`
}

// UserRole ,... Jost for dropping relation table
type UserRole struct {
	UserID uint
	RoleID uint
}

// UserOwnerType ...
type UserOwnerType string

const (
	// OTNone ...
	OTNone UserOwnerType = ""
	// OTGuest ...
	OTGuest UserOwnerType = "guest"
	// OTNormal ...
	OTNormal UserOwnerType = "normal"
	// OTAdmin ...
	OTAdmin UserOwnerType = "admin"
	// OTISQI ...
	OTISQI UserOwnerType = "isqi"
	// OTCompany ...
	OTCompany UserOwnerType = "company"
	// OTAgency ...
	OTAgency UserOwnerType = "agency"
	// OTInspector ...
	OTInspector UserOwnerType = "inspector"
	// OTApi ...
	OTApi UserOwnerType = "api"
)

// Users ...
type Users []*User
