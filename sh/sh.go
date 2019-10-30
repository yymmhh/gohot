package sh

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/syyongx/php2go"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

//读取命令,并且执行,是否拼接当前目录
func readAndRunSh(fileName string, currentPath bool) {
	data, err := ioutil.ReadFile(GetRealPath() + fileName)
	path := ReadConf("listenDir")["path"]

	if err != nil {

		fmt.Println("读取命令失败!", err)
		color.Set(color.BgRed, color.Bold)
		defer color.Unset()

	}
	var shData string = string(data)

	//配置文件中监听目录为空,并且拼接当前目录
	if php2go.Empty(path) && currentPath == true { //如果配置文件为空 就获取当前目录
		path, _ = os.Getwd()
		shData = path + shData

	}

	command := "/bin/bash"
	params := []string{"-c", shData}

	execCommand(command, params)

}

//开始运行时候就启动laravels

func StartSwoole() {
	readAndRunSh("start.sh", true)

}

var LoadCount int

var LoadCountChan chan int

func ReloadSwoole() {
	LoadCount++

	if LoadCount > 5 {
		fmt.Println("休息一下")
		php2go.Sleep(1)
		LoadCount = 0
	}

	readAndRunSh("reload.sh", true)
	LoadCount--

}

var contentArray = make([]string, 0, 5)

//进行Action
func execCommand(commandName string, params []string) bool {
	contentArray = contentArray[0:0]
	cmd := exec.Command(commandName, params...)
	//显示运行的命令
	fmt.Printf("执行命令: %s\n", strings.Join(cmd.Args[1:], " "))
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error=>", err.Error())
		return false
	}
	cmd.Start() // Start开始执行c包含的命令，但并不会等待该命令完成即返回。Wait方法会返回命令的返回状态码并在命令返回后释放相关的资源。

	reader := bufio.NewReader(stdout)

	var index int
	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}

		fmt.Printf("\n %c[1;40;32m%s%c[0m\n\n", 0x1B, line, 0x1B)
		index++
		contentArray = append(contentArray, line)
	}
	fmt.Println(php2go.Date("2006/01/02/ 15:04:05 PM", php2go.Time()))

	cmd.Wait()
	return true
}
