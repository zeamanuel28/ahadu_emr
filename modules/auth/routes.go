package auth

import (
	"saas-api/shared/utils"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(parentGroup *gin.RouterGroup, controller *AuthController) {
	group := parentGroup.Group("/auth")

	// Register Gin routes
	group.POST("/login", controller.Login)
	group.POST("/register", controller.Register)

	// Register Swagger documentation
	utils.Registry.RegisterCustomPath(utils.PathConfig{
		Path:           "/auth/login",
		Method:         "POST",
		Summary:        "Login user",
		Description:    "OK",
		Tags:           []string{"Auth"},
		RequestSchema:  "auth.LoginRequest",
		ResponseSchema: "auth.LoginResponse",
		StatusCode:     "200",
		ErrorResponses: map[string]string{
			"401": "Unauthorized",
		},
		RequireAuth: false,
	})

	utils.Registry.RegisterCustomPath(utils.PathConfig{
		Path:           "/auth/register",
		Method:         "POST",
		Summary:        "Register user",
		Description:    "Created",
		Tags:           []string{"Auth"},
		RequestSchema:  "auth.RegisterRequest",
		ResponseSchema: "user.User",
		StatusCode:     "201",
		ErrorResponses: map[string]string{
			"400": "Bad Request",
		},
		RequireAuth: false,
	})

	// Register schemas
	utils.Registry.RegisterSchema("auth.LoginRequest", LoginRequest{})
	utils.Registry.RegisterSchema("auth.LoginResponse", LoginResponse{})
	utils.Registry.RegisterSchema("auth.RegisterRequest", RegisterRequest{})
}
