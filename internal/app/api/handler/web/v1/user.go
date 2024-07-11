package v1

import (
	"love_knot/internal/app/constants"
	"love_knot/internal/app/schema"
	service "love_knot/internal/app/service/web/v1"
	"love_knot/internal/app/storage/model"
	"love_knot/internal/config"
	ctx "love_knot/internal/pkg/context"
	myErr "love_knot/pkg/error"
	"love_knot/pkg/jwt"
	"love_knot/pkg/response"
	web "love_knot/schema/genproto/web/v1/user"
	"strconv"
	"sync"
	"time"
)

type User struct {
	Config           *config.Config
	UserService      service.IUserService
	EmailService     service.IEmailService
	IpAddressService service.IIpAddressService
	DeviceService    service.IDeviceService
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
		NickName: params.Nickname,
		Password: params.Password,
		Avatar:   params.Avatar,
		Gender:   int(params.Gender),
		Email:    params.Email,
	}); err != nil {
		return myErr.BadRequest("", "用户注册失败: %v", err)
	}

	u.EmailService.Delete(ctx.Ctx(), constants.EmailRegisterChannel, params.Email)

	response.NorResponse(ctx.Context, &web.UserRegisterResponse{}, "注册成功！")
	return nil
}

func (u *User) Login(ctx *ctx.Context) error {
	var wg sync.WaitGroup

	params := &web.UserLoginRequest{}
	if err := ctx.Context.ShouldBind(&params); err != nil {
		return myErr.BadRequest("wrong_parameters", "请求参数错误！")
	}

	var ch = make(chan string, 3)
	defer close(ch)

	wg.Add(1)
	go func() {
		ip := ctx.Context.ClientIP()
		ch <- ip

		address, _ := u.IpAddressService.GetAddress(ip)
		ch <- address

		agent := ctx.Context.Request.Header.Get("User-Agent")
		ch <- agent

		wg.Done()
	}()

	ip := <-ch
	address := <-ch
	agent := <-ch

	var user *model.User
	if params.Password == "" {
		if !u.EmailService.Verify(ctx.Ctx(), constants.EmailLoginChannel, params.Email, params.VerifyCode) {
			return myErr.BadRequest("wrong_verification_code", "验证码错误！")
		}

		u.EmailService.Delete(ctx.Ctx(), constants.EmailLoginChannel, params.Email)
		result, err := u.UserService.GetUserByEmail(ctx.Ctx(), params.Email)
		if err != nil {
			return err
		}
		user = result
	} else {
		result, err := u.UserService.LoginByPassword(ctx.Ctx(), &schema.UserLogin{
			Email:    params.Email,
			Password: params.Password,
		})
		if err != nil {
			return err
		}
		user = result
	}

	// TODO: VERIFY IF THE DEVICE IS AUTHORIZED
	if !u.DeviceService.IsCommonDevice(ctx.Ctx(), user.UserID, ip) && params.VerifyCode == "" {
		return myErr.Forbidden("", "非常用登录设备，请使用邮箱验证登录！")
	}

	_ = u.DeviceService.SetUserCommonDevice(ctx.Ctx(), user.UserID, ip, address, agent)

	// TODO: PUSH USER LOGIN MESSAGE

	wg.Wait()
	response.NorResponse(ctx.Context, &web.UserLoginResponse{
		Type:        "Bearer",
		AccessToken: u.token(user.UserID),
		ExpiresIn:   strconv.FormatInt(u.Config.Jwt.ExpiresTime, 10),
	}, "登录成功！")
	return nil
}

func (u *User) token(uid int64) string {
	expiresAt := time.Now().Add(time.Second * time.Duration(u.Config.Jwt.ExpiresTime))

	token, _ := jwt.GenerateToken("api", u.Config.Jwt.Secret, &jwt.Options{
		ExpiresAt: jwt.NewNumericData(expiresAt),
		ID:        strconv.FormatInt(uid, 10),
		Issuer:    "love_knot.web",
		IssuedAt:  jwt.NewNumericData(time.Now()),
	})

	return token
}
