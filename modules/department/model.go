package department

import (
	"saas-api/core"

	"github.com/google/uuid"
)

type Department struct {
	core.BaseModel
	Name     string     `json:"name"`
	Code     string     `json:"code,omitempty"`
	ParentID *uuid.UUID `gorm:"type:uuid" json:"parentId,omitempty"`

	// Hierarchy relations
	Parent   *Department  `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
	Children []Department `gorm:"foreignKey:ParentID" json:"children,omitempty"`
}
