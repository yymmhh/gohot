package sh

import (
	"github.com/Unknwon/goconfig"
)


//读取配置文件
func ReadConf(section string) map[string]string {

	sec, err := Config.GetSection(section)
	if err!=nil {
		panic(err)
	}

	return sec
}
//加载配置文件
func LoadConfig()  {

	var file string
	if wlDebug==true {
		file=configPath
	}else{
		file= GetRealPath() + "conf.ini"
	}

	cfg, err := goconfig.LoadConfigFile(file)
	if err != nil {
		panic("读取配置文件错误")
	}
	Config=cfg
}