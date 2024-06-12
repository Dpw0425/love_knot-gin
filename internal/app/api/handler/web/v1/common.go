package v1

import (
	service "love_knot/internal/app/service/web/v1"
	ctx "love_knot/internal/pkg/context"
	myErr "love_knot/pkg/error"
	"love_knot/pkg/response"
	"love_knot/schema/genproto/web/v1/common"
)

type Common struct {
	EmailService service.IEmailService
}

func (c *Common) SendEmailCode(ctx *ctx.Context) error {
	params := &web.SendEmailCodeRequest{}
	if err := ctx.Context.ShouldBind(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	err := c.EmailService.Send(ctx.Ctx(), params.Channel, params.Email)
	if err != nil {
		return err
	}

	response.NorResponse(ctx.Context, &web.SendEmailCodeResponse{}, "发送成功！")
	return nil
}
