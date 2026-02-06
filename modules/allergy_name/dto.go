package allergy_name

// CreateAllergyNameDTO defines the fields required to create an allergy name
type CreateAllergyNameDTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description,omitempty"`
}

// UpdateAllergyNameDTO defines the fields allowed for updating an allergy name
type UpdateAllergyNameDTO struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}
