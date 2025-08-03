package products

type CreateProductRequest struct {
	Name        string           `json:"name" binding:"required,min=10"`
	Description string           `json:"description" binding:"required,min=30"`
	Price       *float64         `json:"price" binding:"number"`
	Variants    []ProductVariant `json:"variants" binding:"gt=0"`
	Images      []string         `json:"images" binding:"required,gt=0,dive,string"`
}

type ProductVariant struct {
	Price      float64     `json:"price" binding:"required,numeric"`
	BasePrice  *float64    `json:"base_price" binding:"numeric"`
	Stock      uint        `json:"quantity" binding:"required,number"`
	Attributes []Attribute `json:"attribute" binding:"required,gt=0"`
	Images     []string    `json:"images" binding:"required,gt=0,dive,string"`
}

type Attribute struct {
	AttributeID string   `json:"attribute_id" binding:"required,number"`
	Value       []string `json:"value" binding:"required,gt=0,dive,string"`
}
