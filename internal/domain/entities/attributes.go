package entities

import (
	"strings"

	"github.com/google/uuid"
)

type Attribute struct {
	ID   string
	Name string
}

func NewAttribute(name string) *Attribute {
	return &Attribute{
		ID:   uuid.NewString(),
		Name: strings.ToLower(name),
	}
}

func UpdateAttribute(name string) *Attribute {
	return &Attribute{Name: strings.ToLower(name)}
}
