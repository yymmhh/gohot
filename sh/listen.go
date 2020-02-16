package sh

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
	"github.com/syyongx/php2go"
	"log"
	"strings"
)

//监听管道中个数然后准备重启
func RunListenChan(ch chan int) {
	fmt.Println("开始运行读取管道")
	for {
		select {
		case _ = <-ch:

			if len(ch) > 1 { //消耗掉
				continue
			}

			if ReadConf("listenDir")["ShowLog"] == "true" {
				fmt.Println("管道剩余个数", len(ch))
			}
			//进行重启操作
			Reload()

			php2go.Sleep(1)
		}
	}
}
//启动文件监听
func RunListenDir(loadChan chan int) {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {

		for {
			select {
			case event := <-watcher.Events:

				if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {

					name := event.String()

					//是否输出变动的文件
					if ReadConf("listenDir")["ShowLog"] == "true" {
						fmt.Println(name)
					}

					var index int = -1
					var isHave bool = false
					//检测变动的文件是否在哪些忽略变动的文件中
					for _, v := range ReadConf("ellipsisFile") {
						index = strings.Index(name, v)
						if index == -1 { //没匹配
							isHave = false
						} else { //匹配了
							isHave = true
							break
						}

					}

					if isHave == false {

						loadChan <- php2go.Rand(1, 10) //写入管道

						if ReadConf("listenDir")["ShowLog"] == "true" {
							fmt.Println("写入管道成功,此时管道个数", len(loadChan))
						}

					}

				}

			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	//添加监听的📂
	for _, v := range ReadFile() {
		err = watcher.Add(v)
	}

	if err != nil {
		fmt.Println("出错了,添加的目录太多导致的...!(请去https://github.com/yymmhh/gohot查看具体解决办法)", err)
		color.Set(color.BgRed, color.Bold)
		defer color.Unset()

		log.Fatal(err)
	}

	<-done

}