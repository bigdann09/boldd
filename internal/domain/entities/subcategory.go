package entities

import (
	"github.com/google/uuid"
)

type SubCategory struct {
	ID         string
	Name       string
	CategoryID string
}

func NewSubCategory(name string, categoryID string) *SubCategory {
	return &SubCategory{
		ID:         uuid.NewString(),
		Name:       name,
		CategoryID: categoryID,
	}
}

func UpdateSubCategory(name string) *SubCategory {
	return &SubCategory{Name: name}
}
