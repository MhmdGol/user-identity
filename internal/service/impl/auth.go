package service

import (
	"Identity/internal/model"
	"Identity/internal/repository"
	"Identity/internal/service"
	"Identity/pkg/bcrypthash"
	"Identity/pkg/jwt"
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type AuthService struct {
	userRepo    repository.UserRepo
	sessionRepo repository.SessionRepo
	jwt         *jwt.JwtToken
	redisClient *redis.Client
}

var _ service.AuthService = (*AuthService)(nil)

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (as *AuthService) Login(ctx context.Context, l model.LoginInfo) (model.JwtToken, error) {
	user, err := as.userRepo.ByUsername(l.Username)
	if err != nil {
		return "", err
	}

	err = bcrypthash.ValidatePassword(user.HashedPassword, l.Password)
	if err != nil {
		return "", err
	}

	session := model.Session{
		UserID:     user.ID,
		SessionExp: time.Now().Add(time.Hour),
	}

	err = as.sessionRepo.Add(session)
	if err != nil {
		return "", err
	}

	err = as.redisClient.Set(ctx, fmt.Sprint(user.ID), session, time.Hour).Err()
	if err != nil {
		return "", err
	}

	token, err := as.jwt.MakeToken(model.TokenClaim{
		ID: user.ID,
	})

	return token, err
}
