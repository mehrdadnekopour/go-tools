package models

// WorkingHistory ...
type WorkingHistory struct {
	ID                 uint   `json:"id"`
	Duration           int    `json:"duration"`
	PresentCertificate int    `json:"present_certificate"`
	Description        string `json:"description"`
	CompanyName        string `json:"company_name"`
	CertificateType    int    `json:"certificate_type"`
}
