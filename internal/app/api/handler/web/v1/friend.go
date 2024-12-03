package v1

import (
	service "love_knot/internal/app/service/web/v1"
	"love_knot/internal/config"
	ctx "love_knot/internal/pkg/context"
	myErr "love_knot/pkg/error"
	"love_knot/pkg/response"
	web "love_knot/schema/genproto/web/v1/friend"
)

type Friend struct {
	Config        *config.Config
	FriendService service.IFriendService
}

func (f *Friend) List(ctx *ctx.Context) error {
	list, err := f.FriendService.List(ctx.Context, ctx.UserID())
	if err != nil {
		return myErr.BadRequest("", "查询失败！")
	}

	items := make([]*web.FriendListResponse_Item, 0, len(list))
	for _, item := range list {
		items = append(items, &web.FriendListResponse_Item{
			UserId:   item.UserID,
			NickName: item.NickName,
			Avatar:   item.Avatar,
			Gender:   int32(item.Gender),
		})
	}

	response.NorResponse(ctx.Context, &web.FriendListResponse{Items: items}, "查询成功！")
	return nil
}
