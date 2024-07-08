package model

import (
	"gorm.io/gorm"
	"time"
)

type Device struct {
	gorm.Model
	UserID    int64     `gorm:"column:user_id;NOT NULL" json:"user_id"` // 关联用户的 uid
	Address   string    `gorm:"column:address;NOT NULL" json:"address"` // 位置信息
	IP        string    `gorm:"column:ip;NOT NULL" json:"ip"`           // ip 地址
	Agent     string    `gorm:"column:agent;NOT NULL" json:"agent"`     // 设备信息
	LoginTime time.Time `gorm:"column:login_time" json:"login_time"`    // 上次登录时间
	Status    int       `gorm:"column:status;default:0;NOT NULL"`       // 设备状态[0:正常状态;1:禁用状态]
}

func (Device) TableName() string {
	return "device"
}
