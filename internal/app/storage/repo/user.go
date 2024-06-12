package repo

import (
	"context"
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
