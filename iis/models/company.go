package models

import (
	"github.com/mehrdadnekopour/go-tools/mypes"
)

// Company ...
type Company struct {
	mypes.Model
	Name        string        `json:"name"`
	Code        int           `json:"code"`
	Description string        `json:"description"`
	TypeID      uint          `json:"type_id"`
	Type        *CompanyType  `json:"type"`
	CEO         string        `json:"ceo"`
	Address     string        `json:"address"`
	Tel         string        `json:"tel"`
	Fax         string        `json:"fax"`
	LogoID      uint          `json:"logo_id"`
	Logo        *mypes.Mimage `json:"logo"`
	Agencies    Agencies      `json:"agencies"`
	Cars        Cars          `json:"cars"`
}

// Companies ...
type Companies []*Company

// CompanyType ...
type CompanyType struct {
	mypes.Model
	Title string `json:"title"`
}

// CompanyTypes ...
type CompanyTypes []*CompanyType
