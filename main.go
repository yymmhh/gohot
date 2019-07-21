package main

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/fsnotify/fsnotify"
	"github.com/syyongx/php2go"
	"hello/services/watch/sh"
	"log"
	"strings"
)

var countChan = make(chan int, 0)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

//读取监听文件
func readFile() []string {
	path := readConf("listenDir")["path"]

	if php2go.Empty(path) { //如果配置文件为空 就获取当前目录
		path = sh.GetCurrentDirectory()
	}

	fmt.Printf("\n %c[1;40;44m%s%c[0m\n\n", 0x1B, "监听目录"+path, 0x1B)

	list, err := sh.GetDirList(path)
	if err != nil {
		fmt.Println(err)

	}

	less := readConf("ellipsisDir")

	listenDirs := []string{}

	var index int = -1
	var isHave bool = false

	for _, dir := range list {

		isHave = false
		for _, v := range less {

			index = strings.Index(dir, v)
			if index == -1 { //没匹配
				isHave = false
			} else { //匹配了
				isHave = true
				break
			}
		}
		if isHave == false {
			//fmt.Println(dir)
			listenDirs = append(listenDirs, dir)
		}

	}

	return listenDirs

}

//读取配置文件
func readConf(section string) map[string]string {
	cfg, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		panic("错误")
	}

	sec, err := cfg.GetSection(section)

	return sec

}
func main() {

	runPHP()

}

func runPHP() {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	} else {

		fmt.Printf("\n %c[1;40;32m%s%c[0m\n\n", 0x1B, ""+
			"  WL_HotSwoole   \n"+
			"     V 1.0       \n"+
			" ====开始运行===== "+
			"", 0x1B)
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {

		for {
			select {
			case event := <-watcher.Events:
				//log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
					//log.Println("modified/**/ file:", event.Name)

					name := event.String()

					//是否输出变动的文件
					if readConf("listenDir")["ShowLog"] == "true" {
						fmt.Println(name)
					}

					var index int = -1
					var isHave bool = false

					for _, v := range readConf("ellipsisFile") {
						index = strings.Index(name, v)
						if index == -1 { //没匹配
							isHave = false
						} else { //匹配了
							isHave = true
							break
						}

					}

					if isHave == false {

						reload()

					}

				}

			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	//添加监听的📂
	for _, v := range readFile() {
		err = watcher.Add(v) //也可以监听文件夹
	}

	if err != nil {
		log.Fatal(err)
	}

	<-done

}

func reload() {

	sh.ReloadSwoole()

}
