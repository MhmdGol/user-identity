package sql

import (
	"Identity/internal/model"
	"Identity/internal/repository"
	"Identity/internal/repository/sql/sqlmodel"
	"context"
	"fmt"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/pquerna/otp"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

var _ repository.UserRepo = (*UserRepo)(nil)

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (ur *UserRepo) Create(ctx context.Context, u model.RawUser) error {
	user := sqlmodel.UserInfo{
		UUN:            u.UUN,
		Username:       u.Username,
		HashedPassword: u.Password,
		CreatedAt:      time.Now(),
		Email:          u.Email,
		TotpSecret:     u.TotpSecret.String(),
		Role:           u.Role,
		Status:         u.Status,
	}

	err := ur.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepo) ByID(ctx context.Context, id model.ID) (model.UserInfo, error) {
	var user sqlmodel.UserInfo
	err := ur.db.WithContext(ctx).Where(&sqlmodel.UserInfo{ID: fmt.Sprint(id)}).First(&user).Error
	if err != nil {
		return model.UserInfo{}, err
	}

	sId, _ := snowflake.ParseString(user.ID)
	parsedKey, _ := otp.NewKeyFromURL(user.TotpSecret)

	return model.UserInfo{
		ID:             model.ID(sId),
		UUN:            user.UUN,
		Username:       user.Username,
		HashedPassword: user.HashedPassword,
		Created_at:     user.CreatedAt,
		Email:          user.Email,
		TotpSecret:     parsedKey,
		Role:           user.Role,
		Status:         user.Status,
	}, nil
}

func (ur *UserRepo) ByUsername(ctx context.Context, username string) (model.UserInfo, error) {
	var user sqlmodel.UserInfo
	err := ur.db.WithContext(ctx).Where(&sqlmodel.UserInfo{Username: username}).First(&user).Error
	if err != nil {
		return model.UserInfo{}, err
	}

	sId, _ := snowflake.ParseString(user.ID)
	parsedKey, _ := otp.NewKeyFromURL(user.TotpSecret)

	return model.UserInfo{
		ID:             model.ID(sId),
		UUN:            user.UUN,
		Username:       user.Username,
		HashedPassword: user.HashedPassword,
		Created_at:     user.CreatedAt,
		Email:          user.Email,
		TotpSecret:     parsedKey,
		Role:           user.Role,
		Status:         user.Status,
	}, nil
}

func (ur *UserRepo) UpdateByID(ctx context.Context, u model.UserInfo) error {
	var user sqlmodel.UserInfo
	err := ur.db.WithContext(ctx).Where(&sqlmodel.UserInfo{ID: fmt.Sprint(u.ID)}).First(&user).Error
	if err != nil {
		return err
	}

	if u.UUN != "" {
		user.UUN = u.UUN
	}
	if u.Username != "" {
		user.Username = u.Username
	}
	if u.HashedPassword != "" {
		user.HashedPassword = u.HashedPassword
	}
	if u.Email != "" {
		user.Email = u.Email
	}
	if u.TotpSecret != nil {
		user.TotpSecret = u.TotpSecret.String()
	}
	if u.Role != "" {
		user.Role = u.Role
	}
	if u.Status != "" {
		user.Status = u.Status
	}

	return ur.db.WithContext(ctx).Save(&user).Error
}

// logical flaw: what if service call this function to update username, eg. < set mhmd to erfan where username = mhmd >
func (ur *UserRepo) UpdateByUsername(ctx context.Context, u model.UserInfo) error {
	var user sqlmodel.UserInfo
	err := ur.db.WithContext(ctx).Where(&sqlmodel.UserInfo{Username: u.Username}).First(&user).Error
	if err != nil {
		return err
	}

	if fmt.Sprint(u.ID) != "" {
		user.ID = fmt.Sprint(u.ID)
	}
	if u.UUN != "" {
		user.UUN = u.UUN
	}
	if u.HashedPassword != "" {
		user.HashedPassword = u.HashedPassword
	}
	if u.Email != "" {
		user.Email = u.Email
	}
	if u.TotpSecret != nil {
		user.TotpSecret = u.TotpSecret.String()
	}
	if u.Role != "" {
		user.Role = u.Role
	}
	if u.Status != "" {
		user.Status = u.Status
	}

	return ur.db.WithContext(ctx).Save(&user).Error
}
