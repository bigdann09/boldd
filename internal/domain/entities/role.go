package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	UUID string
	Name string
}

func NewRole(name string) *Role {
	return &Role{
		UUID: uuid.NewString(),
		Name: name,
	}
}
