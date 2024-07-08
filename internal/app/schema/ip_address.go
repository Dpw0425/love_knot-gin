package schema

type IPAddressResponse struct {
	Status   string // 查询结果
	Info     string // 返回状态说明
	Infocode string // 状态码
	Country  string // 国家
	Province string // 省份名称
	City     string // 城市名称
	Adcode   string // 区域编码
	District string // 地区名称
	ISP      string // 运营商
}
