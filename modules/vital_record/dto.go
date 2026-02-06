package vital_record

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

// CreateVitalRecordDTO defines fields for creating a vital record
type CreateVitalRecordDTO struct {
	VisitID uuid.UUID      `json:"visit_id" binding:"required"`
	Values  datatypes.JSON `json:"values" binding:"required"`
}

// UpdateVitalRecordDTO defines fields for updating a vital record
type UpdateVitalRecordDTO struct {
	VisitID *uuid.UUID      `json:"visit_id,omitempty"`
	Values  *datatypes.JSON `json:"values,omitempty"`
}
