package entities

type Attribute struct {
	ID   string
	Name string
}

func NewAttribute(name string) *Attribute {
	return &Attribute{Name: name}
}
