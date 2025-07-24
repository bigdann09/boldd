package entities

import (
	"strings"

	"github.com/google/uuid"
)

type Category struct {
	ID   string
	Name string
}

func NewCategory(name string) *Category {
	return &Category{
		ID:   uuid.NewString(),
		Name: strings.ToLower(name),
	}
}

func UpdateCategory(name string) *Category {
	return &Category{Name: strings.ToLower(name)}
}
