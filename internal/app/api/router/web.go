package router

import (
	"github.com/gin-gonic/gin"
	"love_knot/internal/app/api/handler/web"
	ctx "love_knot/internal/pkg/context"
)

func RegisterWebRouter(secret string, router *gin.Engine, handler *web.Handler) {
	// authorize := middleware.Auth(secret, "api")

	v1 := router.Group("/api/v1")
	{
		common := v1.Group("/common")
		{
			common.POST("/send_email_code", ctx.HandlerFunc(handler.V1.Common.SendEmailCode)) // 邮箱验证码
		}

		user := v1.Group("/user")
		{
			user.POST("/register", ctx.HandlerFunc(handler.V1.User.Register)) // 用户注册
		}
	}
}
