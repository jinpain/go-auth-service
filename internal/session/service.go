package session

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type TokenService interface {
	DeleteToken(sessionID string) error
}

type SessionRepository interface {
	CreateSession(ctx context.Context, session *Model) error
	RevokeSession(ctx context.Context, userID, sessionID uuid.UUID) error
}

type Service struct {
	sessionRepository SessionRepository
	tokenService      TokenService
}

func NewService(sessionRepository SessionRepository, tokenService TokenService) *Service {
	return &Service{
		sessionRepository: sessionRepository,
		tokenService:      tokenService,
	}
}

func (s *Service) CreateSession(session *Model) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := s.sessionRepository.CreateSession(ctx, session)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) RevokeSession(userID, sessionID string) error {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return err
	}

	sessionUUID, err := uuid.Parse(sessionID)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.sessionRepository.RevokeSession(ctx, userUUID, sessionUUID); err != nil {
		return err
	}

	if err := s.tokenService.DeleteToken(sessionID); err != nil {
		return err
	}

	return nil
}
