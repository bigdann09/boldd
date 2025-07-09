package auth

type RegisterRequest struct {
	FullName    string `json:"fullname" binding:"required,min=15,max=100"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phonenumber" binding:"required"`
	Password    string `json:"password" binding:"required,min=6"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
