package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"love_knot/internal/app"
	"love_knot/internal/app/api"
	"love_knot/internal/config"
	"love_knot/internal/job"
	"love_knot/pkg/logger"
)

func NewHttpCommand() app.Command {
	return app.Command{
		Name:  "http",
		Usage: "Http Command - http API 接口服务",
		Action: func(ctx *cli.Context, conf *config.Config) error {
			logger.InitLogger(fmt.Sprintf("%s/logs/app_log", conf.Log.Path), "http")
			return api.RunHttpServer(ctx, NewHttpInjector(conf))
		},
	}
}

func NewSQLCommand() app.Command {
	return app.Command{
		Name:  "sql",
		Usage: "sql command - 数据库初始化",
		Action: func(ctx *cli.Context, conf *config.Config) error {
			logger.InitLogger(fmt.Sprintf("%s/logs/app_log", conf.Log.Path), "sql")
			return job.Run(ctx, NewSQLInjector(conf))
		},
	}
}

//func NewOtherCommand() app.Command {
//	return app.Command{
//		Name:  "other",
//		Usage: "other command - 其它临时命令",
//		SubCommands: []app.Command{
//			{
//				Name:  "test",
//				Usage: "test command",
//				Action: func(ctx *cli.Context, conf *config.Config) error {
//					logger.InitLogger(fmt.Sprintf("%s/logs/app_log", conf.Log.Path), "test")
//					return
//				},
//			},
//		},
//	}
//}

func main() {
	newApp := app.NewApp()
	newApp.Register(NewHttpCommand())
	newApp.Register(NewSQLCommand())
	newApp.Run()
}
