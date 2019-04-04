# go语言入门实例

## 1、环境搭建（windows）

1.1 go环境安装包下载  https://golang.org/dl/

1.2 环境变量配置，在path环境变量中配置go的安装位置，例如：

![](https://github.com/yejingxuan/pic-depot/blob/master/article-pic/TIM%E6%88%AA%E5%9B%BE20190404101644.png)

1.3 在cmd控制台输入 go version 查看是否安装成功

![](https://github.com/yejingxuan/study/blob/master/go%E8%AF%AD%E8%A8%80/TIM%E6%88%AA%E5%9B%BE20190404101557.png)



## 2、第一个go程序

2.1 新建 main.go文件，代码如下

```go
package main

import "fmt"

func main() {
    /* 第一个go程序 */
   fmt.Println("Hello, yejingxuan!")
}
```

2.2 在cmd控制台进入main.go所在文件夹，执行命令
```
go run main.go
```
控制台输出：Hello, World!




