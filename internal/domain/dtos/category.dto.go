package dtos

type CategoryResponse struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type CategoryQueryFilter struct {
	Page     int    `json:"page" binding:"required"`
	PageSize int    `json:"page_size" binding:"required"`
	SortBy   string `json:"sort_by" binding:"one_of=name,created_at"`
	Order    string `json:"order" binding:"one_of=asc,desc"`
}
