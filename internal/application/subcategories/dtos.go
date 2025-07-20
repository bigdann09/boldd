package subcategories

type CreateSubCategoryRequest struct {
	CategoryID uint   `json:"category_id" binding:"required,number"`
	Name       string `json:"name" binding:"required"`
}

type UpdateSubCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}
