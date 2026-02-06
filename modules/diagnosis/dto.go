package diagnosis

import (
	"github.com/google/uuid"
)

// CreateDiagnosisDTO defines the fields required to create a diagnosis
type CreateDiagnosisDTO struct {
	VisitID         uuid.UUID       `json:"visit_id" binding:"required"`
	DiagnosisCodeID uuid.UUID       `json:"diagnosis_code_id" binding:"required"`
	Role            DiagnosisRole   `json:"role" binding:"required,oneof=PRIMARY SECONDARY"`
	Status          DiagnosisStatus `json:"status" binding:"required,oneof=CONFIRMED PROVISIONAL"`
	Notes           string          `json:"notes,omitempty"`
}

// UpdateDiagnosisDTO defines the fields allowed for updating a diagnosis
type UpdateDiagnosisDTO struct {
	VisitID         *uuid.UUID       `json:"visit_id,omitempty"`
	DiagnosisCodeID *uuid.UUID       `json:"diagnosis_code_id,omitempty"`
	Role            *DiagnosisRole   `json:"role,omitempty" binding:"omitempty,oneof=PRIMARY SECONDARY"`
	Status          *DiagnosisStatus `json:"status,omitempty" binding:"omitempty,oneof=CONFIRMED PROVISIONAL"`
	Notes           *string          `json:"notes,omitempty"`
}
