package model

import "gorm.io/gorm"

type Robot struct {
	gorm.Model
	RobotName string `gorm:"column:robot_name;NOT NULL" json:"robot_name"` // 机器人名称
	Describe  string `gorm:"column:describe" json:"describe"`              // 描述信息
	Logo      string `gorm:"column:logo" json:"logo"`                      // 机器人图标
	Status    int    `gorm:"status" json:"status"`                         // 机器人状态
}

func (Robot) TableName() string {
	return "robot"
}
