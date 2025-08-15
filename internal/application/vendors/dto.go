package vendors

type CreateVendorRequest struct {
	Name            string `json:"name" binding:"required,min=6,max=100"`
	BusinessEmail   string `json:"business_email" binding:"required,email"`
	BusinessAddress string `json:"business_address" binding:"required"`
	BusinessPhone   string `json:"business_phone" binding:"required"`
	Description     string `json:"description" binding:"required,min=20,max=500"`
}
