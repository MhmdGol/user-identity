package service

import (
	"Identity/internal/model"
	"Identity/internal/repository"
	"Identity/internal/service"
	"Identity/pkg/bcrypthash"
	"Identity/pkg/jwt"
	"Identity/pkg/smtp"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/pquerna/otp/totp"
	"github.com/redis/go-redis/v9"
)

type AuthService struct {
	userRepo    repository.UserRepo
	sessionRepo repository.SessionRepo
	trackRepo   repository.TrackRepo
	jwt         *jwt.JwtToken
	redisClient *redis.Client
}

var _ service.AuthService = (*AuthService)(nil)

func NewAuthService(
	u repository.UserRepo,
	s repository.SessionRepo,
	t repository.TrackRepo,
	jt *jwt.JwtToken,
	rc *redis.Client,
) *AuthService {
	return &AuthService{
		userRepo:    u,
		sessionRepo: s,
		trackRepo:   t,
		jwt:         jt,
		redisClient: rc,
	}
}

func (as *AuthService) Login(ctx context.Context, l model.LoginInfo) (model.JwtToken, error) {
	user, err := as.userRepo.ByUsername(ctx, l.Username)
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

	err = as.sessionRepo.Add(ctx, session)
	if err != nil {
		return "", err
	}

	sessionJSON, _ := json.Marshal(session)

	err = as.redisClient.Set(ctx, fmt.Sprint(user.ID), sessionJSON, time.Hour).Err()
	if err != nil {
		return "", err
	}

	token, err := as.jwt.MakeToken(model.TokenClaim{
		ID:       user.ID,
		Username: l.Username,
	})

	as.trackRepo.Create(ctx, model.TrackInfo{
		ID:        user.ID,
		Action:    "Login",
		Timestamp: time.Now(),
	})
	return token, err
}

func (as *AuthService) Logout(ctx context.Context, id model.ID) error {
	as.redisClient.Del(ctx, fmt.Sprint(id))

	as.sessionRepo.Remove(ctx, id)

	as.trackRepo.Create(ctx, model.TrackInfo{
		ID:        id,
		Action:    "Logout",
		Timestamp: time.Now(),
	})
	return nil
}

func (as *AuthService) UpdatePassword(ctx context.Context, up model.UpdatePassword) error {
	user, err := as.userRepo.ByUsername(ctx, up.Username)
	if err != nil {
		return err
	}

	err = bcrypthash.ValidatePassword(user.HashedPassword, up.OldPass)
	if err != nil {
		return err
	}

	hpass, err := bcrypthash.HashPassword(up.NewPass)
	if err != nil {
		return err
	}

	as.trackRepo.Create(ctx, model.TrackInfo{
		ID:        user.ID,
		Action:    "Pass Change",
		Timestamp: time.Now(),
	})

	return as.userRepo.UpdateByID(ctx, model.UserInfo{
		ID:             user.ID,
		HashedPassword: hpass,
	})
}

// pass recovery stages logic have some vulnerbilities
func (as *AuthService) PasswordRecovery(ctx context.Context, username string) error {
	user, err := as.userRepo.ByUsername(ctx, username)
	if err != nil {
		return err
	}

	as.trackRepo.Create(ctx, model.TrackInfo{
		ID:        user.ID,
		Action:    "Pass Recovery",
		Timestamp: time.Now(),
	})

	t, _ := as.jwt.MakeToken(model.TokenClaim{
		ID:       user.ID,
		Username: user.Username,
	})

	smtp.RunAndSendEmail("support@identity.com", user.Email, string(t))

	return nil
}

func (as *AuthService) ResetPassword(ctx context.Context, username string, password string) error {
	hpass, _ := bcrypthash.HashPassword(password)

	return as.userRepo.UpdateByUsername(ctx, model.UserInfo{
		Username:       username,
		HashedPassword: hpass,
	})
}

func (as *AuthService) TwoFactorAuth(ctx context.Context, l model.LoginInfo, code string) (model.JwtToken, error) {
	user, _ := as.userRepo.ByUsername(ctx, l.Username)

	isValid := totp.Validate(code, user.TotpSecret.String())
	if !isValid {
		return "", fmt.Errorf("login failed")
	}

	err := bcrypthash.ValidatePassword(user.HashedPassword, l.Password)
	if err != nil {
		return "", err
	}

	session := model.Session{
		UserID:     user.ID,
		SessionExp: time.Now().Add(time.Hour),
	}

	err = as.sessionRepo.Add(ctx, session)
	if err != nil {
		return "", err
	}

	sessionJSON, _ := json.Marshal(session)

	err = as.redisClient.Set(ctx, fmt.Sprint(user.ID), sessionJSON, time.Hour).Err()
	if err != nil {
		return "", err
	}

	token, err := as.jwt.MakeToken(model.TokenClaim{
		ID:       user.ID,
		Username: l.Username,
	})

	return token, err
}
