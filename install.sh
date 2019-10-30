echo "欢迎安装Gohot"

sudo rm -rf /home/linux.zip

sudo wget -P /home/ http://zp.wlphp.cn/linux.zip


sudo rm -rf /usr/local/gohot/
sudo unzip -d /usr/local/gohot /home/linux.zip


sudo rm -rf /usr/bin/gohot
sudo rm -rf /usr/bin/conf.ini

sudo ln -s  /usr/local/gohot/gohot  /usr/bin/gohot