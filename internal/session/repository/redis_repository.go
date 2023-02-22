package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/sayuen0/go-to-gym/config"
	"github.com/sayuen0/go-to-gym/internal/models"
	"github.com/sayuen0/go-to-gym/internal/session"
	"github.com/sayuen0/go-to-gym/pkg/utils"
	"time"
)

type sessionRepo struct {
	redisClient *redis.Client
	cfg         *config.Config
}

func NewSessionRepository(redisClient *redis.Client, cfg *config.Config) session.Repository {
	return &sessionRepo{
		redisClient: redisClient, cfg: cfg,
	}
}

func (s *sessionRepo) CreateSession(ctx context.Context, sess *models.Session, expiration int) (string, error) {
	sess.SessionID = utils.NewUUIDStr()
	sessionKey := s.createKey(sess.SessionID)

	sessBytes, err := json.Marshal(&sess)
	if err != nil {
		return "", errors.Wrap(err, "sessionRepo.CreateSession.json.Marshal")
	}
	if err := s.redisClient.Set(ctx, sessionKey, sessBytes, time.Second*time.Duration(expiration)).Err(); err != nil {
		return "", errors.Wrap(err, "sessionRepo.CreateSession.redisClient.Set")
	}
	return sessionKey, nil
}

func (s *sessionRepo) DeleteByID(ctx context.Context, id string) error {
	if err := s.redisClient.Del(ctx, id).Err(); err != nil {
		return errors.Wrap(err, "sessionRepo.DeleteByID")
	}
	return nil
}

func (s *sessionRepo) createKey(sessionID string) string {
	return fmt.Sprintf("%s: %s", s.cfg.Session.Prefix, sessionID)
}
