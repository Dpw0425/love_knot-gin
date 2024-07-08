package repo

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"love_knot/internal/app/storage/model"
	ctx "love_knot/internal/pkg/context"
)

type UserRepo struct {
	ctx.Repo[model.User]
}

func NewUsers(db *gorm.DB) *UserRepo {
	return &UserRepo{Repo: ctx.NewRepo[model.User](db)}
}

func (u *UserRepo) IsExist(ctx context.Context, email string) bool {
	if len(email) == 0 {
		return false
	}

	isExist, _ := u.Repo.IsExist(ctx, "email = ?", email)
	return isExist
}

func (u *UserRepo) Create(ctx context.Context, user *model.User) error {
	if err := u.Repo.Create(ctx, user); err != nil {
		return err
	}

	return nil
}

func (u *UserRepo) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	if len(email) == 0 {
		return nil, fmt.Errorf("email empty")
	}

	return u.Repo.FindByWhere(ctx, "email = ?", email)
}
