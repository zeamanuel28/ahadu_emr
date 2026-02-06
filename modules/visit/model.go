package visit

import (
	"saas-api/core"
	"saas-api/modules/diagnosis"
	"saas-api/modules/disposition"
	"saas-api/modules/observation"
	"saas-api/modules/vital_record"

	"github.com/google/uuid"
)

type VisitType string

const (
	VisitNew       VisitType = "New"
	VisitFollowup  VisitType = "Followup"
	VisitEmergency VisitType = "Emergency"
)

type VisitStatus string

const (
	VisitWaiting    VisitStatus = "WAITING"
	VisitInProgress VisitStatus = "IN_PROGRESS"
	VisitCompleted  VisitStatus = "COMPLETED"
)

type Severity string

const (
	SeverityMild     Severity = "MILD"
	SeverityModerate Severity = "MODERATE"
	SeveritySevere   Severity = "SEVERE"
)

type Visit struct {
	core.BaseModel

	PatientIDNo string `gorm:"column:patient_id_no;not null;index" json:"patient_id_no"`

	Type     VisitType   `gorm:"type:varchar(20);not null" json:"type"`
	Status   VisitStatus `gorm:"type:varchar(20);not null;default:WAITING" json:"status"`
	Severity Severity    `gorm:"type:varchar(20);not null" json:"severity"`

	DepartmentID *uuid.UUID `gorm:"type:uuid" json:"department_id,omitempty"`

	Observations []observation.Observation  `gorm:"foreignKey:VisitID;constraint:OnDelete:CASCADE" json:"observations,omitempty"`
	Vitals       []vital_record.VitalRecord `gorm:"foreignKey:VisitID;constraint:OnDelete:CASCADE" json:"vitals,omitempty"`
	Diagnoses    []diagnosis.Diagnosis      `gorm:"foreignKey:VisitID;constraint:OnDelete:CASCADE" json:"diagnoses,omitempty"`
	Disposition  *disposition.Disposition   `gorm:"foreignKey:VisitID;constraint:OnDelete:CASCADE" json:"disposition,omitempty"`
}
