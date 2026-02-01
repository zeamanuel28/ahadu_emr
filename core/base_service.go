package core

import (
	"errors"
	"math"
	"reflect"
	"strings"

	"gorm.io/gorm"
)

type FilterParam struct {
	Field    string      `json:"field"`
	Operator string      `json:"operator"` // eq, neq, gt, gte, lt, lte, like, ilike, in
	Value    interface{} `json:"value"`
}

type SortParam struct {
	Field string `json:"field"`
	Order string `json:"order"` // asc, desc
}

type QueryParams struct {
	Page         int
	Limit        int
	Search       string
	SearchFields []string
	Filters      []FilterParam
	Sorts        []SortParam
	Includes     []string
}

type BaseService[T any] struct {
	DB               *gorm.DB
	Preloads         []string // Default fields to preload on Get/GetAll
	SearchableFields []string // Default fields to search if none specified
}

func NewBaseService[T any](db *gorm.DB) *BaseService[T] {
	return &BaseService[T]{
		DB: db,
	}
}

func (s *BaseService[T]) GetAll(params QueryParams) ([]T, int64, int, error) {
	var results []T
	var total int64

	db := s.DB.Model(new(T))

	// 1. Automatic Preloading from Service Config
	for _, preload := range s.Preloads {
		db = db.Preload(preload)
	}

	// 2. Complex Filtering
	for _, filter := range params.Filters {
		column := s.getDBColumnName(filter.Field)
		switch filter.Operator {
		case "eq":
			db = db.Where(column+" = ?", filter.Value)
		case "neq":
			db = db.Where(column+" <> ?", filter.Value)
		case "gt":
			db = db.Where(column+" > ?", filter.Value)
		case "gte":
			db = db.Where(column+" >= ?", filter.Value)
		case "lt":
			db = db.Where(column+" < ?", filter.Value)
		case "lte":
			db = db.Where(column+" <= ?", filter.Value)
		case "like":
			db = db.Where(column+" LIKE ?", "%"+filter.Value.(string)+"%")
		case "ilike":
			db = db.Where(column+" ILIKE ?", "%"+filter.Value.(string)+"%")
		case "in":
			db = db.Where(column+" IN ?", filter.Value)
		}
	}

	// Legacy soft delete filter
	db = db.Where("is_deleted = ?", false)

	// 3. Search
	if params.Search != "" {
		searchQuery := "%" + params.Search + "%"
		searchFields := params.SearchFields
		if len(searchFields) == 0 {
			searchFields = s.SearchableFields
		}

		if len(searchFields) > 0 {
			searchCondition := ""
			searchArgs := []interface{}{}
			for i, field := range searchFields {
				if i > 0 {
					searchCondition += " OR "
				}
				searchCondition += s.getDBColumnName(field) + " ILIKE ?"
				searchArgs = append(searchArgs, searchQuery)
			}
			db = db.Where(searchCondition, searchArgs...)
		}
	}

	// 4. Count total
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, 0, err
	}

	// 5. Sorting
	if len(params.Sorts) > 0 {
		for _, sort := range params.Sorts {
			column := s.getDBColumnName(sort.Field)
			order := "ASC"
			if sort.Order == "desc" || sort.Order == "DESC" {
				order = "DESC"
			}
			db = db.Order(column + " " + order)
		}
	}

	// 6. Pagination
	if params.Page <= 0 {
		params.Page = 1
	}
	if params.Limit <= 0 {
		params.Limit = 10
	}
	offset := (params.Page - 1) * params.Limit

	err := db.Limit(params.Limit).Offset(offset).Find(&results).Error
	totalPages := int(math.Ceil(float64(total) / float64(params.Limit)))

	return results, total, totalPages, err
}

// getDBColumnName maps JSON field name to DB column name using reflection & GORM naming
func (s *BaseService[T]) getDBColumnName(field string) string {
	t := reflect.TypeOf(new(T)).Elem()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		jsonTag := f.Tag.Get("json")
		// Check if json tag matches field (simple split)
		if strings.Split(jsonTag, ",")[0] == field {
			// Found match, check gorm tag for column name override
			gormTag := f.Tag.Get("gorm")
			parts := strings.Split(gormTag, ";")
			for _, p := range parts {
				if strings.HasPrefix(p, "column:") {
					return strings.TrimPrefix(p, "column:")
				}
			}
			// Fallback to NamingStrategy
			return s.DB.NamingStrategy.ColumnName("", f.Name)
		}
	}
	// If no match found, assume it is already a column name or unknown
	return field
}

func (s *BaseService[T]) GetById(id string) (*T, error) {
	var result T
	db := s.DB.Model(new(T)).Where("id = ? AND is_deleted = ?", id, false)

	// Automatic Preloading
	for _, preload := range s.Preloads {
		db = db.Preload(preload)
	}

	err := db.First(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}

func (s *BaseService[T]) Create(data *T) error {
	return s.DB.Create(data).Error
}

func (s *BaseService[T]) Update(id string, data interface{}) (*T, error) {
	var result T
	if err := s.DB.First(&result, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if err := s.DB.Model(&result).Updates(data).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *BaseService[T]) Patch(id string, data interface{}) (*T, error) {
	var result T
	if err := s.DB.First(&result, "id = ?", id).Error; err != nil {
		return nil, err
	}

	if err := s.DB.Model(&result).Updates(data).Error; err != nil {
		return nil, err
	}

	return &result, nil
}
