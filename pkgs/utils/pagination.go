package utils

import (
	"gorm.io/gorm"
)

type PaginationResponse[Entity any] struct {
	Items      []Entity `json:"data"`
	Total      int      `json:"total"`
	Page       int      `json:"page"`
	PageSize   int      `json:"page_size"`
	TotalPages int      `json:"total_pages"`
}

func NewPaginationResponse[Entity any](page, limit int, db *gorm.DB) (PaginationResponse[Entity], error) {
	var response []Entity

	var total int64
	clone := db.Session(&gorm.Session{})
	clone.Count(&total)

	offset := (page - 1) * int(limit)
	result := db.Offset(offset).Limit(limit).Scan(&response)

	totalPages := 0
	if total > 0 {
		totalPages = int(total) / limit
	}

	if len(response) == 0 {
		response = []Entity{}
	}

	return PaginationResponse[Entity]{
		Total:      int(total),
		Items:      response,
		Page:       page,
		PageSize:   limit,
		TotalPages: totalPages,
	}, result.Error
}

func (pagination PaginationResponse[Entity]) PaginationItems() []Entity {
	return pagination.Items
}

func (pagination PaginationResponse[Entity]) TotalItems() int {
	return pagination.Total
}
