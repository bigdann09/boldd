package utils

import (
	"gorm.io/gorm"
)

type PaginationResponse[Entity any] struct {
	Items      []Entity
	Total      int
	Page       int
	PageSize   int
	TotalPages int
}

func NewPaginationResponse[Entity any](page, limit int, db *gorm.DB) (PaginationResponse[Entity], error) {
	var response []Entity

	var total int64
	clone := db.Session(&gorm.Session{})
	clone.Count(&total)

	offset := (page - 1) * int(limit)
	result := db.Offset(offset).Limit(limit).Scan(&response)

	return PaginationResponse[Entity]{
		Total:      int(total),
		Items:      response,
		Page:       page,
		PageSize:   limit,
		TotalPages: int(total) / limit,
	}, result.Error
}

func (pagination PaginationResponse[Entity]) PaginationItems() []Entity {
	return pagination.Items
}

func (pagination PaginationResponse[Entity]) TotalItems() int {
	return pagination.Total
}
