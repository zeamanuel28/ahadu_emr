package allergy

// CreateAllergyDTO defines the fields required to create an allergy
type CreateAllergyDTO struct {
	PatientID     string   `json:"patient_id_no" binding:"required"`
	AllergyNameID uint     `json:"allergy_name_id" binding:"required"`
	Severity      Severity `json:"severity" binding:"required,oneof=MILD MODERATE SEVERE"`
	ReactionIDs   []uint   `json:"reaction_ids,omitempty"`
	Notes         string   `json:"notes,omitempty"`
}

// UpdateAllergyDTO defines the fields allowed for updating an allergy
type UpdateAllergyDTO struct {
	PatientID     *string   `json:"patient_id_no,omitempty"`
	AllergyNameID *uint     `json:"allergy_name_id,omitempty"`
	Severity      *Severity `json:"severity,omitempty" binding:"omitempty,oneof=MILD MODERATE SEVERE"`
	ReactionIDs   *[]uint   `json:"reaction_ids,omitempty"`
	Notes         *string   `json:"notes,omitempty"`
}
