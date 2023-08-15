package sql

import (
	"Identity/internal/model"
	"Identity/internal/repository"
	"Identity/internal/repository/sql/sqlmodel"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type TrackRepo struct {
	db *gorm.DB
}

var _ repository.TrackRepo = (*TrackRepo)(nil)

func NewTrackRepo(db *gorm.DB) *TrackRepo {
	return &TrackRepo{
		db: db,
	}
}

func (tr *TrackRepo) Create(ctx context.Context, t model.TrackInfo) error {
	trackInfo := sqlmodel.TrackInfo{
		UserID:    fmt.Sprint(t.ID),
		Action:    t.Action,
		Timestamp: t.Timestamp,
	}

	err := tr.db.WithContext(ctx).Create(&trackInfo).Error
	if err != nil {
		return err
	}

	return nil
}
