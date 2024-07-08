package service

import (
	"github.com/goccy/go-json"
	"love_knot/internal/app/schema"
	"love_knot/internal/app/storage/repo"
	"love_knot/internal/config"
	"love_knot/internal/pkg/client"
	myErr "love_knot/pkg/error"
	"love_knot/utils/slice_utils"
	"net/url"
	"strings"
)

var _ IIpAddressService = (*IpAddressService)(nil)

type IIpAddressService interface {
	GetAddress(ip string) (string, error)
}

type IpAddressService struct {
	*repo.Source
	Config *config.Config
	Client *client.RequestClient
}

func (i *IpAddressService) GetAddress(ip string) (string, error) {
	if val, err := i.getCache(ip); err == nil {
		return val, nil
	}

	params := &url.Values{}
	params.Add("ip", ip)
	params.Add("key", i.Config.App.GaoDeKey)

	resp, err := i.Client.Get("https://restapi.amap.com/v3/ip?", params)
	if err != nil {
		return "", nil
	}

	data := &schema.IPAddressResponse{}
	if err := json.Unmarshal(resp, data); err != nil {
		return "", err
	}

	if data.Adcode == "" {
		return "", myErr.BadRequest("ip_address_service_error", "位置获取失败！")
	}

	arr := []string{data.Country, data.Province, data.City, data.ISP}
	result := strings.Join(slice_utils.Unique(arr), " ")
	result = strings.TrimSpace(result)

	_ = i.setCache(ip, result)

	return result, nil
}

func (i *IpAddressService) getCache(ip string) (string, error) {
	return i.Source.Redis().HGet("rds:hash:ip-address", ip).Result()
}

func (i *IpAddressService) setCache(ip string, value string) error {
	return i.Source.Redis().HSet("rds:hash:ip-address", ip, value).Err()
}
