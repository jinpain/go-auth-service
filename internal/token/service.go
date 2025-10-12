package token

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type TokenRepository interface {
	SetToken(ctx context.Context, sessionID, refresh string) error
	DeleteToken(ctx context.Context, sessionID string) error
	GetTokenBySessionID(ctx context.Context, sessionID string) (string, error)
}

type TokenProvider interface {
	GenerateToken(userID, sessionID string) (string, error)
}

type Service struct {
	tokenRepository TokenRepository
	tokenProvider   TokenProvider
}

func NewService(tokenRepository TokenRepository, tokenProvider TokenProvider) *Service {
	return &Service{
		tokenRepository: tokenRepository,
		tokenProvider:   tokenProvider,
	}
}

func (s *Service) SetToken(sessionID, refresh string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return s.tokenRepository.SetToken(ctx, sessionID, refresh)
}

func (s *Service) DeleteToken(sessionID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return s.tokenRepository.DeleteToken(ctx, sessionID)
}

func (s *Service) RefreshToken(token *Model) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	refresh, err := s.tokenRepository.GetTokenBySessionID(ctx, token.SessionID)
	if err != nil {
		return err
	}

	if token.RefreshToken != refresh {
		return errors.New("invalid refresh token")
	}

	token.RefreshToken = uuid.NewString()

	if err := s.tokenRepository.SetToken(ctx, token.SessionID, token.RefreshToken); err != nil {
		return err
	}

	token.AccessToken, err = s.tokenProvider.GenerateToken(token.UserID, token.SessionID)
	if err != nil {
		return err
	}

	return nil
}
