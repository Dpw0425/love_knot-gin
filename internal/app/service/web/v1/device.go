package service

import (
	"context"
	"love_knot/internal/app/storage/model"
	"love_knot/internal/app/storage/repo"
	"time"
)

var _ IDeviceService = (*DeviceService)(nil)

type IDeviceService interface {
	SetUserCommonDevice(ctx context.Context, uid int64, ip, address, agent string) error
	IsCommonDevice(ctx context.Context, uid int64, ip string) bool
}

type DeviceService struct {
	DeviceRepo *repo.DeviceRepo
}

// SetUserCommonDevice 绑定常用登录设备
func (d *DeviceService) SetUserCommonDevice(ctx context.Context, uid int64, ip, address, agent string) error {
	return d.DeviceRepo.Create(ctx, &model.Device{
		UserID:    uid,
		Address:   address,
		IP:        ip,
		Agent:     agent,
		LoginTime: time.Now(),
	})
}

// IsCommonDevice 判断是否为常用登录设备
func (d *DeviceService) IsCommonDevice(ctx context.Context, uid int64, ip string) bool {
	result, err := d.DeviceRepo.GetLoginDevice(ctx, uid)
	if err != nil {
		return false
	}

	if result.IP != ip {
		return false
	}

	return true
}
