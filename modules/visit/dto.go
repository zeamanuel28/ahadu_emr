package visit

import (
	"github.com/google/uuid"
)

// CreateVisitDTO defines the fields required to create a visit
type CreateVisitDTO struct {
	PatientIDNo  string      `json:"patient_id_no" binding:"required"`
	Type         VisitType   `json:"type" binding:"required,oneof=New Followup Emergency"`
	Status       VisitStatus `json:"status" binding:"omitempty,oneof=WAITING IN_PROGRESS COMPLETED"`
	Severity     Severity    `json:"severity" binding:"required,oneof=MILD MODERATE SEVERE"`
	DepartmentID *uuid.UUID  `json:"department_id,omitempty"`
}

// UpdateVisitDTO defines the fields allowed for updating a visit
type UpdateVisitDTO struct {
	Type         *VisitType   `json:"type,omitempty" binding:"omitempty,oneof=New Followup Emergency"`
	Status       *VisitStatus `json:"status,omitempty" binding:"omitempty,oneof=WAITING IN_PROGRESS COMPLETED"`
	Severity     *Severity    `json:"severity,omitempty" binding:"omitempty,oneof=MILD MODERATE SEVERE"`
	DepartmentID *uuid.UUID   `json:"department_id,omitempty"`
}
