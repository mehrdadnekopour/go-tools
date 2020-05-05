package models

import (
	"github.com/mehrdadnekopour/go-tools/mypes"
)

// CourseInfo ...
type CourseInfo struct {
	mypes.Model
	Domain    InspectionDomain `json:"domain" gorm:"index"`
	CourseID  uint             `json:"course_id"`
	TopicID   uint             `json:"topic_id"`
	Jobtitle  uint             `json:"jobtitle"`
	CarID     uint             `json:"car_id"`
	CompanyID uint             `json:"company_id"`
	Course    *Course          `json:"-"`
	Topic     *Topic           `json:"-"`
	Car       *Car             `json:"-"`
	Company   *Company         `json:"-"`
	// Employees *Employees      `json:"-"`
}

// CourseInfos ...
type CourseInfos []*CourseInfo
