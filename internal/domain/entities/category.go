package entities

import (
	"github.com/google/uuid"
)

type Category struct {
	ID   string
	Name string
}

func NewCategory(name string) *Category {
	return &Category{
		ID:   uuid.NewString(),
		Name: name,
	}
}

func UpdateCategory(name string) *Category {
	return &Category{Name: name}
}
