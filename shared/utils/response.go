package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Meta struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalCount int64 `json:"totalCount"`
	TotalPages int   `json:"totalPages"`
}

type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func SuccessResponse(c *gin.Context, data interface{}, meta *Meta) {
	c.JSON(http.StatusOK, Response{
		Data: data,
		Meta: meta,
	})
}

func CreatedResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, Response{
		Data: data,
	})
}

func ErrorResponse(c *gin.Context, status int, message string, err string) {
	c.JSON(status, Response{
		Message: message,
		Error:   err,
	})
}
