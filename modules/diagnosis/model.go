package diagnosis

import (
	"saas-api/core"
	"saas-api/modules/diagnosis_code"

	"github.com/google/uuid"
)

type DiagnosisRole string
type DiagnosisStatus string

const (
	Primary   DiagnosisRole = "PRIMARY"
	Secondary DiagnosisRole = "SECONDARY"

	Confirmed   DiagnosisStatus = "CONFIRMED"
	Provisional DiagnosisStatus = "PROVISIONAL"
)

type Diagnosis struct {
	core.BaseModel

	// Link to the Visit
	VisitID uuid.UUID `gorm:"type:uuid;not null;index" json:"visit_id"`

	// Link to the Master Diagnosis Code (ICD10, etc.)
	DiagnosisCodeID uuid.UUID                     `gorm:"type:uuid;not null" json:"diagnosis_code_id"`
	DiagnosisCode   *diagnosis_code.DiagnosisCode `gorm:"foreignKey:DiagnosisCodeID" json:"diagnosis_code,omitempty"`

	Role   DiagnosisRole   `gorm:"type:varchar(20);not null" json:"role"`
	Status DiagnosisStatus `gorm:"type:varchar(20);not null" json:"status"`

	Notes string `gorm:"type:text" json:"notes"`
}

func (Diagnosis) TableName() string {
	return "diagnoses"
}
