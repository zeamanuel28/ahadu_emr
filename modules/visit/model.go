package visit

import (
	"saas-api/core"

	"gorm.io/datatypes"
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

	PatientIDNo string `gorm:"not null;index" json:"patient_id_no"`

	Type     VisitType   `gorm:"type:varchar(20);not null" json:"type"`
	Status   VisitStatus `gorm:"type:varchar(20);not null;default:WAITING" json:"status"`
	Severity Severity    `gorm:"type:varchar(20);not null" json:"severity"`

	DepartmentID *string `json:"department_id,omitempty"`

	Observations []Observation `gorm:"foreignKey:VisitID;constraint:OnDelete:CASCADE" json:"observations,omitempty"`
	Vitals       []VitalRecord `gorm:"foreignKey:VisitID;constraint:OnDelete:CASCADE" json:"vitals,omitempty"`
}

type Observation struct {
	core.BaseModel

	VisitID string `gorm:"not null;index" json:"visit_id"`

	ChiefComplaintID string `gorm:"not null" json:"chief_complaint_id"`

	Note *string `json:"note,omitempty"`
}

type VitalRecord struct {
	core.BaseModel

	VisitID string `gorm:"not null;index" json:"visit_id"`

	Values datatypes.JSON `gorm:"type:jsonb" json:"values"`
}
