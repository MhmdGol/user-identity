package repository

import (
	"Identity/internal/model"
	"context"
)

type SessionRepo interface {
	Add(context.Context, model.Session) error
	Remove(context.Context, model.ID) error
	ByID(context.Context, model.ID) (model.Session, error)
}

type UserRepo interface {
	Create(context.Context, model.RawUser) error
	ByID(context.Context, model.ID) (model.UserInfo, error)
	ByUsername(context.Context, string) (model.UserInfo, error)
	UpdateByID(context.Context, model.UserInfo) error
	UpdateByUsername(context.Context, model.UserInfo) error
}

type TrackRepo interface {
	Create(context.Context, model.TrackInfo) error
}

// goose -dir ./internal/migrations sqlserver "sqlserver://sa:MyPass1234@localhost:1433?database=identityDB" create initial_migration sql
// goose -dir ./internal/migrations sqlserver "sqlserver://sa:MyPass1234@localhost:1433?database=identityDB" up
