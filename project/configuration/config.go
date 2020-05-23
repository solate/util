package configuration

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
)

const (
	//配置文件默认位置
	Path = "./app/config/config.toml"
)

var (
	configPath string // 配置文件路径
)

//加载配置
func LoadConfig(s interface{}) (err error) {
	//获取参数
	flag.StringVar(&configPath, "config", Path, "set config file path!")
	flag.Parse()

	// 解析参数
	if _, err = toml.DecodeFile(configPath, s); err != nil {
		log.Println("failed to load config.toml" + err.Error())
		return
	}
	return
}
