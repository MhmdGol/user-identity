package repository

import "Identity/internal/model"

type SessionRepo interface {
	Add(model.Session) error
	Remove(model.ID) error
	UpdateByID(model.ID) error
	ByID(model.ID) (model.Session, error)
}

type UserRepo interface {
	Create(model.RawUser) error
	ByID(model.ID) (model.UserInfo, error)
	ByUsername(string) (model.UserInfo, error)
}
