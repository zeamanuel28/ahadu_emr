package problem

import (
	"time"

	"github.com/google/uuid"
)

// CreateProblemDTO defines the fields required to create a problem
type CreateProblemDTO struct {
	PatientID       uuid.UUID     `json:"patient_id" binding:"required"`
	DiagnosisCodeID uuid.UUID     `json:"diagnosis_code_id" binding:"required"`
	Status          ProblemStatus `json:"status" binding:"omitempty,oneof=ACTIVE RESOLVED"`
	OnsetDate       *time.Time    `json:"onset_date"`
	ResolvedDate    *time.Time    `json:"resolved_date"`
	Notes           string        `json:"notes,omitempty"`
}

// UpdateProblemDTO defines the fields allowed for updating a problem
type UpdateProblemDTO struct {
	PatientID       *uuid.UUID     `json:"patient_id,omitempty"`
	DiagnosisCodeID *uuid.UUID     `json:"diagnosis_code_id,omitempty"`
	Status          *ProblemStatus `json:"status,omitempty" binding:"omitempty,oneof=ACTIVE RESOLVED"`
	OnsetDate       *time.Time     `json:"onset_date,omitempty"`
	ResolvedDate    *time.Time     `json:"resolved_date,omitempty"`
	Notes           *string        `json:"notes,omitempty"`
}
