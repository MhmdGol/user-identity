package repository

import (
	"Identity/internal/model"
	"context"
)

type SessionRepo interface {
	Add(context.Context, model.Session) error
	Remove(context.Context, model.ID) error
	UpdateByID(context.Context, model.ID) error
	ByID(context.Context, model.ID) (model.Session, error)
}

type UserRepo interface {
	Create(context.Context, model.RawUser) error
	ByID(context.Context, model.ID) (model.UserInfo, error)
	ByUsername(context.Context, string) (model.UserInfo, error)
}

type TrackRepo interface {
	Create(context.Context, model.TrackInfo) error
	All(context.Context) (model.TrackInfo, error)
}
