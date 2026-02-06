package allergy_reaction

// CreateAllergyReactionDTO defines the fields required to create an allergy reaction
type CreateAllergyReactionDTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description,omitempty"`
}

// UpdateAllergyReactionDTO defines the fields allowed for updating an allergy reaction
type UpdateAllergyReactionDTO struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}
