package models

import (
	"github.com/mehrdadnekopour/go-tools/mypes"
)

// Agency ...
type Agency struct {
	mypes.Model
	Code            int             `json:"code"`
	Owner           string          `json:"owner"`
	OwnerMobile     string          `json:"owner_mobile"`
	Tel             string          `json:"tel"`
	Fax             string          `json:"fax"`
	Address         string          `json:"address"`
	CNGSupport      bool            `json:"cng_support"`
	Warranty        float64         `json:"waranty"`
	Guarantee       float64         `json:"guarantee"`
	Reception       float64         `json:"reception"`
	WorkingHours    float64         `json:"working_hours"`
	MechanicHours   float64         `json:"mechanic_hours"`
	ElectricHours   float64         `json:"electric_hours"`
	SuspensionHours float64         `json:"suspension_hours"`
	GasHours        float64         `json:"gas_hours"`
	LogoID          uint            `json:"logo_id"`
	Logo            *mypes.Mimage   `json:"logo"`
	CompanyCode     int             `json:"company_code"`
	CompanyID       uint            `json:"company_id"`
	Company         *Company        `json:"company"`
	CompanyTypeID   uint            `json:"company_type_id"`
	CityID          uint            `json:"city_id"`
	City            *mypes.City     `json:"city"`
	ProvinceID      uint            `json:"province_id"`
	Province        *mypes.Province `json:"province"`
	BVTypeID        int             `json:"bv_type_id"`
	New             bool            `json:"new"`
	// Cars           *Cars    `json:"cars"`
}

// Agencies ...
type Agencies []*Agency

// AccessSI     bool   `json:"access_si" gorm:"default:'true'"`
