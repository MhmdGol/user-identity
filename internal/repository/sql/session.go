package sql

import (
	"Identity/internal/model"
	"Identity/internal/repository"
	"Identity/internal/repository/sql/sqlmodel"
	"context"
	"fmt"

	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
)

type SessionRepo struct {
	db *gorm.DB
}

var _ repository.SessionRepo = (*SessionRepo)(nil)

func NewSessionRepo(db *gorm.DB) *SessionRepo {
	return &SessionRepo{
		db: db,
	}
}

func (sr *SessionRepo) Add(ctx context.Context, s model.Session) error {
	session := sqlmodel.Session{
		UserID:     fmt.Sprint(s.UserID),
		SessionExp: s.SessionExp,
	}

	err := sr.db.WithContext(ctx).Create(&session).Error
	if err != nil {
		return err
	}

	return nil
}

func (sr *SessionRepo) Remove(ctx context.Context, id model.ID) error {
	var session sqlmodel.Session
	err := sr.db.WithContext(ctx).Find(&session).Error
	if err != nil {
		return err
	}

	err = sr.db.WithContext(ctx).Delete(&session).Error
	if err != nil {
		return err
	}

	return nil
}

func (sr *SessionRepo) ByID(ctx context.Context, id model.ID) (model.Session, error) {
	var session sqlmodel.Session
	err := sr.db.WithContext(ctx).Where(&sqlmodel.Session{UserID: fmt.Sprint(id)}).First(&session).Error
	if err != nil {
		return model.Session{}, err
	}

	sId, _ := snowflake.ParseString(session.UserID)

	return model.Session{
		UserID:     model.ID(sId),
		SessionExp: session.SessionExp,
	}, nil
}
