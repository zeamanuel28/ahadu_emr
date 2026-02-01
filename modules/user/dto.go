package user

type CreateUserDTO struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"fullName"`
	Status   string `json:"status" example:"active"`
}

type UpdateUserDTO struct {
	Username *string `json:"username"`
	Email    *string `json:"email"`
	FullName *string `json:"fullName"`
	Status   *string `json:"status"`
}
