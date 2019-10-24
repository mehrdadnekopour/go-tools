package models

import (
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"

	"github.com/mehrdadnekopour/go-tools/mypes"
)

// Inspection ...
type Inspection struct {
	mypes.Model
	Domain      InspectionDomain `json:"domain"`
	Round       int              `json:"round"`
	Period      InspectionPeriod `json:"period"`
	Shift       InspectionShift  `json:"shift"`
	Status      InspectionStatus `json:"status"`
	Type        InspectionType   `json:"type"`
	Parts       InspectionParts  `json:"parts"`
	DateAssiged *time.Time       `json:"date_assigned"`
	DateOppened *time.Time       `json:"date_openned"`
	DateDone    *time.Time       `json:"date_done"`
	AgencyCode  int              `json:"agency_code"`
	AgencyOwner string           `json:"agency_owner"`
	CompanyID   int              `json:"company_id"`
	Checklist   postgres.Jsonb   `json:"checklist"` // Think about content of checklist
}

// InspectionDomain ...
type InspectionDomain string

// Domains ...
const (
	// DomainAfterSale ...
	DomainAfterSale InspectionDomain = "aftersale"
	// DomainSale ...
	DomainSale InspectionDomain = "sale"
)

// InspectionShift ...
type InspectionShift int

// Shifts ...
const (
	// ShiftMorning ...
	ShiftMorning InspectionShift = 0
	// ShiftAfternoon ...
	ShiftAfternoon InspectionShift = 1
)

// InspectionStatus ...
type InspectionStatus int

// Status ...
const (
	// StatusNone ...
	StatusNone InspectionStatus = 0
	// StatusFinishTemp ...
	StatusFinishTemp InspectionStatus = 1
	// StatusFinished ...
	StatusFinished InspectionStatus = 2
	// StatusUnderConstruction ...
	StatusUnderConstruction InspectionStatus = 4
	// StatusClosed ...
	StatusClosed InspectionStatus = 8
	// StatusDisaffiliation ...
	StatusDisaffiliation InspectionStatus = 16
	// StatusOtherReason ...
	StatusOtherReason InspectionStatus = 32
	// StatusFinishedForce ...
	StatusFinishedForce InspectionStatus = 1024
)

// InspectionType ...
type InspectionType int

// Types ...
const (
	// TypeSingle ...
	TypeSingle InspectionType = 1
	// TypeDouble ...
	TypeDouble InspectionType = 2
)

// InspectionPeriod ....
type InspectionPeriod int

// Periods ...
const (
	// FirstPeriod ...
	FirstPeriod InspectionPeriod = 1
	// SecondPeriod ...
	SecondPeriod InspectionPeriod = 2
	// ThirdPeriod ...
	ThirdPeriod InspectionPeriod = 3
	// FourthPeriod ...
	FourthPeriod InspectionPeriod = 4
	// FifthPeriod ...
	FifthPeriod InspectionPeriod = 5
)

// InspectionParts ...
type InspectionParts int

// Parts
const (
	// PartTools ...
	PartTools InspectionParts = 1
	// PartEmployees ...
	PartEmployees InspectionParts = 2
	// PartComplete ...
	PartComplete InspectionParts = 3
)
