package router

import (
	"github.com/gin-gonic/gin"
	"love_knot/internal/app/api/handler/web"
	"love_knot/internal/app/middleware"
	ctx "love_knot/internal/pkg/context"
)

func RegisterWebRouter(secret string, router *gin.Engine, handler *web.Handler) {
	authorizer := middleware.Auth(secret, "api")

	v1 := router.Group("/api/v1")
	{
		common := v1.Group("/common")
		{
			common.POST("/send_email_code", ctx.HandlerFunc(handler.V1.Common.SendEmailCode)) // 邮箱验证码
		}

		user := v1.Group("/user")
		{
			user.POST("/register", ctx.HandlerFunc(handler.V1.User.Register)) // 用户注册
			user.POST("/login", ctx.HandlerFunc(handler.V1.User.Login))       // 用户登录
		}

		friend := v1.Group("/friend").Use(authorizer)
		{
			friend.GET("/list", ctx.HandlerFunc(handler.V1.Friend.List)) // 好友列表
		}
	}
}
