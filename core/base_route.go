package core

import (
	"saas-api/shared/utils"

	"github.com/gin-gonic/gin"
)

type RouteOptions struct {
	Tag            string
	PathPrefix     string
	CreateSchema   string
	UpdateSchema   string
	ResponseSchema string
}

type BaseRoute[T any, CreateDTO any, UpdateDTO any] struct {
	Controller CRUDController
}

func NewBaseRoute[T any, CreateDTO any, UpdateDTO any](controller CRUDController) *BaseRoute[T, CreateDTO, UpdateDTO] {
	return &BaseRoute[T, CreateDTO, UpdateDTO]{
		Controller: controller,
	}
}

func (r *BaseRoute[T, CreateDTO, UpdateDTO]) Register(group *gin.RouterGroup, opts RouteOptions) {
	// Register dynamic routes in Gin
	group.GET("", r.Controller.GetAll)
	group.GET("/:id", r.Controller.GetById)
	group.POST("", r.Controller.Create)
	group.PUT("/:id", r.Controller.Update)
	group.PATCH("/:id", r.Controller.Patch)

	// Register metadata in Dynamic Swagger Registry
	r.registerSwagger(opts)
}

func (r *BaseRoute[T, CreateDTO, UpdateDTO]) registerSwagger(opts RouteOptions) {
	pathPrefix := opts.PathPrefix
	tag := opts.Tag

	// Register DTO schemas if provided
	if opts.CreateSchema != "" && opts.CreateSchema != "object" {
		var createDTO CreateDTO
		utils.Registry.RegisterSchema(opts.CreateSchema, createDTO)
	}

	if opts.UpdateSchema != "" && opts.UpdateSchema != "object" {
		var updateDTO UpdateDTO
		utils.Registry.RegisterSchema(opts.UpdateSchema, updateDTO)
	}

	// Register Entity Schema if provided (T)
	if opts.ResponseSchema != "" && opts.ResponseSchema != "object" {
		var entity T
		utils.Registry.RegisterSchema(opts.ResponseSchema, entity)
	}

	// Default schemas if not provided
	createSchema := opts.CreateSchema
	if createSchema == "" {
		createSchema = "object"
	} else {
		createSchema = "#/definitions/" + createSchema
	}

	updateSchema := opts.UpdateSchema
	if updateSchema == "" {
		updateSchema = "object"
	} else {
		updateSchema = "#/definitions/" + updateSchema
	}

	responseSchema := opts.ResponseSchema
	if responseSchema == "" {
		responseSchema = "object" // Default to generic object if no schema provided
	} else {
		responseSchema = "#/definitions/" + responseSchema
	}

	// Helper to get { data: ref } properties
	getDataSchema := func(ref string, isArray bool) map[string]interface{} {
		if ref == "object" {
			return map[string]interface{}{"type": "object"}
		}

		dataSchema := map[string]interface{}{"$ref": ref}
		if isArray {
			dataSchema = map[string]interface{}{
				"type":  "array",
				"items": map[string]interface{}{"$ref": ref},
			}
		}

		return map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"data": dataSchema,
			},
		}
	}

	// GET /
	utils.Registry.RegisterPath(pathPrefix, "get", map[string]interface{}{
		"summary": "List " + tag,
		"tags":    []string{tag},
		"parameters": []map[string]interface{}{
			{"name": "page", "in": "query", "required": false, "type": "integer", "description": "Page number (default 1)"},
			{"name": "limit", "in": "query", "required": false, "type": "integer", "description": "Items per page (default 10)"},
			{"name": "q", "in": "query", "required": false, "type": "string", "description": "Search term"},
			{"name": "searchFields", "in": "query", "required": false, "type": "string", "description": "Comma-separated fields to search in (e.g. name,email)"},
			{"name": "filters", "in": "query", "required": false, "type": "string", "description": "JSON array of filters (e.g. [{\"field\":\"status\",\"operator\":\"eq\",\"value\":\"ACTIVE\"}])"},
			{"name": "sort", "in": "query", "required": false, "type": "string", "description": "JSON array of sort fields (e.g. [{\"field\":\"name\",\"order\":\"asc\"}])"},
		},
		"responses": map[string]interface{}{
			"200": map[string]interface{}{
				"description": "Success",
				"schema":      getDataSchema(responseSchema, true),
			},
		},
		"security": []map[string]interface{}{{"BearerAuth": []interface{}{}}},
	})

	// GET /:id
	utils.Registry.RegisterPath(pathPrefix+"/{id}", "get", map[string]interface{}{
		"summary": "Get " + tag + " by ID",
		"tags":    []string{tag},
		"parameters": []map[string]interface{}{
			{"name": "id", "in": "path", "required": true, "type": "string"},
		},
		"responses": map[string]interface{}{
			"200": map[string]interface{}{
				"description": "Success",
				"schema":      getDataSchema(responseSchema, false),
			},
			"404": map[string]interface{}{"description": "Not found"},
		},
		"security": []map[string]interface{}{{"BearerAuth": []interface{}{}}},
	})

	// POST /
	utils.Registry.RegisterPath(pathPrefix, "post", map[string]interface{}{
		"summary": "Create " + tag,
		"tags":    []string{tag},
		"parameters": []map[string]interface{}{
			{
				"name":     "request",
				"in":       "body",
				"required": true,
				"schema":   r.getSchemaRef(createSchema),
			},
		},
		"responses": map[string]interface{}{
			"201": map[string]interface{}{
				"description": "Created",
				"schema":      getDataSchema(responseSchema, false),
			},
		},
		"security": []map[string]interface{}{{"BearerAuth": []interface{}{}}},
	})

	// PUT /:id
	utils.Registry.RegisterPath(pathPrefix+"/{id}", "put", map[string]interface{}{
		"summary": "Update " + tag,
		"tags":    []string{tag},
		"parameters": []map[string]interface{}{
			{"name": "id", "in": "path", "required": true, "type": "string"},
			{
				"name":     "request",
				"in":       "body",
				"required": true, // PUT expects full resource
				"schema":   r.getSchemaRef(createSchema),
			},
		},
		"responses": map[string]interface{}{
			"200": map[string]interface{}{
				"description": "Updated",
				"schema":      getDataSchema(responseSchema, false),
			},
		},
		"security": []map[string]interface{}{{"BearerAuth": []interface{}{}}},
	})

	// PATCH /:id
	utils.Registry.RegisterPath(pathPrefix+"/{id}", "patch", map[string]interface{}{
		"summary": "Patch " + tag,
		"tags":    []string{tag},
		"parameters": []map[string]interface{}{
			{"name": "id", "in": "path", "required": true, "type": "string"},
			{
				"name":     "request",
				"in":       "body",
				"required": true,
				"schema":   r.getSchemaRef(updateSchema),
			},
		},
		"responses": map[string]interface{}{
			"200": map[string]interface{}{
				"description": "Patched",
				"schema":      getDataSchema(responseSchema, false),
			},
		},
		"security": []map[string]interface{}{{"BearerAuth": []interface{}{}}},
	})
}

func (r *BaseRoute[T, CreateDTO, UpdateDTO]) getSchemaRef(schema string) map[string]interface{} {
	if schema == "object" {
		return map[string]interface{}{"type": "object"}
	}
	return map[string]interface{}{"$ref": schema}
}
