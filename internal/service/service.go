package service

import (
	"Identity/internal/model"
	"context"
)

type AuthService interface {
	Login(context.Context, model.LoginInfo) (model.JwtToken, error)
}
