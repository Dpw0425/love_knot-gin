package repo

import (
	"context"
	"gorm.io/gorm"
	"love_knot/internal/app/storage/model"
	ctx "love_knot/internal/pkg/context"
)

type DeviceRepo struct {
	ctx.Repo[model.Device]
}

func NewDevice(db *gorm.DB) *DeviceRepo {
	return &DeviceRepo{ctx.NewRepo[model.Device](db)}
}

func (d *DeviceRepo) Create(ctx context.Context, device *model.Device) error {
	if err := d.Repo.Create(ctx, device); err != nil {
		return err
	}

	return nil
}

// GetLoginDevice 获取登录设备信息
func (d *DeviceRepo) GetLoginDevice(ctx context.Context) (*model.Device, error) {
	// TODO: COMPLETE THIS QUERY STATEMENT
	return d.Repo.FindByWhere(ctx, "user_id = ? and status = ?")
}
