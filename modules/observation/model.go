package observation

import (
	"saas-api/core"

	"github.com/google/uuid"
)

// ChiefComplaint represents a reference for symptoms/complaints
type ChiefComplaint struct {
	core.BaseModel
	Name        string `gorm:"unique;not null;index" json:"name"`
	Description string `gorm:"type:text" json:"description"`
}

func (ChiefComplaint) TableName() string {
	return "chief_complaints"
}

// Observation represents a specific symptom recorded during a visit
type Observation struct {
	core.BaseModel
	VisitID          uuid.UUID       `gorm:"type:uuid;not null;index" json:"visit_id"`
	ChiefComplaintID uuid.UUID       `gorm:"type:uuid;not null" json:"chief_complaint_id"`
	ChiefComplaint   *ChiefComplaint `gorm:"foreignKey:ChiefComplaintID" json:"chief_complaint"`
	Note             string          `gorm:"type:text" json:"note"`
}

func (Observation) TableName() string {
	return "observations"
}
