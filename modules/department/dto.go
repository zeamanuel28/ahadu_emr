package department

import (
	"github.com/google/uuid"
)

type CreateDepartmentDTO struct {
	Name     string     `json:"name" binding:"required"`
	Code     string     `json:"code" binding:"required"`
	ParentID *uuid.UUID `json:"parentId"`
}

type UpdateDepartmentDTO struct {
	Name     *string    `json:"name"`
	Code     *string    `json:"code"`
	ParentID *uuid.UUID `json:"parentId"`
}
