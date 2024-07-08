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
	//IsCommonDevice() bool
}

type DeviceService struct {
	DeviceRepo *repo.DeviceRepo
}

func (d *DeviceService) SetUserCommonDevice(ctx context.Context, uid int64, ip, address, agent string) error {
	return d.DeviceRepo.Create(ctx, &model.Device{
		UserID:    uid,
		Address:   address,
		IP:        ip,
		Agent:     agent,
		LoginTime: time.Now(),
	})
}

//func (d *DeviceService) IsCommonDevice() bool {
//
//}
