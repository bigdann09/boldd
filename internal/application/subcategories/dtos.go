package subcategories

type CreateSubCategoryRequest struct {
	CategoryID string `json:"category_id" binding:"required"`
	Name       string `json:"name" binding:"required"`
}

type UpdateSubCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}
