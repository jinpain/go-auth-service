package verification

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

func (r *Repository) SaveVerificationCode(ctx context.Context, code, userID string) error {
	return r.redisClient.Set(ctx, code, userID, 15*time.Minute)
}

func (r *Repository) DeleteVerificationCode(ctx context.Context, code string) error {
	return r.redisClient.Del(ctx, code)
}

func (r *Repository) GetUserIDByVerificationCode(ctx context.Context, code string) (string, error) {
	return r.redisClient.Get(ctx, code)
}
