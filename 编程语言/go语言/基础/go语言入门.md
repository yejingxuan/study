# go语言入门实例

- [go语言入门实例](#go%e8%af%ad%e8%a8%80%e5%85%a5%e9%97%a8%e5%ae%9e%e4%be%8b)
  - [1、环境搭建（windows）](#1%e7%8e%af%e5%a2%83%e6%90%ad%e5%bb%bawindows)
    - [1.1 go环境安装包下载](#11-go%e7%8e%af%e5%a2%83%e5%ae%89%e8%a3%85%e5%8c%85%e4%b8%8b%e8%bd%bd)
    - [1.2 环境变量配置](#12-%e7%8e%af%e5%a2%83%e5%8f%98%e9%87%8f%e9%85%8d%e7%bd%ae)
    - [1.3 查看是否安装成功](#13-%e6%9f%a5%e7%9c%8b%e6%98%af%e5%90%a6%e5%ae%89%e8%a3%85%e6%88%90%e5%8a%9f)
  - [2、第一个go程序](#2%e7%ac%ac%e4%b8%80%e4%b8%aago%e7%a8%8b%e5%ba%8f)
    - [2.1 新建 main.go文件](#21-%e6%96%b0%e5%bb%ba-maingo%e6%96%87%e4%bb%b6)
    - [2.2 运行 main.go文件](#22-%e8%bf%90%e8%a1%8c-maingo%e6%96%87%e4%bb%b6)

> 参考文章：[Go 语言极速入门](https://www.cnblogs.com/java-zhao/p/9942311.html)

## 1、环境搭建（windows）

### 1.1 go环境安装包下载  

https://golang.org/dl/

### 1.2 环境变量配置

在path环境变量中配置go的安装位置，例如：

![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20191211161221.png)

### 1.3 查看是否安装成功

在cmd控制台输入 go version

![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20191211161152.png)


## 2、第一个go程序


### 2.1 新建 main.go文件

代码如下

```go
package main

import "fmt"

func main() {
    /* 第一个go程序 */
   fmt.Println("Hello, yejingxuan!")
}
```

### 2.2 运行 main.go文件

在cmd控制台进入main.go所在文件夹，执行命令
```shell
go run main.go
```
控制台输出：Hello, World!



