package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"love_knot/pkg/logger"
	"love_knot/utils/encrypt"
	"love_knot/utils/random"
	"os"
	"time"
)

type Config struct {
	sid    string  // 服务运行 ID
	App    *App    `json_utils:"app" yaml:"app" mapstructure:"app"`
	Server *Server `json_utils:"server" yaml:"server" mapstructure:"server"`
	Log    *Log    `json_utils:"log" yaml:"log" mapstructure:"log"`
	Mysql  *Mysql  `json_utils:"mysql" yaml:"mysql" mapstructure:"mysql"`
}

func Load(filename string) *Config {
	content, err := os.ReadFile(filename)
	if err != nil {
		logger.Panicf("Read Config Error: %v1", err)
		panic(err)
	}

	var conf Config
	if yaml.Unmarshal(content, &conf) != nil {
		logger.Panicf("Parsing Config Failed: %v1", err)
		panic(fmt.Sprintf("解析 config.yaml 失败: %v1", err))
	}

	conf.sid = encrypt.Md5(fmt.Sprintf("%d%s", time.Now().UnixNano(), random.Random(6)))

	return &conf
}

// ServerID 服务运行 ID
func (c *Config) ServerID() string {
	return c.sid
}

// RunMode app 运行模式
func (c *Config) RunMode() string {
	return c.App.RunMode
}

// LogPath 日志存储目录
func (c *Config) LogPath() string {
	return c.Log.Path
}
