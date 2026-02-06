package patient

import (
	"fmt"
	"time"

	"saas-api/core"
	"saas-api/modules/patient_allergy"
	"saas-api/modules/problem"

	"gorm.io/gorm"
)

type Sex string
type BloodType string
type EmergencyContactRelation string

// Sex constants
const (
	SexMale   Sex = "MALE"
	SexFemale Sex = "FEMALE"
)

// BloodType constants
const (
	BloodTypeAPos  BloodType = "A+"
	BloodTypeANeg  BloodType = "A-"
	BloodTypeBPos  BloodType = "B+"
	BloodTypeBNeg  BloodType = "B-"
	BloodTypeABPos BloodType = "AB+"
	BloodTypeABNeg BloodType = "AB-"
	BloodTypeOPos  BloodType = "O+"
	BloodTypeONeg  BloodType = "O-"
)

// EmergencyContactRelation constants
const (
	RelationParent  EmergencyContactRelation = "PARENT"
	RelationSpouse  EmergencyContactRelation = "SPOUSE"
	RelationSibling EmergencyContactRelation = "SIBLING"
	RelationFriend  EmergencyContactRelation = "FRIEND"
	RelationOther   EmergencyContactRelation = "OTHER"
)

// Patient model
type Patient struct {
	core.BaseModel
	IDNo                     string                   `gorm:"uniqueIndex;not null" json:"idNo"` // MRN generated automatically
	PatientImage             string                   `json:"patientImage,omitempty"`
	FullName                 string                   `gorm:"not null" json:"fullName"`
	Sex                      Sex                      `gorm:"type:varchar(10);not null" json:"sex"`
	DateOfBirth              time.Time                `gorm:"type:date;not null" json:"dateOfBirth"`
	BloodType                BloodType                `gorm:"type:varchar(5)" json:"bloodType,omitempty"`
	Phone                    string                   `gorm:"uniqueIndex;size:10" json:"phone,omitempty"`
	Email                    string                   `gorm:"uniqueIndex" json:"email,omitempty"`
	City                     string                   `json:"city,omitempty"`
	SubCity                  string                   `json:"subcity,omitempty"`
	Wereda                   string                   `json:"wereda,omitempty"`
	EmergencyContactName     string                   `json:"emergencyContactName,omitempty"`
	EmergencyContactPhone    string                   `gorm:"size:10" json:"emergencyContactPhone,omitempty"`
	EmergencyContactRelation EmergencyContactRelation `gorm:"type:varchar(10)" json:"emergencyContactRelation,omitempty"`
	Status                   string                   `gorm:"type:varchar(10);default:ACTIVE" json:"status"` // ACTIVE | INACTIVE | DECEASED

	// Relationships - One patient can have many allergies
	Allergies []patient_allergy.Allergy `json:"allergies,omitempty" gorm:"foreignKey:PatientID;references:IDNo"`
	Problems  []problem.Problem         `json:"problems,omitempty" gorm:"foreignKey:PatientID"`
}

// BeforeCreate hook to generate MRN
func (p *Patient) BeforeCreate(tx *gorm.DB) (err error) {
	if p.IDNo == "" {
		nextMRN, err := generateNextMRN(tx)
		if err != nil {
			return err
		}
		p.IDNo = nextMRN
	}
	return nil
}

// Helper to generate next MRN
func generateNextMRN(tx *gorm.DB) (string, error) {
	var lastPatient Patient
	err := tx.Order("id_no desc").First(&lastPatient).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return "", err
	}

	if lastPatient.IDNo == "" {
		return "MRNAA00001", nil
	}

	prefix := lastPatient.IDNo[:5]     // MRNAA
	numberPart := lastPatient.IDNo[5:] // 00001
	var nextNumber int
	fmt.Sscanf(numberPart, "%05d", &nextNumber)
	nextNumber++
	return fmt.Sprintf("%s%05d", prefix, nextNumber), nil
}
