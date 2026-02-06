package diagnosis_code

import (
	"saas-api/core"
)

// DiagnosisSystem defines the coding standard used (e.g., ICD10, SNOMED)
type DiagnosisSystem string

// DiagnosisSystem constants
const (
	ICD10  DiagnosisSystem = "ICD10"
	SNOMED DiagnosisSystem = "SNOMED"
	Custom DiagnosisSystem = "CUSTOM"
)

// DiagnosisCode represents a standardized medical diagnosis code
type DiagnosisCode struct {
	core.BaseModel

	// Code: "I10", "E11.9" - Unique index on code + system
	Code string `gorm:"type:varchar(20);not null;index:idx_code_system,unique" json:"code"`
	Name string `gorm:"type:varchar(255);not null" json:"name"`

	// System: Enum-like field for the coding system
	System DiagnosisSystem `gorm:"type:varchar(20);not null;index:idx_code_system,unique" json:"system"`

	Description string `gorm:"type:text" json:"description"`
	Active      bool   `gorm:"default:true" json:"active"`
}

// TableName returns the table name for GORM
func (DiagnosisCode) TableName() string {
	return "diagnosis_codes"
}
