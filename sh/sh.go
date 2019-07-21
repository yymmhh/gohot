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

//读取重启命令
func readSh() string {
	data, err := ioutil.ReadFile("./reload.sh")
	if err != nil {

		fmt.Println("读取命令失败!", err)
		color.Set(color.BgRed, color.Bold)
		defer color.Unset()

	}

	return string(data)
}

func ReloadSwoole() {
	sh := readSh()

	command := "/bin/bash"
	params := []string{"-c", sh}

	execCommand(command, params)

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
