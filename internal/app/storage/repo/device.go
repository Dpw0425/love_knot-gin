package repo

import (
	"context"
	"gorm.io/gorm"
	"love_knot/internal/app/constants"
	"love_knot/internal/app/storage/model"
	ctx "love_knot/internal/pkg/context"
)

type DeviceRepo struct {
	ctx.Repo[model.Device]
}

func NewDevice(db *gorm.DB) *DeviceRepo {
	return &DeviceRepo{ctx.NewRepo[model.Device](db)}
}

// Create 创建常用登录设备
func (d *DeviceRepo) Create(ctx context.Context, device *model.Device) error {
	if err := d.Repo.Create(ctx, device); err != nil {
		return err
	}

	return nil
}

// GetLoginDevice 获取登录设备信息
func (d *DeviceRepo) GetLoginDevice(ctx context.Context, uid int64) (*model.Device, error) {
	return d.Repo.FindByWhere(ctx, "user_id = ? and status = ?", uid, constants.NormalStatus)
}
