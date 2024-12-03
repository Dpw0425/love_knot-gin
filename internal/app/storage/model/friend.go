package model

import "gorm.io/gorm"

type Friend struct {
	gorm.Model
	UserID   int64 `gorm:"column:user_id_1;NOT NULL" json:"user_id"`       // 第一个用户 id
	FriendID int64 `gorm:"column:user_id_2;NOT;NULL" json:"friend_id"`     // 第二个用户 id
	Status   int   `gorm:"column:status;default:0;NOT NULL" json:"status"` // 0: 未认证, 1: 好友关系, 2: 特别关心
}

func (Friend) TableName() string { return "friends" }
