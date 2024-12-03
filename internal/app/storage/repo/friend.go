package repo

import (
	"gorm.io/gorm"
	"love_knot/internal/app/storage/model"
	ctx "love_knot/internal/pkg/context"
)

type FriendRepo struct {
	ctx.Repo[model.Friend]
}

func NewFriend(db *gorm.DB) *FriendRepo { return &FriendRepo{ctx.NewRepo[model.Friend](db)} }
