package models

import (
	"github.com/mehrdadnekopour/go-tools/mypes"
)

// Car ...
type Car struct {
	mypes.Model
	Name          string            `json:"name"`
	Code          int               `json:"code"`
	CompanyID     uint              `json:"company_id"`
	CompanyTypeID uint              `json:"company_type_id"`
	Company       *Company          `json:"company"`
	CoverID       uint              `json:"cover_id"`
	Cover         *mypes.Mimage     `json:"cover"`
	Gallery       *mypes.Mimagables `json:"gallery"`
	Type          CarTypeEnum       `json:"type"`
}

// Cars ...
type Cars []*Car

// CarTypeEnum ...
type CarTypeEnum int

const (
	// CarTypeNormal ...
	CarTypeNormal CarTypeEnum = 61
	// CarTypeTruck ...
	CarTypeTruck CarTypeEnum = 62
	// CarTypeBus ...
	CarTypeBus CarTypeEnum = 63
)
