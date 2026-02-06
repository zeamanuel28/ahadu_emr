package visit

import (
	"gorm.io/datatypes"
)

// CreateVisitDTO defines the fields required to create a visit
type CreateVisitDTO struct {
	PatientIDNo  string      `json:"patient_id_no" binding:"required"`
	Type         VisitType   `json:"type" binding:"required,oneof=New Followup Emergency"`
	Status       VisitStatus `json:"status" binding:"omitempty,oneof=WAITING IN_PROGRESS COMPLETED"`
	Severity     Severity    `json:"severity" binding:"required,oneof=MILD MODERATE SEVERE"`
	DepartmentID *string     `json:"department_id,omitempty"`
}

// UpdateVisitDTO defines the fields allowed for updating a visit
type UpdateVisitDTO struct {
	Type         *VisitType   `json:"type,omitempty" binding:"omitempty,oneof=New Followup Emergency"`
	Status       *VisitStatus `json:"status,omitempty" binding:"omitempty,oneof=WAITING IN_PROGRESS COMPLETED"`
	Severity     *Severity    `json:"severity,omitempty" binding:"omitempty,oneof=MILD MODERATE SEVERE"`
	DepartmentID *string      `json:"department_id,omitempty"`
}

// CreateObservationDTO defines the fields required to create an observation
type CreateObservationDTO struct {
	VisitID          string  `json:"visit_id" binding:"required"`
	ChiefComplaintID string  `json:"chief_complaint_id" binding:"required"`
	Note             *string `json:"note,omitempty"`
}

// CreateVitalRecordDTO defines the fields required to create a vital record
type CreateVitalRecordDTO struct {
	VisitID string         `json:"visit_id" binding:"required"`
	Values  datatypes.JSON `json:"values" binding:"required"`
}
