package disposition

import (
	"errors"
	"saas-api/core"
	"saas-api/modules/department"

	"gorm.io/gorm"
)

type DispositionService struct {
	*core.BaseService[Disposition]
}

func NewDispositionService(db *gorm.DB) *DispositionService {
	service := core.NewBaseService[Disposition](db)

	return &DispositionService{
		BaseService: service,
	}
}

// Create validates that the visit exists before creating the disposition
func (s *DispositionService) Create(disposition *Disposition) error {
	// Validate that the visit exists before creating the disposition
	var count int64
	err := s.DB.Table("visits").Where("id = ? AND deleted_at IS NULL", disposition.VisitID).Count(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("visit with ID " + disposition.VisitID.String() + " does not exist")
	}

	// Validate department if provided
	if disposition.DepartmentID != nil {
		var existingDept department.Department
		err := s.DB.Where("id = ?", *disposition.DepartmentID).First(&existingDept).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return errors.New("department with ID " + disposition.DepartmentID.String() + " does not exist")
			}
			return err
		}
	}

	// Proceed with creating the disposition
	return s.BaseService.Create(disposition)
}
