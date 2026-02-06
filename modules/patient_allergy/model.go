package patient_allergy

import (
	"saas-api/core"
	"saas-api/modules/allergy_name"
	"saas-api/modules/allergy_reaction"

	"github.com/google/uuid"
)

// Severity represents the severity level of an allergy
type Severity string

const (
	SeverityMild     Severity = "MILD"
	SeverityModerate Severity = "MODERATE"
	SeveritySevere   Severity = "SEVERE"
)

// Allergy represents a patient's allergy
type Allergy struct {
	core.BaseModel

	PatientID     string    `gorm:"column:patient_id_no;not null;index" json:"patient_id_no"`
	AllergyNameID uuid.UUID `gorm:"column:allergy_name_id;not null;index" json:"allergy_name_id"`

	AllergyName *allergy_name.AllergyName           `json:"allergy_name,omitempty" gorm:"foreignKey:AllergyNameID"`
	Reactions   []*allergy_reaction.AllergyReaction `json:"reactions,omitempty" gorm:"many2many:patient_allergy_reactions;"`

	Severity Severity `gorm:"type:varchar(20);not null" json:"severity"`
	Notes    string   `json:"notes,omitempty"`
}

func (Allergy) TableName() string {
	return "patient_allergies"
}
