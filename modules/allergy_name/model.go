package allergy_name

import (
	"saas-api/core"
)

// AllergyName represents a type of allergy (reference/lookup table)
type AllergyName struct {
	core.BaseModel
	Name        string `gorm:"uniqueIndex;not null" json:"name"` // e.g., "Peanuts", "Penicillin", "Latex"
	Description string `json:"description,omitempty"`
}

func (AllergyName) TableName() string {
	return "allergy_names"
}
