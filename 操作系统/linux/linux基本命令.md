## linux基本命令

- [linux基本命令](#linux%e5%9f%ba%e6%9c%ac%e5%91%bd%e4%bb%a4)
  - [用户相关](#%e7%94%a8%e6%88%b7%e7%9b%b8%e5%85%b3)
  - [权限相关](#%e6%9d%83%e9%99%90%e7%9b%b8%e5%85%b3)
  - [文件相关](#%e6%96%87%e4%bb%b6%e7%9b%b8%e5%85%b3)
  - [进程相关](#%e8%bf%9b%e7%a8%8b%e7%9b%b8%e5%85%b3)
  - [网络相关](#%e7%bd%91%e7%bb%9c%e7%9b%b8%e5%85%b3)

## 用户相关
- sodu -i： 切换到root用户

### 用户相关

- ssh -p 50022 my@127.0.0.1
  >ssh登录

- sudo -i 
  >免密切换到root用户

## 文件相关

- tar zxvf  filename : 解压文件

    > x : 从 tar 包中把文件提取出来  
    z : 表示 tar 包是被 gzip 压缩过的，所以解压时需要用 gunzip 解压  
    v : 显示详细信息  
    f xxx.tar.gz : 指定被处理的文件是 xxx.tar.gz 
 
- dpkg : deb软件包命令
    > dpkg -i filename.deb ：安装  
    dpkg -l : 显示已安装的

## 进程相关

### 权限相关

- chmod 777 file
  >赋予所有用户该文件可读可写可执行权限


### 进程相关

- ps -aux | grep mysql
   >查看mysql相关进程的详细信息和占用内存

- netstat -apn|grep 8080
  >查看某个端口占用的进程信息

- top
  >查看占用资源最高的进程

- top -p pid
  >查看进程号为pid的进程占用资源信息

- ps -ef|grep java
  >查看java占用的进程

### 网络相关

- ping IP
  >测试IP是否连通

- ping IP PORT
  >测试IP、端口是否连通


### 文件相关

- scp root@107.172.27.254:/home/test.txt 
  >下载文件

- scp -r root@107.172.27.254:/home/test
  >下载目录

- scp test.txt root@107.172.27.254:/home  
  >上传文件

- scp -r test root@107.172.27.254:/home
  >上传目录


### 增删改查文件

- tail -n 20 -f  filename
  >查看文件后20行并进行追踪
