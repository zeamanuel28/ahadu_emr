package department

type CreateDepartmentDTO struct {
	Name     string  `json:"name" binding:"required"`
	Code     string  `json:"code" binding:"required"`
	ParentID *string `json:"parentId"`
}

type UpdateDepartmentDTO struct {
	Name     *string `json:"name"`
	Code     *string `json:"code"`
	ParentID *string `json:"parentId"`
}
