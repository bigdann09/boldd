package dtos

type VendorResponse struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	BusinessEmail   string `json:"business_email"`
	BusinessAddress string `json:"business_address"`
	BusinessPhone   string `json:"business_phone"`
	Description     string `json:"description"`
}

type VendorQueryFilter struct {
	Page     int    `form:"page" binding:"number"`
	PageSize int    `form:"page_size" binding:"number"`
	SortBy   string `form:"sort_by" binding:""`
	Order    string `form:"order" binding:""`
}
