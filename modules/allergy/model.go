package allergy

import (
	"saas-api/core"
)

// Severity represents the severity level of an allergy
type Severity string

const (
	SeverityMild     Severity = "MILD"
	SeverityModerate Severity = "MODERATE"
	SeveritySevere   Severity = "SEVERE"
)

// AllergyName represents a type of allergy (reference/lookup table)
type AllergyName struct {
	core.BaseModel
	Name        string `gorm:"uniqueIndex;not null" json:"name"` // e.g., "Peanuts", "Penicillin", "Latex"
	Description string `json:"description,omitempty"`
}

// AllergyReaction represents a type of reaction (reference/lookup table)
type AllergyReaction struct {
	core.BaseModel
	Name        string `gorm:"uniqueIndex;not null" json:"name"` // e.g., "Hives", "Anaphylaxis", "Rash"
	Description string `json:"description,omitempty"`
}

// Allergy represents a patient's allergy
type Allergy struct {
	core.BaseModel

	PatientID     string `gorm:"not null;index" json:"patient_id_no"`
	AllergyNameID uint   `gorm:"not null;index" json:"allergy_name_id"`

	
	AllergyName *AllergyName       `json:"allergy_name,omitempty" gorm:"foreignKey:AllergyNameID"`
	Reactions   []*AllergyReaction `json:"reactions,omitempty" gorm:"many2many:allergy_allergy_reactions;"`

	Severity Severity `gorm:"type:varchar(20);not null" json:"severity"`
	Notes    string   `json:"notes,omitempty"`
}
