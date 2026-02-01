package employee

import (
	"saas-api/core"

	"gorm.io/gorm"
)

type EmployeeService struct {
	core.BaseService[Employee]
}

func NewEmployeeService(db *gorm.DB) *EmployeeService {
	service := &EmployeeService{
		BaseService: *core.NewBaseService[Employee](db),
	}
	// Configure default preloads and searchable fields
	service.BaseService.Preloads = []string{"Branch", "Department", "Position"}
	service.BaseService.SearchableFields = []string{"full_name_english", "employee_no", "email", "phone"}
	return service
}

// GetAll and GetById are now handled by BaseService using configured Preloads
