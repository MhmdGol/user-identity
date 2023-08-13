package service

import (
	"Identity/internal/model"
	"Identity/internal/repository"
	"Identity/internal/service"
	"Identity/pkg/jwt"
	"context"

	"github.com/redis/go-redis/v9"
)

type UserService struct {
	userRepo    repository.UserRepo
	sessionRepo repository.SessionRepo
	jwt         *jwt.JwtToken
	redisClient *redis.Client
}

var _ service.UserService = (*UserService)(nil)

func NewUserService() *AuthService {
	return &AuthService{}
}

func (us *UserService) Create(ctx context.Context, ru model.RawUser) error {

}
