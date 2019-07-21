# hotswoole

用作swoole 的重启

先后台运行 laravels 的启动

    bin/laravels start &
    
    
    /Volumes/E/www/php/aix-system/bin/laravels reload
    [2019-07-21 15:44:28] [INFO] Swoole [PID=2847] is reloaded.
    

可在conf.ini 中定义重启的命令 
已经忽略的目录
 
忽略的文件
  
是否开始日志   


如果为空则监听当前目录
    [listenDir]
    path=/Volumes/E/www/php/aix-system       


重启的命令在  reload.sh    

需要加入环境变量

    export hotswoole=/Volumes/E/www/go/src/wl_HotSwoole
    
    export PATH=$PATH:$GOBIN:$GOROOT:$GOTOOLDIR:$php:$hotswoole
