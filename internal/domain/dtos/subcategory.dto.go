package dtos

type SubCategoryResponse struct {
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	CategoryID uint   `json:"category_id"`
}

type SubCategoryQueryFilter struct {
	Page     int    `form:"page" binding:"number"`
	PageSize int    `form:"page_size" binding:"number"`
	SortBy   string `form:"sort_by" binding:""`
	Order    string `form:"order" binding:""`
}
