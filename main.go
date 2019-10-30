package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
	"github.com/syyongx/php2go"
	"gohot/sh"
	"log"
	"os"
	"strings"
)

//读取监听文件
func readFile() []string {
	path := sh.ReadConf("listenDir")["path"]

	if php2go.Empty(path) { //如果配置文件为空 就获取当前目录
		path, _ = os.Getwd()
	}

	fmt.Printf("\n %c[1;40;44m%s%c[0m\n\n", 0x1B, "监听目录"+path, 0x1B)

	list, err := sh.GetDirList(path)
	if err != nil {
		fmt.Println(err)

	}

	less := sh.ReadConf("ellipsisDir")

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

//读取管道中个数然后准备重启
func RunChan(ch chan int) {

	for {
		select {
		case _ = <-ch:

			if len(ch) > 1 { //消耗掉
				continue
			}

			if sh.ReadConf("listenDir")["ShowLog"] == "true" {
				fmt.Println("管道剩余个数", len(ch))
			}

			reload()
			php2go.Sleep(1)
		}
	}
}

func main() {
	LoadCountChan := make(chan int, 100)

	//协程不影响后头运行
	go func() {
		sh.StartSwoole()
	}()

	//启动管道监听
	go RunChan(LoadCountChan)

	//启动文件监听
	runPHP(LoadCountChan)

}

func runPHP(loadChan chan int) {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	} else {

		fmt.Printf("\n %c[1;40;32m%s%c[0m\n\n", 0x1B, ""+
			"  Wl_GoHot   \n"+
			"     V 1.1       \n"+
			"  Author:yymmhh"+
			" ====开始运行===== "+
			"", 0x1B)
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {

		for {
			select {
			case event := <-watcher.Events:

				if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
					//log.Println("modified/**/ file:", event.Name)

					name := event.String()

					//是否输出变动的文件
					if sh.ReadConf("listenDir")["ShowLog"] == "true" {
						fmt.Println(name)
					}

					var index int = -1
					var isHave bool = false

					for _, v := range sh.ReadConf("ellipsisFile") {
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

						if sh.ReadConf("listenDir")["ShowLog"] == "true" {
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
	for _, v := range readFile() {
		err = watcher.Add(v) //也可以监听文件夹
	}

	if err != nil {
		fmt.Println("出错了,添加的目录太多导致的...!", err)
		color.Set(color.BgRed, color.Bold)
		defer color.Unset()

		log.Fatal(err)
	}

	<-done

}

func reload() {

	sh.ReloadSwoole()

}
