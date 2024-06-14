package service

import (
	"context"
	"love_knot/internal/app/schema"
	"love_knot/internal/app/storage/model"
	"love_knot/internal/app/storage/repo"
	myErr "love_knot/pkg/error"
	"love_knot/utils/encrypt"
	"love_knot/utils/generator"
)

var _ IUserService = (*UserService)(nil)

type IUserService interface {
	Register(ctx context.Context, sur *schema.UserRegister) error
}

type UserService struct {
	UserRepo *repo.UserRepo
}

func (u *UserService) Register(ctx context.Context, sur *schema.UserRegister) error {
	if u.UserRepo.IsExist(ctx, sur.Email) {
		return myErr.BadRequest("", "账号已存在！")
	}

	return u.UserRepo.Create(ctx, &model.User{
		UserID:   generator.IDGenerator(),
		NickName: sur.NickName,
		Password: encrypt.HashPassword(sur.Password),
		Avatar:   sur.Avatar,
		Gender:   sur.Gender,
		Email:    sur.Email,
	})
}
