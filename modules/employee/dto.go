package employee

import (
	"time"
)

type CreateEmployeeDTO struct {
	EmployeeNo      string    `json:"employeeNo" binding:"required"`
	FullNameEnglish string    `json:"fullNameEnglish" binding:"required"`
	FullNameAmharic string    `json:"fullNameAmharic"` // Optional
	Email           string    `json:"email" binding:"required,email"`
	Phone           string    `json:"phone" binding:"required"`
	DateOfBirth     time.Time `json:"dateOfBirth"`
	Gender          string    `json:"gender"`
	ProfilePicture  string    `json:"profilePicture"`
	Status          string    `json:"status"` // Default ACTIVE handled by DB but allowed in DTO
	HireDate        time.Time `json:"hireDate"`
	BranchID        string    `json:"branchId"`
	DepartmentID    string    `json:"departmentId"`
	PositionID      string    `json:"positionId"`
}

type UpdateEmployeeDTO struct {
	// Pointers for partial updates (PATCH)
	FullNameEnglish *string    `json:"fullNameEnglish"`
	FullNameAmharic *string    `json:"fullNameAmharic"`
	Email           *string    `json:"email"` // Should validate email format if present, usually handled by binding
	Phone           *string    `json:"phone"`
	DateOfBirth     *time.Time `json:"dateOfBirth"`
	Gender          *string    `json:"gender"`
	ProfilePicture  *string    `json:"profilePicture"`
	Status          *string    `json:"status"`
	HireDate        *time.Time `json:"hireDate"`
	TerminationDate *time.Time `json:"terminationDate"`
	BranchID        *string    `json:"branchId"`
	DepartmentID    *string    `json:"departmentId"`
	PositionID      *string    `json:"positionId"`
}
