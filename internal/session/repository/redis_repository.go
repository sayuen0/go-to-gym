package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/sayuen0/go-to-gym/config"
	"github.com/sayuen0/go-to-gym/internal/models"
	"github.com/sayuen0/go-to-gym/internal/session"
	"github.com/sayuen0/go-to-gym/pkg/utils"
)

type sessionRepo struct {
	redisClient *redis.Client
	cfg         *config.Config
}

// NewSessionRepo is a constructor of session.Repository
func NewSessionRepo(redisClient *redis.Client, cfg *config.Config) session.Repository {
	return &sessionRepo{
		redisClient: redisClient, cfg: cfg,
	}
}

// CreateSession creates a new session
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

// GetSessionById returns a session
func (s *sessionRepo) GetSessionByID(ctx context.Context, id string) (*models.Session, error) {
	sessBytes, err := s.redisClient.Get(ctx, id).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "sessionRepo.GetSessionByID.redisClient.Get")
	}

	sess := &models.Session{}
	if err := json.Unmarshal(sessBytes, sess); err != nil {
		return nil, errors.Wrap(err, "sessionRepo.GetSessionByID.json.Unmarshal")
	}

	return sess, nil
}

// DeleteByID deletes the session
func (s *sessionRepo) DeleteByID(ctx context.Context, id string) error {
	if err := s.redisClient.Del(ctx, id).Err(); err != nil {
		return errors.Wrap(err, "sessionRepo.DeleteByID")
	}

	return nil
}

func (s *sessionRepo) createKey(sessionID string) string {
	return fmt.Sprintf("%s: %s", s.cfg.Session.Prefix, sessionID)
}
