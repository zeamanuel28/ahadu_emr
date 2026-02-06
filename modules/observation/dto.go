package observation

import "github.com/google/uuid"

// CreateChiefComplaintDTO defines fields for creating a chief complaint
type CreateChiefComplaintDTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description,omitempty"`
}

// UpdateChiefComplaintDTO defines fields for updating a chief complaint
type UpdateChiefComplaintDTO struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// CreateObservationDTO defines fields for creating an observation
type CreateObservationDTO struct {
	VisitID          uuid.UUID `json:"visit_id" binding:"required"`
	ChiefComplaintID uuid.UUID `json:"chief_complaint_id" binding:"required"`
	Note             string    `json:"note,omitempty"`
}

// UpdateObservationDTO defines fields for updating an observation
type UpdateObservationDTO struct {
	VisitID          *uuid.UUID `json:"visit_id,omitempty"`
	ChiefComplaintID *uuid.UUID `json:"chief_complaint_id,omitempty"`
	Note             *string    `json:"note,omitempty"`
}
