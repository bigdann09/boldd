package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	UUID string
	Name string
}

func NewCategory(name string) *Category {
	return &Category{
		UUID: uuid.NewString(),
		Name: name,
	}
}
