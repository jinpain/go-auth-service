package auth

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jinpain/go-auth-service/internal/session"
	"github.com/jinpain/go-auth-service/internal/user"
	"github.com/jinpain/go-auth-service/pkg/util"
)

type UserService interface {
	CreateUser(user *user.Model) error
	GetUserByEmail(email string) (*user.Model, error)
	ExistsUserByEmail(email string) (bool, error)
}

type SessionService interface {
	CreateSession(session *session.Model) error
}

type VerificationService interface {
	SaveVerificationCode(code, userID string) error
}

type TokenService interface {
	SetToken(sessionID, refreshToken string) error
}

type TokenProvider interface {
	GenerateToken(userID, sessionID string) (string, error)
}

type Service struct {
	userService         UserService
	sessionService      SessionService
	verificationService VerificationService
	tokenService        TokenService
	tokenProvider       TokenProvider
}

func NewService(userService UserService, sessionService SessionService,
	verificationService VerificationService, tokenService TokenService,
	tokenProvider TokenProvider) *Service {
	return &Service{
		userService:         userService,
		sessionService:      sessionService,
		verificationService: verificationService,
		tokenService:        tokenService,
		tokenProvider:       tokenProvider,
	}
}

func (s *Service) Register(user *user.Model) error {
	if err := ValidatePassword(user.Password); err != nil {
		return err
	}

	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	if err := s.userService.CreateUser(user); err != nil {
		return err
	}

	code := uuid.NewString()

	if err := s.verificationService.SaveVerificationCode(code, user.ID.String()); err != nil {
		return err
	}

	return nil
}

func (s *Service) Login(email, password string, session *session.Model) (*Model, error) {
	userExists, err := s.userService.ExistsUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if !userExists {
		return nil, errors.New("account not found")
	}

	user, err := s.userService.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if user.Blocked {
		return nil, errors.New("user is blocked")
	}

	if !util.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid password")
	}

	session.UserID = user.ID

	err = s.sessionService.CreateSession(session)
	if err != nil {
		return nil, errors.New("invalid password")
	}

	jwt, err := s.tokenProvider.GenerateToken(session.UserID.String(), session.ID.String())
	if err != nil {
		return nil, errors.New("the service is temporarily unavailable. Please try again later")
	}

	refreshToken := uuid.NewString()

	if err := s.tokenService.SetToken(session.ID.String(), refreshToken); err != nil {
		return nil, errors.New("the service is temporarily unavailable. Please try again later")
	}

	authTokens := Model{
		AccessToken:  jwt,
		RefreshToken: refreshToken,
	}

	return &authTokens, nil
}
