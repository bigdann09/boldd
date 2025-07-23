package attributes

type CreateAttributeRequest struct {
	Name string `json:"name" binding:"required"`
}
