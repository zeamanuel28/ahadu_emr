package position

import (
	"saas-api/core"
)

type Position struct {
	core.BaseModel
	DepartmentID string `json:"departmentId"`
	Title        string `json:"title"`
	Code         string `json:"code,omitempty"`
	Grade        string `json:"grade,omitempty"`
	IsManagerial bool   `json:"isManagerial"`
}
