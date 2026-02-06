package problem

import (
	"time"

	"saas-api/core"
	"saas-api/modules/diagnosis_code"

	"github.com/google/uuid"
)

type ProblemStatus string

const (
	ProblemActive   ProblemStatus = "ACTIVE"
	ProblemResolved ProblemStatus = "RESOLVED"
)

type Problem struct {
	core.BaseModel

	// Linked directly to the Patient's lifetime record
	PatientID uuid.UUID `gorm:"type:uuid;not null;index" json:"patient_id"`

	// Linked to the Master Diagnosis Code (e.g., ICD-10 for Diabetes)
	DiagnosisCodeID uuid.UUID                     `gorm:"type:uuid;not null" json:"diagnosis_code_id"`
	DiagnosisCode   *diagnosis_code.DiagnosisCode `gorm:"foreignKey:DiagnosisCodeID" json:"diagnosis_code,omitempty"`

	Status       ProblemStatus `gorm:"type:varchar(20);default:'ACTIVE'" json:"status"`
	OnsetDate    *time.Time    `json:"onset_date"`
	ResolvedDate *time.Time    `json:"resolved_date"`
	Notes        string        `gorm:"type:text" json:"notes"`
}

func (Problem) TableName() string {
	return "problems"
}
