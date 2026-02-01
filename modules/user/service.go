package user

import (
	"crypto/rand"
	"encoding/hex"

	"saas-api/core"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	core.BaseService[User]
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		BaseService: *core.NewBaseService[User](db),
	}
}

// Create overrides BaseService.Create to handle password hashing and role assignment
func (s *UserService) Create(user *User) error {
	// Generate random password if not provided
	if user.Password == "" {
		bytes := make([]byte, 8)
		if _, err := rand.Read(bytes); err != nil {
			return err
		}
		user.TempPassword = hex.EncodeToString(bytes)
		user.Password = user.TempPassword
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return s.DB.Create(user).Error
}

// GetByEmail finds a user by email
func (s *UserService) GetByEmail(email string) (*User, error) {
	var u User
	err := s.DB.Preload("Roles").Where("email = ? AND is_deleted = ?", email, false).First(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}
