package service

import (
	"github.com/gin-gonic/gin"
	"love_knot/internal/app/constants"
	"love_knot/internal/app/schema"
	"love_knot/internal/app/storage/repo"
)

var _ IFriendService = (*FriendService)(nil)

type IFriendService interface {
	List(ctx *gin.Context, UserID int) ([]*schema.FriendListItem, error)
}

type FriendService struct {
	FriendRepo *repo.FriendRepo
}

func (f *FriendService) List(ctx *gin.Context, UserID int) ([]*schema.FriendListItem, error) {
	list := make([]*schema.FriendListItem, 0)

	tx := f.FriendRepo.Model(ctx)
	tx.Select([]string{
		"users.id",
		"users.nick_name",
		"users.avatar",
		"users.gender",
	})
	tx.Joins("inner join `users` ON `users.id` = friends.friend_id")
	tx.Where("friends.user_id = ? and friends.status = ?", UserID, constants.NormalFriendStatus)

	if err := tx.Scan(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}
