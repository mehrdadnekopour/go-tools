package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mehrdadnekopour/go-tools/mypes"
)

// Role ...
type Role struct {
	mypes.Model
	Name        string    `json:"name" validate:"required"`
	Code        RolesEnum `json:"code" validate:"required"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description"`
	Policies    *Policies `json:"policies"`

	// Assignments *Users    `json:"users" gorm:"many2many:user_roles;"`
	// Tags
}

// BeforeCreate ...
func (r *Role) BeforeCreate(scope *gorm.Scope) error {

	scope.SetColumn("Title", r.Code.Title())

	return nil
}

// Roles ...
type Roles []*Role

// Contains ...
func (rs Roles) Contains(role RolesEnum) bool {
	c := len(rs)
	for i := 0; i < c; i++ {
		r := rs[i]
		if r.Code == role {
			return true
		}
	}
	return false
}

// List ...
func (rs Roles) List() []int {
	var roleNos []int
	for _, r := range rs {
		roleNos = append(roleNos, int(r.Code))
	}

	return roleNos
}

// RolesEnum ...
type RolesEnum int

const (
	// RoleGuest ...
	RoleGuest RolesEnum = 0
	// RoleNormalUser ...
	RoleNormalUser RolesEnum = 1
	// RoleSuperUser ... me
	RoleSuperUser RolesEnum = 10
	// RoleAdmin ... customer > ISQI
	RoleAdmin RolesEnum = 100
	// RoleCompany ...
	RoleCompany RolesEnum = 500
	// RoleCompanySale ...
	RoleCompanySale RolesEnum = 501
	// RoleCompanyAfterSale ...
	RoleCompanyAfterSale RolesEnum = 502
	// RoleCompanyIndustrialMachinraies ...
	RoleCompanyIndustrialMachinraies RolesEnum = 503
	// RoleCompanyMotorCycles ...
	RoleCompanyMotorCycles RolesEnum = 504
	// RoleRegionalOffice ...
	RoleRegionalOffice RolesEnum = 600
	// RoleAgency ...
	RoleAgency RolesEnum = 700
	// RoleInspector ...
	RoleInspector RolesEnum = 900
	// RoleInspectorSale ...
	RoleInspectorSale RolesEnum = 901
	// RoleInspectorAfterSale ...
	RoleInspectorAfterSale RolesEnum = 902
	// RoleAPICall ...
	RoleAPICall RolesEnum = 1000
)

// Title ...
func (r RolesEnum) Title() string {
	switch r {
	case RoleGuest:
		return "guest"
	case RoleNormalUser:
		return "normal_user"
	case RoleSuperUser:
		return "super_user"
	case RoleAdmin:
		return "admin"
	case RoleCompany:
		return "company"
	case RoleCompanySale:
		return "company"
	case RoleCompanyAfterSale:
		return "company"
	case RoleCompanyIndustrialMachinraies:
		return "company"
	case RoleCompanyMotorCycles:
		return "company"
	case RoleRegionalOffice:
		return "regional_office"
	case RoleAgency:
		return "agency"
	case RoleInspector:
		return "inspector"
	case RoleInspectorSale:
		return "inspector"
	case RoleInspectorAfterSale:
		return "inspector"
	case RoleAPICall:
		return "api"
	}

	return ""
}
