package vital_type

import (
	"saas-api/core"
)

type VitalType struct {
	core.BaseModel
	Name      string   `gorm:"unique;not null;index" json:"name"` // e.g., "BP", "HR"
	Unit      string   `json:"unit"`                              // e.g., "mmHg", "bpm"
	NormalMin *float64 `json:"normal_min"`                        // Reference range min
	NormalMax *float64 `json:"normal_max"`                        // Reference range max
}

func (VitalType) TableName() string {
	return "vital_types"
}
