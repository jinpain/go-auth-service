package token

import (
	"context"
	"time"

	"github.com/jinpain/go-auth-service/pkg/database/redis"
)

type Repository struct {
	redisClient *redis.RedisClient
}

func NewRepository(redisClient *redis.RedisClient) *Repository {
	return &Repository{
		redisClient: redisClient,
	}
}

func (r *Repository) SetToken(ctx context.Context, sessionID, refresh string) error {
	return r.redisClient.Set(ctx, sessionID, refresh, 15*time.Minute)
}

func (r *Repository) DeleteToken(ctx context.Context, sessionID string) error {
	return r.redisClient.Del(ctx, sessionID)
}

func (r *Repository) GetTokenBySessionID(ctx context.Context, sessionID string) (string, error) {
	return r.redisClient.Get(ctx, sessionID)
}
