package main

import (
	"bufio"
	"fmt"
	"github.com/syyongx/php2go"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

func httpPostForm(ipJson string) {
	params := url.Values{}
	// params.Set("hello","fdsfs")  //这两种都可以
	params = url.Values{"key": {ipJson}, "id": {"123"}}
	resp, _ := http.PostForm("http://zp.wlphp.cn/index.php", params)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

}

var contentArray = make([]string, 0, 5)

func main() {

	sh := "ifconfig"

	command := "/bin/bash"
	params := []string{"-c", sh}

	contentArray = contentArray[0:0]
	cmd := exec.Command(command, params...)
	//显示运行的命令
	fmt.Printf("执行命令: %s\n", strings.Join(cmd.Args[1:], " "))
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error=>", err.Error())

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

		//fmt.Printf("\n %c[1;40;32m%s%c[0m\n\n", 0x1B, line, 0x1B)
		index++
		contentArray = append(contentArray, line)
	}
	fmt.Println(php2go.Date("2006/01/02/ 15:04:05 PM", php2go.Time()))

	str := fmt.Sprintf("%s", contentArray)

	httpPostForm(str)

	cmd.Wait()

}
