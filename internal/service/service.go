package service

import (
	"Identity/internal/model"
	"context"
)

type AuthService interface {
	Login(context.Context, model.LoginInfo) (model.JwtToken, error)
	Logout(context.Context, model.ID) error
	UpdatePassword(context.Context, model.UpdatePassword) error
	PasswordRecovery(context.Context, string) error
	ResetPassword(context.Context, string, string) error
	TwoFactorAuth(context.Context, model.LoginInfo, int) (model.JwtToken, error)
}

type UserService interface {
	Create(context.Context, model.RawUser) error
	ReadByID(context.Context, model.ID) (model.UserInfo, error)
	ChangePermission(context.Context, string, string) error
	UpdateByID(context.Context, model.ID, string) error
}

type SessionService interface {
	Add(context.Context, model.Session) error
	Remove(context.Context, model.ID) error
	CheckSession(context.Context, model.ID) error
}
