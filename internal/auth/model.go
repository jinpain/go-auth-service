package auth

type Model struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}
