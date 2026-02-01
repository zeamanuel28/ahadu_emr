package core

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"saas-api/shared/utils"

	"github.com/gin-gonic/gin"
)

type BaseController[T any, CreateDTO any, UpdateDTO any] struct {
	Service *BaseService[T]
}

type CRUDController interface {
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Patch(c *gin.Context)
}

func NewBaseController[T any, CreateDTO any, UpdateDTO any](service *BaseService[T]) *BaseController[T, CreateDTO, UpdateDTO] {
	return &BaseController[T, CreateDTO, UpdateDTO]{
		Service: service,
	}
}

func (ctrl *BaseController[T, CreateDTO, UpdateDTO]) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.Query("q")

	// Parse new parameters
	include := c.Query("include")
	var includes []string
	if include != "" {
		// Simple comma separation, assumes no commas in relation names
		for _, inc := range strings.Split(include, ",") {
			includes = append(includes, strings.TrimSpace(inc))
		}
	}

	searchFieldsStr := c.Query("searchFields")
	var searchFields []string
	if searchFieldsStr != "" {
		for _, f := range strings.Split(searchFieldsStr, ",") {
			searchFields = append(searchFields, strings.TrimSpace(f))
		}
	}

	filtersJSON := c.Query("filters")
	var filters []FilterParam
	if filtersJSON != "" {
		// Use core.FilterParam (same package)
		json.Unmarshal([]byte(filtersJSON), &filters)
	}

	// Also support legacy filters via query map?
	// The new BaseService GetAll expects []FilterParam.
	// We can convert simple query params to Eq filters here if needed.

	sortJSON := c.Query("sort")
	var sorts []SortParam
	if sortJSON != "" {
		json.Unmarshal([]byte(sortJSON), &sorts)
	}

	params := QueryParams{
		Page:         page,
		Limit:        limit,
		Search:       search,
		SearchFields: searchFields,
		Filters:      filters,
		Sorts:        sorts,
		Includes:     includes,
	}

	data, total, totalPages, err := ctrl.Service.GetAll(params)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch data", err.Error())
		return
	}

	utils.SuccessResponse(c, data, &utils.Meta{
		Page:       page,
		Limit:      limit,
		TotalCount: total,
		TotalPages: totalPages,
	})
}

func (ctrl *BaseController[T, CreateDTO, UpdateDTO]) GetById(c *gin.Context) {
	id := c.Param("id")
	data, err := ctrl.Service.GetById(id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch record", err.Error())
		return
	}
	if data == nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Record not found", "")
		return
	}

	utils.SuccessResponse(c, data, nil)
}

func (ctrl *BaseController[T, CreateDTO, UpdateDTO]) Create(c *gin.Context) {
	var dto CreateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	// Convert DTO to Model (T)
	dtoBytes, err := json.Marshal(dto)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to process data", err.Error())
		return
	}

	var model T
	if err := json.Unmarshal(dtoBytes, &model); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to map to model", err.Error())
		return
	}

	if err := ctrl.Service.Create(&model); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create record", err.Error())
		return
	}

	utils.CreatedResponse(c, model)
}

func (ctrl *BaseController[T, CreateDTO, UpdateDTO]) Update(c *gin.Context) {
	id := c.Param("id")
	// PUT uses CreateDTO to enforce full resource validation (Replace semantics)
	var dto CreateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	dtoBytes, err := json.Marshal(dto)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to process data", err.Error())
		return
	}

	var data map[string]interface{}
	if err := json.Unmarshal(dtoBytes, &data); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to map data", err.Error())
		return
	}

	result, err := ctrl.Service.Update(id, data)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update record", err.Error())
		return
	}

	utils.SuccessResponse(c, result, nil)
}

func (ctrl *BaseController[T, CreateDTO, UpdateDTO]) Patch(c *gin.Context) {
	id := c.Param("id")
	// PATCH uses UpdateDTO (Partial semantics)
	var dto UpdateDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	dtoBytes, err := json.Marshal(dto)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to process data", err.Error())
		return
	}

	var data map[string]interface{}
	if err := json.Unmarshal(dtoBytes, &data); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to map data", err.Error())
		return
	}

	result, err := ctrl.Service.Patch(id, data)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update record", err.Error())
		return
	}

	utils.SuccessResponse(c, result, nil)
}
