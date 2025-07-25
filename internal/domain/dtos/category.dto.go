package dtos

type CategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CategoryQueryFilter struct {
	Page     int    `form:"page" binding:"number"`
	PageSize int    `form:"page_size" binding:"number"`
	SortBy   string `form:"sort_by" binding:""`
	Order    string `form:"order" binding:""`
}
