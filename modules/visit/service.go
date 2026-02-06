package visit

import (
	"errors"
	"saas-api/core"
	"saas-api/modules/department"
	"saas-api/modules/patient"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VisitService struct {
	*core.BaseService[Visit]
}

func NewVisitService(db *gorm.DB) *VisitService {
	service := core.NewBaseService[Visit](db)

	// Define searchable fields
	service.SearchableFields = []string{
		"patient_id_no",
	}

	// Default preloads
	service.Preloads = []string{"Observations", "Vitals", "Diagnoses", "Disposition"}

	return &VisitService{
		BaseService: service,
	}
}

// Create creates a new visit after validating that the patient exists
func (s *VisitService) Create(visit *Visit) error {
	// Validate that the patient exists before creating the visit
	var existingPatient patient.Patient
	err := s.DB.Where("id_no = ?", visit.PatientIDNo).First(&existingPatient).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("patient with IDNo " + visit.PatientIDNo + " does not exist")
		}
		return err
	}

	// Validate department if provided
	if visit.DepartmentID != nil && *visit.DepartmentID != uuid.Nil {
		var existingDept department.Department
		err := s.DB.Where("id = ?", *visit.DepartmentID).First(&existingDept).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return errors.New("department with ID " + visit.DepartmentID.String() + " does not exist")
			}
			return err
		}
	}

	// Proceed with creating the visit
	return s.BaseService.Create(visit)
}
