package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secret string
}

func NewJWTService(secret string) *JWTService {
	return &JWTService{
		secret: secret,
	}
}

func (s *JWTService) GenerateToken(userID, sessionID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"sid": sessionID,
		"exp": time.Now().Add(15 * time.Minute).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secret))
}
