package service

import (
	"Identity/internal/model"
	"Identity/internal/repository"
	"Identity/internal/service"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type SessionService struct {
	sessionRepo repository.SessionRepo
	redisClient *redis.Client
}

var _ service.SessionService = (*SessionService)(nil)

func NewSessionService() *SessionService {
	return &SessionService{}
}

func (ss *SessionService) CheckSession(ctx context.Context, id model.ID) error {
	sessionJSON, err := ss.redisClient.Get(ctx, fmt.Sprint(id)).Result()
	if err == redis.Nil {
		session, err := ss.sessionRepo.ByID(ctx, id)
		if err == nil {
			return fmt.Errorf("session doesn't exist")
		}
		if time.Now().UTC().After(session.SessionExp) {
			ss.sessionRepo.Remove(ctx, id)
			return fmt.Errorf("session expired")
		}
		return nil
	}

	var retrievedSession model.Session
	json.Unmarshal([]byte(sessionJSON), &retrievedSession)

	if time.Now().UTC().After(retrievedSession.SessionExp) {
		return fmt.Errorf("session expired")
	}

	return nil

}
