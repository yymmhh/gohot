package sh

import (
	"fmt"
	"github.com/Unknwon/goconfig"
)

var time_total = 0

//读取配置文件
func ReadConf(section string) map[string]string {
	file := GetRealPath() + "conf.ini"

	if time_total < 1 {
		fmt.Printf("配置文件: %s\n", file)
	}
	time_total++

	cfg, err := goconfig.LoadConfigFile(file)
	if err != nil {
		panic("读取配置文件错误")
	}

	sec, err := cfg.GetSection(section)

	return sec

}
