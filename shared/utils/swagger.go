package utils

import (
	"encoding/json"
	"reflect"
	"strings"
	"sync"
	"time"
)

// SwaggerRegistry holds dynamic Swagger definitions
type SwaggerRegistry struct {
	mu      sync.RWMutex
	Paths   map[string]map[string]interface{} `json:"paths"`
	Schemas map[string]interface{}            `json:"schemas"`
}

var Registry = &SwaggerRegistry{
	Paths:   make(map[string]map[string]interface{}),
	Schemas: make(map[string]interface{}),
}

// RegisterPath adds a path to the dynamic registry
func (r *SwaggerRegistry) RegisterPath(path, method string, operation interface{}) {
	r.mu.Lock()
	defer r.mu.Unlock()

	method = strings.ToLower(method)
	if _, ok := r.Paths[path]; !ok {
		r.Paths[path] = make(map[string]interface{})
	}
	r.Paths[path][method] = operation
}

// RegisterSchema adds a schema definition to the dynamic registry
func (r *SwaggerRegistry) RegisterSchema(name string, dto interface{}) {
	r.mu.Lock()
	defer r.mu.Unlock()

	schema := generateSchemaFromStruct(dto)
	r.Schemas[name] = schema
}

// generateSchemaFromStruct converts a Go struct to OpenAPI schema using reflection
func generateSchemaFromStruct(dto interface{}) map[string]interface{} {
	return generateSchemaWithRecursion(dto, make(map[reflect.Type]bool))
}

func generateSchemaWithRecursion(dto interface{}, visited map[reflect.Type]bool) map[string]interface{} {
	t := reflect.TypeOf(dto)

	// Handle pointer types
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return map[string]interface{}{"type": "object"}
	}

	// Check for recursion
	if visited[t] {
		return map[string]interface{}{"type": "object"} // Break recursion
	}
	visited[t] = true
	defer delete(visited, t) // Cleanup after visiting

	properties := make(map[string]interface{})
	required := []string{}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Skip unexported fields
		if !field.IsExported() {
			continue
		}

		// Get JSON tag
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		// Parse JSON tag (handle "name,omitempty" format)
		jsonName := strings.Split(jsonTag, ",")[0]

		// Get binding tag for validation
		bindingTag := field.Tag.Get("binding")

		// Generate property schema
		propSchema := generatePropertySchema(field.Type, bindingTag, visited)
		properties[jsonName] = propSchema

		// Check if field is required
		if strings.Contains(bindingTag, "required") {
			// For pointer types in UpdateDTO, don't mark as required
			if field.Type.Kind() != reflect.Ptr {
				required = append(required, jsonName)
			}
		}
	}

	schema := map[string]interface{}{
		"type":       "object",
		"properties": properties,
	}

	if len(required) > 0 {
		schema["required"] = required
	}

	return schema
}

// generatePropertySchema generates OpenAPI schema for a field type
func generatePropertySchema(t reflect.Type, bindingTag string, visited map[reflect.Type]bool) map[string]interface{} {
	// Handle pointer types
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	schema := make(map[string]interface{})

	switch t.Kind() {
	case reflect.String:
		schema["type"] = "string"
		// Check for email validation
		if strings.Contains(bindingTag, "email") {
			schema["format"] = "email"
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		schema["type"] = "integer"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		schema["type"] = "integer"
	case reflect.Float32, reflect.Float64:
		schema["type"] = "number"
	case reflect.Bool:
		schema["type"] = "boolean"
	case reflect.Slice, reflect.Array:
		schema["type"] = "array"
		schema["items"] = generatePropertySchema(t.Elem(), "", visited)
	case reflect.Struct:
		// Handle special types
		if t == reflect.TypeOf(time.Time{}) {
			schema["type"] = "string"
			schema["format"] = "date-time"
		} else {
			// Check for recursion BEFORE generating nested schema
			if visited[t] {
				schema["type"] = "object" // Break recursion
			} else {
				// Create new instance of struct to pass to generateSchemaWithRecursion
				// (Should ideally refactor to accept Type, but this works)
				newStruct := reflect.New(t).Interface()
				schema = generateSchemaWithRecursion(newStruct, visited)
			}
		}
	default:
		schema["type"] = "object"
	}

	// Add validation constraints from binding tag
	if bindingTag != "" {
		parts := strings.Split(bindingTag, ",")
		for _, part := range parts {
			if strings.HasPrefix(part, "min=") {
				schema["minimum"] = strings.TrimPrefix(part, "min=")
			} else if strings.HasPrefix(part, "max=") {
				schema["maximum"] = strings.TrimPrefix(part, "max=")
			} else if strings.HasPrefix(part, "len=") {
				schema["minLength"] = strings.TrimPrefix(part, "len=")
				schema["maxLength"] = strings.TrimPrefix(part, "len=")
			}
		}
	}

	return schema
}

// PathConfig holds configuration for custom path registration
type PathConfig struct {
	Path           string
	Method         string
	Summary        string
	Description    string
	Tags           []string
	RequestSchema  string
	ResponseSchema string
	StatusCode     string
	ErrorResponses map[string]string
	RequireAuth    bool
}

// SwaggerConfig holds API metadata configuration
type SwaggerConfig struct {
	Title       string
	Description string
	Version     string
	Host        string
	BasePath    string
}

// RegisterCustomPath registers a custom endpoint (non-CRUD) in the registry
func (r *SwaggerRegistry) RegisterCustomPath(config PathConfig) {
	r.mu.Lock()
	defer r.mu.Unlock()

	method := strings.ToLower(config.Method)
	if _, ok := r.Paths[config.Path]; !ok {
		r.Paths[config.Path] = make(map[string]interface{})
	}

	operation := map[string]interface{}{
		"summary":  config.Summary,
		"tags":     config.Tags,
		"consumes": []string{"application/json"},
		"produces": []string{"application/json"},
	}

	// Swagger 2.0 uses parameters in body for request bodies
	parameters := []map[string]interface{}{}

	// Add path parameters if any (simplified detection)
	if strings.Contains(config.Path, "{id}") {
		parameters = append(parameters, map[string]interface{}{
			"name":     "id",
			"in":       "path",
			"required": true,
			"type":     "string",
		})
	}

	// Add request body as a parameter (Swagger 2.0 format)
	if config.RequestSchema != "" {
		parameters = append(parameters, map[string]interface{}{
			"name":     "request",
			"in":       "body",
			"required": true,
			"schema": map[string]interface{}{
				"$ref": "#/definitions/" + config.RequestSchema,
			},
		})
	}

	if len(parameters) > 0 {
		operation["parameters"] = parameters
	}

	// Add responses (Swagger 2.0 format)
	responses := make(map[string]interface{})
	if config.ResponseSchema != "" {
		responses[config.StatusCode] = map[string]interface{}{
			"description": config.Description,
			"schema": map[string]interface{}{
				"allOf": []interface{}{
					map[string]interface{}{"$ref": "#/definitions/utils.Response"},
					map[string]interface{}{
						"type": "object",
						"properties": map[string]interface{}{
							"data": map[string]interface{}{
								"$ref": "#/definitions/" + config.ResponseSchema,
							},
						},
					},
				},
			},
		}
	} else {
		responses[config.StatusCode] = map[string]interface{}{
			"description": config.Description,
		}
	}

	// Add error responses
	if config.ErrorResponses != nil {
		for code, desc := range config.ErrorResponses {
			responses[code] = map[string]interface{}{
				"description": desc,
				"schema": map[string]interface{}{
					"$ref": "#/definitions/utils.Response",
				},
			}
		}
	}

	operation["responses"] = responses

	// Add security if required
	if config.RequireAuth {
		operation["security"] = []map[string]interface{}{
			{"BearerAuth": []interface{}{}},
		}
	}

	r.Paths[config.Path][method] = operation
}

// GenerateFullSpec generates a complete Swagger 2.0 specification
func (r *SwaggerRegistry) GenerateFullSpec(config SwaggerConfig) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	spec := map[string]interface{}{
		"swagger": "2.0",
		"info": map[string]interface{}{
			"title":          config.Title,
			"description":    config.Description,
			"version":        config.Version,
			"termsOfService": "http://swagger.io/terms/",
			"contact": map[string]interface{}{
				"name":  "API Support",
				"url":   "http://www.swagger.io/support",
				"email": "support@swagger.io",
			},
			"license": map[string]interface{}{
				"name": "Apache 2.0",
				"url":  "http://www.apache.org/licenses/LICENSE-2.0.html",
			},
		},
		"host":     config.Host,
		"basePath": config.BasePath,
		"schemes":  []string{},
		"securityDefinitions": map[string]interface{}{
			"BearerAuth": map[string]interface{}{
				"type":        "apiKey",
				"name":        "Authorization",
				"in":          "header",
				"description": "Enter your token directly or use \"Bearer <token>\" format.",
			},
		},
		"paths":       r.Paths,
		"definitions": r.Schemas,
	}

	merged, err := json.Marshal(spec)
	if err != nil {
		return "", err
	}

	return string(merged), nil
}
