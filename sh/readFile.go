package sh

import (
	"fmt"
	"github.com/syyongx/php2go"
	"os"
	"strconv"
)

/**
读取要监听的文件
 */
var ListenDirs []string

//读取监听文件
func ReadFile() [] string{
	path := ReadConf("listenDir")["path"]

	if php2go.Empty(path) { //如果配置文件为空 就获取当前目录
		path, _ = os.Getwd()
	}

	fmt.Printf("\n %c[1;40;44m%s%c[0m\n\n", 0x1B, "监听目录"+path, 0x1B)

	//再获取配置文件中监听哪些目录
	DirPath := ReadConf("listenFileDir")


	//循环配置文件中监听的这些目录,然后再去递归获取他们下面的文件夹
	for _, v := range DirPath{
		fmt.Println(path+"/"+v)

		list, err := GetDirList(path+"/"+v)
		if err != nil {
			fmt.Println("获取目录出错",err)
		}
		ListenDirs=append(ListenDirs,list...)
	}
	ListenDirsLength:=len(ListenDirs)
	fmt.Println("一共监听了"+strconv.Itoa(ListenDirsLength)+"个文件夹")

	return  ListenDirs

}