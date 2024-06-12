package job

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
	"love_knot/internal/app/storage/model"
	"love_knot/internal/config"
	"love_knot/pkg/logger"
	"os"
	"strings"
)

type SQLProvider struct {
	Config *config.Config
	DB     *gorm.DB
}

func Run(ctx *cli.Context, app *SQLProvider) error {
	fmt.Println("数据库初始化中...")

	content, err := os.ReadFile("./docs/sql/love_knot.sql")
	if err != nil {
		fmt.Println("数据库导入失败: ", err)
		logger.Errorf("Databases Import Failed: %s", err)
		fmt.Println("正在自动生成数据库表...")
		app.DB.AutoMigrate(&model.User{})
		fmt.Println("数据库表初始化成功！")
		logger.Info("Databases Init Successful!")
	}

	for _, sql := range strings.Split(string(content), ";;") {
		if len(sql) > 0 {
			_ = app.DB.Exec(strings.TrimSpace(sql)).Error
			fmt.Println("数据库导入成功！")
			logger.Info("Databases Import Successful!")
		}
	}

	return nil
}
