package models

import (
	"github.com/mehrdadnekopour/go-tools/mypes"
)

// Course ...
type Course struct {
	mypes.Model
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Domain      InspectionDomain `json:"domain"`
}

// Courses ...
type Courses []*Course
