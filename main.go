package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/syyongx/php2go"
	"log"
	"os"
	"strings"
	"wl_GoHot/sh"
)


//读取监听文件
func readFile() []string {
	path := sh.ReadConf("listenDir")["path"]

	if php2go.Empty(path) { //如果配置文件为空 就获取当前目录
		path,_= os.Getwd()
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


func main() {
	//协程不影响后头运行
	go func() {
		sh.StartSwoole()
	}()

	runPHP()

}



func runPHP() {


	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	} else {

		fmt.Printf("\n %c[1;40;32m%s%c[0m\n\n", 0x1B, ""+
			"  Wl_GoHot   \n"+
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
