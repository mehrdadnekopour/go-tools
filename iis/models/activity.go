package models

import (
	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mehrdadnekopour/go-tools/mypes"
)

// Activity ...
type Activity struct {
	mypes.Model
	Action        Action         `json:"action"`
	Username      string         `json:"username"`
	UserTitle     string         `json:"user_title"`
	UserID        int            `json:"user_id"`
	UserOwnerType UserOwnerType  `json:"user_owner_type"`
	UserOwnerID   int            `json:"user_owner_id"`
	ObjectType    ObjectType     `json:"object_type"`
	Info          postgres.Jsonb `json:"info"`
}

// Activities ...
type Activities []*Activity

// InspectorActivityInfo ... this type saves as a json in activity table, info column
type InspectorActivityInfo struct {
	Inspector        *Inspector       `json:"inspector"`
	ObjectID         interface{}      `json:"object_id"` // can be 'int' or 'bson.ObjectId'(for results)
	InspectionDomain InspectionDomain `json:"inspection_domain"`
	CompanyID        int              `json:"company_id"`
	CompanyName      string           `json:"company_name"`
	Company          *Company         `json:"company"`
	AgencyCode       int              `json:"agency_code"`
	AgencyOwner      string           `json:"agency_owner"`
	Round            int              `json:"round"`
	Period           InspectionPeriod `json:"period"`
	Parts            InspectionParts  `json:"parts"`
}

// ActivityRequest ...
type ActivityRequest struct {
	Activity   Activity
	Inspection Inspection
}

// Action ...
type Action string

// Actions ...
const (
	// ActionNone ...
	ActionNone Action = ""
	// ActionInspectorLogin  ...
	ActionLogin Action = "login"
	// ActionAuthorize  ...
	ActionAuthorize Action = "authorize"
	// ActionInspectorLogout  ...
	ActionLogout Action = "logout"
	// ActionInspectionCreate ...
	ActionInspectionCreate Action = "inspection-create"
	// ActionInspectionOpen ...
	ActionInspectionOpen Action = "inspection-open"
	// ActionInspectionClose ...
	ActionInspectionClose Action = "inspection-close"
	// ActionInspectionFinishTemp ...
	ActionInspectionFinishTemp Action = "inspection-finish_temp"
	// ActionInspectionFinish ...
	ActionInspectionFinish Action = "inspection-finish"
	// ActionInspectionFinishedForce ...
	ActionInspectionFinishedForce Action = "inspection-finished_force"
	// ActionInspectionFinish ...
	ActionInspectionCreateResult Action = "inspection-create_result"
)

// ObjectType ...
type ObjectType string

const (
	// ObjectNone ...
	ObjectNone ObjectType = ""
	// ObjectInspection ...
	ObjectInspection ObjectType = "inspection"
	// ObjectInspectionResult ...
	ObjectInspectionResult ObjectType = "inspection_result"
)
