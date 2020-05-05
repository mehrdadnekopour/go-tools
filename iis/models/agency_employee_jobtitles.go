package models

import (
	"github.com/mehrdadnekopour/go-tools/mypes"
)

// AgencyEmployeeJobTiles ...
type AgencyEmployeeJobTiles struct {
	mypes.Model
	AgencyID   uint      `json:"agency_id" gorm:"unique_index:ag_emp_jt_unique_index"`
	AgencyCode int       `json:"agency_code"`
	CompanyID  uint      `json:"company_id"`
	EmployeeID uint      `json:"employee_id" gorm:"unique_index:ag_emp_jt_unique_index"`
	NationalID string    `json:"national_id"`
	Jobtitles  int       `json:"job_titles" gorm:"unique_index:ag_emp_jt_unique_index"`
	Employee   *Employee `json:"employee"`
}
