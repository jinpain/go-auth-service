package sessionhttpv1

import "github.com/google/uuid"

type CreateSession struct {
	UserID    uuid.UUID `json:"user_id"`
	Device    string    `json:"device"`
	IpAddress string    `json:"ip_address"`
}

type Session struct {
	ID           uuid.UUID `json:"session_id"`
	UserID       uuid.UUID `json:"user_id"`
	RefreshToken uuid.UUID `json:"refresh_token"`
	AccessToken  uuid.UUID `json:"access_token"`
}
