package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubCategory struct {
	gorm.Model
	UUID       string
	Name       string
	CategoryID uint
}

func NewSubCategory(name string, categoryID uint) *SubCategory {
	return &SubCategory{
		UUID:       uuid.NewString(),
		Name:       name,
		CategoryID: categoryID,
	}
}

func UpdateSubCategory(name string) *SubCategory {
	return &SubCategory{Name: name}
}
