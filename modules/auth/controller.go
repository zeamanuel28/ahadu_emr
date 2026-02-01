package auth

import (
	"net/http"

	"saas-api/modules/user"
	"saas-api/shared/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService *AuthService
}

func NewAuthController(authService *AuthService) *AuthController {
	return &AuthController{
		AuthService: authService,
	}
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	token, u, err := ctrl.AuthService.Login(req.Email, req.Password)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Login failed", err.Error())
		return
	}

	utils.SuccessResponse(c, LoginResponse{
		Token: token,
		User:  *u,
	}, nil)
}

func (ctrl *AuthController) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	u := user.User{
		Username: req.Username,
		Email:    req.Email,
		FullName: req.FullName,
		Password: req.Password,
	}

	registeredUser, err := ctrl.AuthService.Register(&u)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Registration failed", err.Error())
		return
	}

	utils.CreatedResponse(c, registeredUser)
}
