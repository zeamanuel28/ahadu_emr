package branch

import (
	"saas-api/core"

	"gorm.io/gorm"
)

type BranchService struct {
	core.BaseService[Branch]
}

func NewBranchService(db *gorm.DB) *BranchService {
	service := &BranchService{
		BaseService: *core.NewBaseService[Branch](db),
	}
	service.BaseService.SearchableFields = []string{"name", "code", "city", "phone"}
	return service
}
