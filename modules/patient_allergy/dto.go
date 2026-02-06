package patient_allergy

import "github.com/google/uuid"

// CreateAllergyDTO defines the fields required to create a patient allergy
type CreateAllergyDTO struct {
	PatientID     string      `json:"patient_id_no" binding:"required"`
	AllergyNameID uuid.UUID   `json:"allergy_name_id" binding:"required"`
	Severity      Severity    `json:"severity" binding:"required,oneof=MILD MODERATE SEVERE"`
	ReactionIDs   []uuid.UUID `json:"reaction_ids,omitempty"`
	Notes         string      `json:"notes,omitempty"`
}

// UpdateAllergyDTO defines the fields allowed for updating a patient allergy
type UpdateAllergyDTO struct {
	PatientID     *string      `json:"patient_id_no,omitempty"`
	AllergyNameID *uuid.UUID   `json:"allergy_name_id,omitempty"`
	Severity      *Severity    `json:"severity,omitempty" binding:"omitempty,oneof=MILD MODERATE SEVERE"`
	ReactionIDs   *[]uuid.UUID `json:"reaction_ids,omitempty"`
	Notes         *string      `json:"notes,omitempty"`
}
