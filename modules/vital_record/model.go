package vital_record

import (
	"saas-api/core"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

// VitalRecord represents medical vitals stored as JSON
type VitalRecord struct {
	core.BaseModel
	VisitID uuid.UUID `gorm:"type:uuid;not null;index" json:"visit_id"`

	// Stores: {"BP": "120/80", "HR": 75, "TEMP": 37.2}
	Values datatypes.JSON `gorm:"type:jsonb;not null" json:"values"`
}

func (VitalRecord) TableName() string {
	return "vital_records"
}
