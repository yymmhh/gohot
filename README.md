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

需要加入环境变量

    export gohot=/Volumes/E/www/go/src/gohot
    
    export PATH=$PATH:$GOBIN:$GOROOT:$GOTOOLDIR:$php:$gohot


    gohot
    
运行时会自动启动laravels

运行 成功
<img src="https://github.com/yymmhh/hotswoole/blob/master/show.png"/>

修改了php 文件 即刻 重新运行命令


======
2019-10-29 
    加入协程管道重启