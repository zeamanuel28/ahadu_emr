package branch

type CreateBranchDTO struct {
	Name     string `json:"name" binding:"required"`
	Code     string `json:"code" binding:"required"`
	SubCity  string `json:"subCity"`
	Wereda   string `json:"wereda"`
	MangerID string `json:"mangerId"`
}

type UpdateBranchDTO struct {
	Name     *string `json:"name"`
	Code     *string `json:"code"`
	SubCity  *string `json:"subCity"`
	Wereda   *string `json:"wereda"`
	MangerID *string `json:"mangerId"`
}
