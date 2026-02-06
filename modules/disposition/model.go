package disposition

import (
	"saas-api/core"

	"github.com/google/uuid"
)

// Custom types for Enums
type DispositionType string
type DispositionStatus string

const (
	Discharge DispositionType = "DISCHARGE"
	Transfer  DispositionType = "TRANSFER"
	Admit     DispositionType = "ADMIT"
	AMA       DispositionType = "AMA"
	Referred  DispositionType = "REFERRED"
	Deceased  DispositionType = "DECEASED"

	InProgress DispositionStatus = "IN_PROGRESS"
	Completed  DispositionStatus = "COMPLETED"
)

type Disposition struct {
	core.BaseModel

	VisitID uuid.UUID         `gorm:"type:uuid;uniqueIndex;not null" json:"visit_id"`
	Type    DispositionType   `gorm:"type:varchar(20);not null" json:"type"`
	Status  DispositionStatus `gorm:"type:varchar(20);default:'IN_PROGRESS'" json:"status"`
	Note    string            `gorm:"type:text" json:"note"`

	// Optional: Which department are they moving to?
	DepartmentID *uuid.UUID `gorm:"type:uuid" json:"department_id"`
}

func (Disposition) TableName() string {
	return "dispositions"
}
