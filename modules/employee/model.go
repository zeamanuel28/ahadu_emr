package employee

import (
	"time"

	"saas-api/core"
	"saas-api/modules/branch"
	"saas-api/modules/department"
	"saas-api/modules/position"
)

type Gender string

const (
	GenderMale   Gender = "MALE"
	GenderFemale Gender = "FEMALE"
)

type Employee struct {
	core.BaseModel
	ProfilePicture  string     `json:"profilePicture"`
	EmployeeNo      string     `gorm:"uniqueIndex;not null" json:"employeeNo"`
	FullNameEnglish string     `gorm:"not null" json:"fullNameEnglish"`
	FullNameAmharic string     `json:"fullNameAmharic"`
	Gender          Gender     `json:"gender"`
	DateOfBirth     time.Time  `json:"dateOfBirth"`
	Email           string     `gorm:"uniqueIndex" json:"email"`
	Phone           string     `gorm:"uniqueIndex" json:"phone"`
	Status          string     `gorm:"default:ACTIVE" json:"status"`
	HireDate        time.Time  `json:"hireDate"`
	TerminationDate *time.Time `json:"terminationDate,omitempty"`

	BranchID     string `json:"branchId"`
	DepartmentID string `json:"departmentId"`
	PositionID   string `json:"positionId"`

	Branch     *branch.Branch         `gorm:"foreignKey:BranchID" json:"branch,omitempty"`
	Department *department.Department `gorm:"foreignKey:DepartmentID" json:"department,omitempty"`
	Position   *position.Position     `gorm:"foreignKey:PositionID" json:"position,omitempty"`
}
