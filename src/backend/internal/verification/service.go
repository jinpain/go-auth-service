package verification

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type VerificationRepository interface {
	SaveVerificationCode(ctx context.Context, code, userID string) error
	DeleteVerificationCode(ctx context.Context, code string) error
	GetUserIDByVerificationCode(ctx context.Context, code string) (string, error)
}

type UserService interface {
	SetUserVerified(userID uuid.UUID) error
}

type Service struct {
	verificationRepository VerificationRepository
	userService            UserService
}

func NewService(verificationRepository VerificationRepository, userService UserService) *Service {
	return &Service{
		verificationRepository: verificationRepository,
		userService:            userService,
	}
}

func (s *Service) SaveVerificationCode(code, userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return s.verificationRepository.SaveVerificationCode(ctx, code, userID)
}

func (s *Service) DeleteVerificationCode(code string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return s.verificationRepository.DeleteVerificationCode(ctx, code)
}

func (s *Service) CheckVerifyCode(code string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	userID, err := s.verificationRepository.GetUserIDByVerificationCode(ctx, code)
	if err != nil {
		return err
	} else if userID == "" {
		return errors.New("verify code not found")
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return err
	}

	if err := s.userService.SetUserVerified(userUUID); err != nil {
		return err
	}

	if err := s.DeleteVerificationCode(code); err != nil {
		return err
	}

	return nil
}
