package models

// ExpertsInfo ...
type ExpertsInfo struct {
	EmployeeJobtitles []EmployeeJobTitle `json:"employee_job_titles"`
	Employees         []Employee         `json:"employees"`
}

// EmployeeJobTitle ...
type EmployeeJobTitle struct {
	JobTitles  int    `json:"job_titles"`
	NationalID string `json:"national_id"`
}

// ExpertsInfos ...
// type ExpertsInfos []*ExpertsInfo

// AgencyExpertsInfo ...
type AgencyExpertsInfo struct {
	AgencyCode  int              `json:"agency_code"`
	CompanyID   int              `json:"company_id"`
	Domain      InspectionDomain `json:"domain"`
	ExpertsInfo ExpertsInfo      `json:"experts_info"`
}
