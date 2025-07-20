package dtos

type CategoryResponse struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type CategoryQueryFilter struct {
	Page     int    `json:"page" binding:"number,isdefault=1"`
	PageSize int    `json:"page_size" binding:"number,isdefault=10"`
	SortBy   string `json:"sort_by" binding:""`
	Order    string `json:"order" binding:""`
}
