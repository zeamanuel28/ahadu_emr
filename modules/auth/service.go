package auth

import (
	"errors"

	"saas-api/modules/user"
	"saas-api/shared/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserService *user.UserService
}

func NewAuthService(userService *user.UserService) *AuthService {
	return &AuthService{
		UserService: userService,
	}
}

func (s *AuthService) Login(email, password string) (string, *user.User, error) {
	u, err := s.UserService.GetByEmail(email)
	if err != nil {
		return "", nil, err
	}
	if u == nil {
		return "", nil, errors.New("invalid credentials")
	}

	if u.Status != user.UserStatusActive {
		return "", nil, errors.New("account is not active")
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	var roles []string
	for _, role := range u.Roles {
		roles = append(roles, role.Name)
	}

	token, err := utils.GenerateToken(u.ID.String(), u.Email, roles)
	if err != nil {
		return "", nil, err
	}

	return token, u, nil
}

func (s *AuthService) Register(userData *user.User) (*user.User, error) {
	return userData, s.UserService.Create(userData)
}
