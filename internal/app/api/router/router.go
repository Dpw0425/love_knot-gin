package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/sjson"
	"gopkg.in/natefinch/lumberjack.v2"
	"love_knot/internal/app/config"
	"love_knot/internal/app/middleware"
	myErr "love_knot/pkg/error"
	"love_knot/pkg/response"
	"net/http"
)

func NewRouter(conf *config.Config) *gin.Engine {
	router := gin.New()

	accessFilterRule := middleware.NewAccessFilterRule()
	accessFilterRule.AddRule("/api/v1/user/login", func(req *middleware.RequestInfo) {
		req.RequestBody, _ = sjson.Set(req.RequestBody, `password`, "过滤敏感字段")
	})

	router.Use(middleware.AccessLog(&lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/logs/router_log/access.log", conf.Log.Path), // 日志文件的位置
		MaxSize:    100,                                                         // 文件最大尺寸（以MB为单位）
		MaxAge:     7,                                                           // 保留旧文件的最大天数
		MaxBackups: 3,                                                           // 保留的最大旧文件数量
		LocalTime:  true,                                                        // 使用本地时间创建时间戳
		Compress:   true,                                                        // 是否压缩/归档旧文件
	}, accessFilterRule))

	router.Use(gin.RecoveryWithWriter(gin.DefaultWriter, func(c *gin.Context, err any) {
		response.ErrResponse(c, myErr.InternalServerError("", "系统错误，请重试!!!"))
	}))

	router.GET("/", func(c *gin.Context) {
		response.NorResponse(c, http.StatusOK, gin.H{}, "hello world!")
	})

	return router
}
