package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID   int64  `gorm:"column:user_id;unique" json:"id"`                // id
	NickName string `gorm:"column:nick_name;NOT NULL" json:"nick_name"`     // 昵称
	Password string `gorm:"column:password;NOT NULL" json:"-"`              // 密码
	Avatar   string `gorm:"column:avatar;NOT NULL" json:"avatar"`           // 头像
	Gender   int    `gorm:"column:gender;default:0;NOT NULL" json:"gender"` // 0: 保密, 1: 男, 2: 女
	Email    string `gorm:"column:email;NOT NULL" json:"email"`             // 邮箱
}

func (User) TableName() string {
	return "users"
}
