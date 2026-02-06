package allergy_reaction

import (
	"saas-api/core"
)

// AllergyReaction represents a type of reaction (reference/lookup table)
type AllergyReaction struct {
	core.BaseModel
	Name        string `gorm:"uniqueIndex;not null" json:"name"` // e.g., "Hives", "Anaphylaxis", "Rash"
	Description string `json:"description,omitempty"`
}

func (AllergyReaction) TableName() string {
	return "allergy_reactions"
}
