package user

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *Model) error
	GetUserByEmail(ctx context.Context, email string) (*Model, error)
	ExistsUserByEmail(ctx context.Context, email string) (bool, error)
	SetUserVerified(ctx context.Context, userID uuid.UUID) error
}

type Service struct {
	userRepository UserRepository
}

func NewService(userRepository UserRepository) *Service {
	return &Service{
		userRepository: userRepository,
	}
}

func (s *Service) CreateUser(user *Model) error {
	existing, err := s.ExistsUserByEmail(user.Email)
	if err != nil {
		return err
	}

	if existing {
		return errors.New("email already registered")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = s.userRepository.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetUserByEmail(email string) (*Model, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	user, err := s.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) ExistsUserByEmail(email string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	exists, err := s.userRepository.ExistsUserByEmail(ctx, email)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (s *Service) SetUserVerified(userID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return s.userRepository.SetUserVerified(ctx, userID)
}
