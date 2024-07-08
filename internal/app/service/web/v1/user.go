package service

import (
	"context"
	"gorm.io/gorm"
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
	LoginByPassword(ctx context.Context, sul *schema.UserLogin) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
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

func (u *UserService) LoginByPassword(ctx context.Context, sul *schema.UserLogin) (*model.User, error) {
	user, err := u.UserRepo.FindByEmail(ctx, sul.Email)
	if err != nil {
		if myErr.Equal(err, gorm.ErrRecordNotFound) {
			return nil, myErr.BadRequest("", "账号不存在！")
		}

		return nil, myErr.BadRequest("", err.Error())
	}

	if !encrypt.VerifyPassword(user.Password, sul.Password) {
		return nil, myErr.BadRequest("", "密码错误！")
	}

	return user, nil
}

func (u *UserService) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := u.UserRepo.FindByEmail(ctx, email)
	if err != nil {
		if myErr.Equal(err, gorm.ErrRecordNotFound) {
			return nil, myErr.BadRequest("", "账号不存在！")
		}

		return nil, myErr.BadRequest("", err.Error())
	}

	return user, nil
}
