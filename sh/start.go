package sh

import (
	"github.com/Unknwon/goconfig"
)

/**
初始化启动
*/

var Config *goconfig.ConfigFile

var wlDebug bool = false

var configPath string = "/Users/mac/www/go/src/gohot/conf.ini"

func init() {
	//展示一些程序的信息
	ShowAuthor()
	//加载配置文件到内存中
	LoadConfig()

}

func Run() {

	LoadCountChan := make(chan int, 100)
	//初始化启动程序
	go func() {
		StartSwoole()
	}()

	//启动定时任务
	go StartCrobTable()

	//启动监听文件夹
	go RunListenDir(LoadCountChan)
	//启动监听管道(并且保留着常驻进程)
	RunListenChan(LoadCountChan)

}
