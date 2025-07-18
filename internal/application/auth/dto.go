package auth

type RegisterRequest struct {
	FullName    string `json:"fullname" binding:"required,min=6,max=60"`
	Email       string `json:"email" binding:"required,email,unique_email"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
