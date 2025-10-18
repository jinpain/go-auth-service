package token

type Model struct {
	SessionID    string `json:"session_id"`
	UserID       string `json:"user_id"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}
