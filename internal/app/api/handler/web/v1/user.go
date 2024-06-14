package v1

import (
	"love_knot/internal/app/constants"
	"love_knot/internal/app/schema"
	service "love_knot/internal/app/service/web/v1"
	ctx "love_knot/internal/pkg/context"
	myErr "love_knot/pkg/error"
	"love_knot/pkg/response"
	web "love_knot/schema/genproto/web/v1/user"
)

type User struct {
	UserService  service.IUserService
	EmailService service.IEmailService
}

func (u *User) Register(ctx *ctx.Context) error {
	params := &web.UserRegisterRequest{}
	if err := ctx.Context.ShouldBind(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	if !u.EmailService.Verify(ctx.Ctx(), constants.EmailRegisterChannel, params.Email, params.VerifyCode) {
		return myErr.BadRequest("wrong_verification_code", "验证码错误！")
	}

	if err := u.UserService.Register(ctx.Ctx(), &schema.UserRegister{
		NickName: params.NickName,
		Password: params.Password,
		Avatar:   params.Avatar,
		Gender:   int(params.Gender),
		Email:    params.Email,
	}); err != nil {
		return err
	}

	u.EmailService.Delete(ctx.Ctx(), constants.EmailRegisterChannel, params.Email)

	response.NorResponse(ctx.Context, &web.UserRegisterResponse{}, "注册成功！")
	return nil
}
