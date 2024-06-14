package service

import (
	"github.com/google/wire"
	service "love_knot/internal/app/service/web/v1"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(service.EmailService), "*"),
	wire.Bind(new(service.IEmailService), new(*service.EmailService)),

	wire.Struct(new(service.TemplateService), "*"),
	wire.Bind(new(service.ITemplateService), new(*service.TemplateService)),

	wire.Struct(new(service.UserService), "*"),
	wire.Bind(new(service.IUserService), new(*service.UserService)),
)
