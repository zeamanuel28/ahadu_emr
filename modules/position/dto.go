package position

type CreatePositionDTO struct {
	Title        string `json:"title" binding:"required"`
	Code         string `json:"code" binding:"required"`
	IsManagerial bool   `json:"isManagerial"`
	Grade        string `json:"grade"`
}

type UpdatePositionDTO struct {
	Title        *string `json:"title"`
	Code         *string `json:"code"`
	IsManagerial *bool   `json:"isManagerial"`
	Grade        *string `json:"grade"`
}
