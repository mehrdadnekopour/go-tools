package models

import (
	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mehrdadnekopour/go-tools/mypes"
)

// Employee ...
type Employee struct {
	mypes.Model
	NationalID                string              `gorm:"unique_index" json:"national_id"`
	FirstName                 string              `json:"first_name"`
	LastName                  string              `json:"last_name"`
	BirthDate                 string              `json:"birth_date"`
	BirthCertificateID        string              `json:"birth_certificate_id"` // شماره شناسنامه
	Address                   string              `json:"address"`
	Cellphone                 string              `json:"cellphone"`
	Description               string              `json:"description"`
	Email                     string              `json:"email"`
	Gender                    int                 `json:"gender"`
	JobDocs                   bool                `json:"job_docs"`
	EducationDocs             bool                `json:"education_docs"`
	LearningsDocs             bool                `json:"learnings_docs"`
	IdentificationDocs        bool                `json:"identification_docs"`
	PostalCode                string              `json:"postal_code"`
	Tel                       string              `json:"tel"`
	WorkingDayLimit           int                 `json:"working_day_limit"`
	WorkingHistoryDescription string              `json:"working_history_description"`
	EmployeeCourseInfos       EmployeeCourseInfos `json:"course_infos"`
	Education                 postgres.Jsonb      `json:"education"`
	WorkingHistoryIn          postgres.Jsonb      `json:"working_history_in"`
	WorkingHistoryOut         postgres.Jsonb      `json:"working_history_out"`
	WorkingHistories          postgres.Jsonb      `json:"working_histories"`
	// CourseInfos               postgres.Jsonb `json:"course_infos"`
}

// Employees ...
type Employees []*Employees

// EmployeeCourseInfo ...
type EmployeeCourseInfo struct {
	EmployeeID        uint             `json:"employee_id"`
	CourseInfoID      uint             `json:"course_info_id"`
	Domain            InspectionDomain `json:"domain"`
	CourseID          uint             `json:"course_id"`
	TopicID           uint             `json:"topic_id"`
	Jobtitle          uint             `json:"jobtitle"`
	CarID             uint             `json:"car_id"`
	CompanyID         uint             `json:"company_id"`
	HasCertificate    bool             `json:"has_certificate"`
	CertificateNumber string           `json:"certificate_no"`
	CertificateDate   string           `json:"certificate_date"`
	CertificateDesc   string           `json:"certificate_desc"`
}

// TableName ...
func (e EmployeeCourseInfo) TableName() string {
	return "employee_course_infos"
}

// EmployeeCourseInfos ...
type EmployeeCourseInfos []*EmployeeCourseInfo
