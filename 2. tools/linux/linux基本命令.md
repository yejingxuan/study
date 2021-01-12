# linux基本命令

- [linux基本命令](#linux基本命令)
  - [一、用户相关 & 权限相关](#一用户相关--权限相关)
  - [二、文件相关](#二文件相关)
        - [查找文件里的内容](#查找文件里的内容)
        - [统计文件内容](#统计文件内容)
        - [解压文件](#解压文件)
        - [dpkg : deb软件包命令](#dpkg--deb软件包命令)
  - [进程相关](#进程相关)
  - [网络相关](#网络相关)
  - [文件相关](#文件相关)
    - [增删改查文件](#增删改查文件)

## 一、用户相关 & 权限相关
- sudo -i
  > 切换到root用户

- sudo user
  > 切换到user用户

- ssh -p 50022 my@127.0.0.1
  > ssh登录

- sudo -i 
  > 免密切换到root用户

- chmod 777 file
  >赋予所有用户该文件可读可写可执行权限

## 二、文件相关

##### 查找文件里的内容
> https://www.cnblogs.com/kerrycode/p/5802420.html

##### 统计文件内容
grep -o "hello" demo.log | wc -l

-c 只显示有多少行匹配 ，而不具体显示匹配的行
-i 在字符串比较的时候忽略大小写
-n 在每一行前面打印该行在文件中的行数

##### 解压文件 
tar zxvf  filename 
> x : 从 tar 包中把文件提取出来  
z : 表示 tar 包是被 gzip 压缩过的，所以解压时需要用 gunzip 解压  
v : 显示详细信息  
f xxx.tar.gz : 指定被处理的文件是 xxx.tar.gz 
 
##### dpkg : deb软件包命令
    > dpkg -i filename.deb ：安装  
    dpkg -l : 显示已安装的




## 进程相关

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



- iostat
  > 相关参数：
    -C 显示CPU使用情况
    -d 显示磁盘使用情况
    -k 以 KB 为单位显示
    -m 以 M 为单位显示
    -N 显示磁盘阵列(LVM) 信息
    -n 显示NFS 使用情况
    -p[磁盘] 显示磁盘和分区的情况
    -t 显示终端和CPU的信息
    -x 显示详细信息
    -V 显示版本信息
cpu属性值说明：

  %user：CPU处在用户模式下的时间百分比。
  %nice：CPU处在带NICE值的用户模式下的时间百分比。
  %system：CPU处在系统模式下的时间百分比。
  %iowait：CPU等待输入输出完成时间的百分比。
  %steal：管理程序维护另一个虚拟处理器时，虚拟CPU的无意识等待时间百分比。
  %idle：CPU空闲时间百分比。
  注：如果%iowait的值过高，表示硬盘存在I/O瓶颈，%idle值高，表示CPU较空闲，如果%idle值高但系统响应慢时，有可能是CPU等待分配内存，此时应加大内存容量。%idle值如果持续低于10，那么系统的CPU处理能力相对较低，表明系统中最需要解决的资源是CPU。
disk属性值说明：

  rrqm/s: 每秒进行 merge 的读操作数目。即 rmerge/s
  wrqm/s: 每秒进行 merge 的写操作数目。即 wmerge/s
  r/s: 每秒完成的读 I/O 设备次数。即 rio/s
  w/s: 每秒完成的写 I/O 设备次数。即 wio/s
  rsec/s: 每秒读扇区数。即 rsect/s
  wsec/s: 每秒写扇区数。即 wsect/s
  rkB/s: 每秒读K字节数。是 rsect/s 的一半，因为每扇区大小为512字节。
  wkB/s: 每秒写K字节数。是 wsect/s 的一半。
  avgrq-sz: 平均每次设备I/O操作的数据大小 (扇区)。
  avgqu-sz: 平均I/O队列长度。
  await: 平均每次设备I/O操作的等待时间 (毫秒)。
  svctm: 平均每次设备I/O操作的服务时间 (毫秒)。
  %util: 一秒中有百分之多少的时间用于 I/O 操作，即被io消耗的cpu百分比
  备注：如果 %util 接近 100%，说明产生的I/O请求太多，I/O系统已经满负荷，该磁盘可能存在瓶颈。如果 svctm 比较接近 await，说明 I/O 几乎没有等待时间；如果 await 远大于 svctm，说明I/O 队列太长，io响应太慢，则需要进行必要优化。如果avgqu-sz比较大，也表示有当量io在等待。
iostat 2 3 每两秒刷新显示 且显示三次
查看TPS和吞吐量 iostat -d -k 1 1
  tps：该设备每秒的传输次数（Indicate the number of transfers per second that were issued to the device.）。“一次传输”意思是“一次I/O请求”。多个逻辑请求可能会被合并为“一次I/O请求”。“一次传输”请求的大小是未知的。
  kB_read/s：每秒从设备（drive expressed）读取的数据量；
  kB_wrtn/s：每秒向设备（drive expressed）写入的数据量；
  kB_read：读取的总数据量；kB_wrtn：写入的总数量数据量；


## 网络相关

- ping IP
  >测试IP是否连通

- ping IP PORT
  >测试IP、端口是否连通

- netstat 命令
  ```
  -a或–all 显示所有连线中的Socket。
  -l或–listening 显示监控中的服务器的Socket。
  -t或–tcp 显示TCP传输协议的连线状况。
  -u或–udp 显示UDP传输协议的连线状况。
  -n或–numeric 直接使用IP地址，而不通过域名服务器。
  ```


## 文件相关

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
