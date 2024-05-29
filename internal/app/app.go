package app

import (
	"github.com/urfave/cli/v2"
	"love_knot/internal/app/config"
	"os"
)

type App struct {
	app *cli.App
}

type Action func(ctx *cli.Context, conf *config.Config) error

type Command struct {
	Name        string
	Usage       string
	Flags       []cli.Flag
	Action      Action
	SubCommands []Command
}

func NewApp() *App {
	return &App{
		app: &cli.App{
			Name:    "LoveKnot",
			Usage:   "",
			Version: "v0.0.240422",
		},
	}
}

func (c *App) Register(cm Command) {
	c.app.Commands = append(c.app.Commands, c.command(cm))
}

func (c *App) command(cm Command) *cli.Command {
	cd := &cli.Command{
		Name:  cm.Name,
		Usage: cm.Usage,
		Flags: make([]cli.Flag, 0),
	}

	// 批处理
	if len(cm.SubCommands) > 0 {
		for _, v := range cm.SubCommands {
			cd.Subcommands = append(cd.Subcommands, c.command(v))
		}
	} else {
		if cm.Flags != nil && len(cm.Flags) > 0 {
			cd.Flags = append(cd.Flags, cm.Flags...)
		}

		var isConfig bool

		for _, flag := range cd.Flags {
			if flag.Names()[0] == "config" {
				isConfig = true
				break
			}
		}

		if !isConfig {
			cd.Flags = append(cd.Flags, &cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Value:       "./config.yaml",
				Usage:       "配置文件路径",
				DefaultText: "./config.yaml",
			})
		}

		if cm.Action != nil {
			cd.Action = func(ctx *cli.Context) error {
				return cm.Action(ctx, config.Load(ctx.String("config")))
			}
		}
	}

	return cd
}

func (c *App) Run() {
	err := c.app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
