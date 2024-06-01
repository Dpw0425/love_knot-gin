package v1

import (
	service "love_knot/internal/app/service/web/v1"
	ctx "love_knot/internal/pkg/context"
	myErr "love_knot/pkg/error"
	"love_knot/pkg/response"
	"love_knot/schema/genproto/web/v1/common"
)

type Common struct {
	CommonService service.IEmailService
}

func (c *Common) SendEmailCode(ctx *ctx.Context) {
	params := &web.SendEmailCodeRequest{}
	if err := ctx.Context.ShouldBind(&params); err != nil {
		response.ErrResponse(ctx.Context, myErr.BadRequest("", "请求参数错误！"))
		return
	}

	c.CommonService.Verify()
}
