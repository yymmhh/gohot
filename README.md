# wl_gohot

用作swoole laravels 的重启


下载地址  https://github.com/yymmhh/gohot/releases/tag/v1.0



可在conf.ini 中定义重启的命令 
已经忽略的目录
 
忽略的文件
  
是否开始日志   


如果为空则监听当前目录,并且会在执行的启动和重启命令自动拼接当前目录
    [listenDir]
    path=/Volumes/E/www/php/aix-system       


重启的命令在  reload.sh    

懒人安装

    wget -O install.sh http://zp.wlphp.cn/install.sh && sudo sh install.sh

需要加入环境变量

    export gohot=/Volumes/E/www/go/src/gohot
    
    export PATH=$PATH:$GOBIN:$GOROOT:$GOTOOLDIR:$php:$gohot


    gohot
    
运行时会自动启动laravels

运行 成功
<img src="http://cxt.cdn.wlphp.cn/gohot_show.png"/>

修改了php 文件 即刻 重新运行命令


======
2019-10-29 
    加入协程管道重启
    
    
编译指令

    go build -o gohot main.go

交叉编译

    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o test_linux_x64 runIP.go
    
    
======
2020-02-16
    大部分代码重构,变为只监听配置的目录的变化


启动时候遇到 too many open files 或者打开文件过多的问题
  https://github.com/yymmhh/gohot/issues/3 
  
  
定时任务 
  在crontable.sh 里面
  
  配置项目的Settings
  CronTableTime 是多少秒执行一次
  
  open 是否开启

    
