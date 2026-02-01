package branch

import (
	"saas-api/core"
)

type Branch struct {
	core.BaseModel
	Name    string `json:"name"`
	Code    string `gorm:"uniqueIndex;not null" json:"code"`
	City    string `json:"city,omitempty"`
	SubCity string `json:"subCity,omitempty"`
	Wereda  string `json:"wereda,omitempty"`
	Phone   string `json:"phone,omitempty"`
}
