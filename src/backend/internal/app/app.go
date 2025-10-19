package app

import (
	"fmt"

	"github.com/jinpain/go-auth-service/internal/auth"
	authhttpv1 "github.com/jinpain/go-auth-service/internal/auth/transport/http/v1"
	"github.com/jinpain/go-auth-service/internal/config"
	"github.com/jinpain/go-auth-service/internal/session"
	sessionhttpv1 "github.com/jinpain/go-auth-service/internal/session/transport/http/v1"
	"github.com/jinpain/go-auth-service/internal/token"
	tokenhttpv1 "github.com/jinpain/go-auth-service/internal/token/transport/http/v1"
	"github.com/jinpain/go-auth-service/internal/user"
	"github.com/jinpain/go-auth-service/internal/verification"
	verificationhttpv1 "github.com/jinpain/go-auth-service/internal/verification/transport/http/v1"
	"github.com/jinpain/go-auth-service/pkg/database/postgres"
	"github.com/jinpain/go-auth-service/pkg/database/redis"
	"github.com/jinpain/go-auth-service/pkg/httpserver"
	"github.com/jinpain/go-auth-service/pkg/logger"
	"github.com/jinpain/go-auth-service/pkg/middleware"
	"github.com/jinpain/go-auth-service/pkg/sqlstore"
)

func Run(cfg *config.Config) error {
	log := logger.New(cfg.ServiceConfig.Env)

	log.Info("Application running!")

	pgConnStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		cfg.DatabaseConfig.Host, cfg.DatabaseConfig.Port, cfg.DatabaseConfig.User,
		cfg.DatabaseConfig.Password, cfg.DatabaseConfig.Name)
	pgxAdapter, err := postgres.NewPgxAdapter(pgConnStr)
	if err != nil {
		return err
	}

	sqlStore, err := sqlstore.New(cfg.SqlConfig.Path)
	if err != nil {
		return err
	}

	redisToken, err := redis.New(&cfg.RedisTokenConfig)
	if err != nil {
		return err
	}

	redisVerification, err := redis.New(&cfg.RedisVerificationConfig)
	if err != nil {
		return err
	}

	httpServer := httpserver.New(cfg.ServiceConfig.Env)

	// api groups
	publicApi := httpServer.Group("/api")
	protectedApi := httpServer.Group("/api")

	// middlewares
	protectedApi.Use(middleware.JWTMiddleware(cfg.JWTConfig.Secret))

	// repositories
	tokenRepository := token.NewRepository(redisToken)
	userRepository := user.NewRepository(pgxAdapter, sqlStore)
	sessionRepository := session.NewRepository(pgxAdapter, sqlStore)
	verificationRepository := verification.NewRepository(redisVerification)

	// services
	jwtService := token.NewJWTService(cfg.JWTConfig.Secret)
	tokenService := token.NewService(tokenRepository, jwtService)
	userService := user.NewService(userRepository)
	sessionService := session.NewService(sessionRepository, tokenService)
	verificationService := verification.NewService(verificationRepository, userService)
	authService := auth.NewService(userService, sessionService, verificationService, tokenService, jwtService)

	// routers public
	authhttpv1.SetupRouter(publicApi, authService)
	verificationhttpv1.SetupRouter(publicApi, verificationService)

	// routes protected
	sessionhttpv1.SetupRouter(protectedApi, sessionService)
	tokenhttpv1.SetupRouter(protectedApi, tokenService)

	addr := fmt.Sprintf("%s:%d", cfg.ServerConfig.Host, cfg.ServerConfig.Port)
	if err := httpServer.Run(addr); err != nil {
		return err
	}

	return nil
}
