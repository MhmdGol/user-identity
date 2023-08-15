package service

import (
	"Identity/internal/model"
	"Identity/internal/repository"
	"Identity/internal/service"
	"Identity/pkg/bcrypthash"
	"context"

	"github.com/bwmarrin/snowflake"
	"github.com/casbin/casbin"
	"github.com/pquerna/otp/totp"
)

type UserService struct {
	userRepo repository.UserRepo
	enforcer *casbin.Enforcer
	sf       snowflake.Node
}

var _ service.UserService = (*UserService)(nil)

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) Create(ctx context.Context, ru model.RawUser) error {
	hpass, _ := bcrypthash.HashPassword(ru.Password)

	userSecret, _ := totp.Generate(totp.GenerateOpts{
		Issuer:      "IdentityServer",
		AccountName: ru.Username,
	})

	return us.userRepo.Create(ctx, model.RawUser{
		ID:         model.ID(us.sf.Generate()),
		UUN:        ru.UUN,
		Username:   ru.Username,
		Password:   hpass,
		Email:      ru.Email,
		TotpSecret: userSecret,
		Role:       ru.Role,
		Status:     ru.Status,
	})
}

func (us *UserService) ReadByID(ctx context.Context, id model.ID) (model.UserInfo, error) {
	user, _ := us.userRepo.ByID(ctx, id)

	return model.UserInfo{
		ID:             user.ID,
		UUN:            user.UUN,
		Username:       user.Username,
		HashedPassword: user.HashedPassword,
		Created_at:     user.Created_at,
		Email:          user.Email,
		TotpSecret:     user.TotpSecret,
		Role:           user.Role,
		Status:         user.Status,
	}, nil
}

func (us *UserService) ChangePermission(ctx context.Context, username string, role string) error {
	us.enforcer.RemoveFilteredGroupingPolicy(0, username)
	us.enforcer.AddGroupingPolicy(username, role)

	return us.userRepo.UpdateByUsername(ctx, model.UserInfo{
		Username: username,
		Role:     role,
	})
}

func (us *UserService) UpdateByID(ctx context.Context, u model.UserInfo) error {
	return us.userRepo.UpdateByID(ctx, u)
}
