package attributes

type CreateAttributeRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateAttributeRequest struct {
	Name string `json:"name" binding:"required"`
}
