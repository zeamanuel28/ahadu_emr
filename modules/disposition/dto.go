package disposition

import (
	"github.com/google/uuid"
)

// CreateDispositionDTO defines the fields required to create a disposition
type CreateDispositionDTO struct {
	VisitID      uuid.UUID         `json:"visit_id" binding:"required"`
	Type         DispositionType   `json:"type" binding:"required,oneof=DISCHARGE TRANSFER ADMIT AMA REFERRED DECEASED"`
	Status       DispositionStatus `json:"status" binding:"omitempty,oneof=IN_PROGRESS COMPLETED"`
	Note         string            `json:"note"`
	DepartmentID *uuid.UUID        `json:"department_id,omitempty"`
}

// UpdateDispositionDTO defines the fields allowed for updating a disposition
type UpdateDispositionDTO struct {
	Type         *DispositionType   `json:"type,omitempty" binding:"omitempty,oneof=DISCHARGE TRANSFER ADMIT AMA REFERRED DECEASED"`
	Status       *DispositionStatus `json:"status,omitempty" binding:"omitempty,oneof=IN_PROGRESS COMPLETED"`
	Note         *string            `json:"note,omitempty"`
	DepartmentID *uuid.UUID         `json:"department_id,omitempty"`
}
