package authhttpv1

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}
