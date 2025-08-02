package products

type CreateProductCategory struct {
	Name        string           `json:"name" binding:"required,min=10"`
	Description string           `json:"description" binding:"required,min=30"`
	Price       float64          `json:"price" binding:"required,number"`
	Variants    []ProductVariant `json:"variants" binding:"required"`
}

type ProductVariant struct {
	AttributeID    string      `json:"attribute_id" binding:"required"`
	AttributeValue interface{} `json:"attribute_vaule" binding:"required"`
	BasePrice      float64     `json:"base_price" binding:"required"`
	AdjustedPrice  float64     `json:"adjusted_price" binding:"required"`
	Quantity uint `form:"quantity"`
}
