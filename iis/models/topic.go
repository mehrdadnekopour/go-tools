package models

import (
	"github.com/mehrdadnekopour/go-tools/mypes"
)

// Topic ...
type Topic struct {
	mypes.Model
	Name        string           `json:"name"`
	Description string           `json:"description"`
	IsCNG       bool             `json:"is_cng"`
	Domain      InspectionDomain `json:"domain"`
}

// Topics ...
type Topics []*Topic
