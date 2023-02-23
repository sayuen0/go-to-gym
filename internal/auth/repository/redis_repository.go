package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/sayuen0/go-to-gym/internal/auth"
	"github.com/sayuen0/go-to-gym/internal/models"
)

const (
	basePrefix = "api-auth:"
)

type authRedisRepo struct {
	cli *redis.Client
}

// NewRedisRepo is a constructor of auth.RedisRepository
func NewRedisRepo(redisClient *redis.Client) auth.RedisRepository {
	return &authRedisRepo{
		cli: redisClient,
	}
}

// GetByID return a models.User found by a key with userID
func (r *authRedisRepo) GetByID(ctx context.Context, userID string) (*models.User, error) {
	// TODO: Redis専用モデルを作成すべき
	userBytes, err := r.cli.Get(ctx, r.generateUserKey(userID)).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "authRedisRepo.GetByID.redis.Client.Get")
	}

	user := &models.User{}
	if err := json.Unmarshal(userBytes, user); err != nil {
		return nil, errors.Wrap(err, "authRedisRepo.GetByID.json.Unmarshal")
	}

	return user, nil
}

// SetUser caches a user with duration in seconds
func (r *authRedisRepo) SetUser(ctx context.Context, userID string, duration int, user *models.User) error {
	userBytes, err := json.Marshal(user)
	if err != nil {
		return errors.Wrap(err, "authRedisRepo.SetUser.json.Marshal")
	}

	if err := r.cli.Set(ctx, r.generateUserKey(userID),
		userBytes, time.Second*time.Duration(duration)).Err(); err != nil {
		return errors.Wrap(err, "authRedisRepo.SetUser.redis.Client.Set")
	}
	return nil
}

func (r *authRedisRepo) generateUserKey(userID string) string {
	return fmt.Sprintf("%s: %s", basePrefix, userID)
}
