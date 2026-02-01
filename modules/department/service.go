package department

import (
	"saas-api/core"

	"gorm.io/gorm"
)

type DepartmentService struct {
	core.BaseService[Department]
}

func NewDepartmentService(db *gorm.DB) *DepartmentService {
	service := &DepartmentService{
		BaseService: *core.NewBaseService[Department](db),
	}
	service.BaseService.SearchableFields = []string{"name", "code"}
	service.BaseService.Preloads = []string{"Parent", "Children"}
	return service
}
