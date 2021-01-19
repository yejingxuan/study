# go语言开发指南

- [go语言开发指南](#go语言开发指南)
  - [一、quick-start](#一quick-start)
    - [1、 go环境安装包下载](#1-go环境安装包下载)
    - [2、 环境变量配置](#2-环境变量配置)
    - [3、 查看是否安装成功](#3-查看是否安装成功)
    - [4、第一个go程序](#4第一个go程序)
  - [二、内存模型&垃圾回收](#二内存模型垃圾回收)
    - [1、内存模型](#1内存模型)
    - [2、垃圾回收算法](#2垃圾回收算法)
  - [三、数据结构](#三数据结构)
    - [1、基本数据类型](#1基本数据类型)
    - [2、go里面的集合](#2go里面的集合)
    - [3、指针类型](#3指针类型)
  - [四、struct & interface & 面向对象](#四struct--interface--面向对象)
    - [1、struct](#1struct)
    - [2、interface](#2interface)
  - [五、go的并发](#五go的并发)
    - [1、goroutine协程](#1goroutine协程)
    - [2、让协程按顺序执行](#2让协程按顺序执行)
    - [3、go中的那些锁](#3go中的那些锁)
    - [4、other](#4other)
  - [六、包管理](#六包管理)
    - [1、go get](#1go-get)
    - [2、go vender](#2go-vender)
    - [3、go mod](#3go-mod)
    - [4、go build打包](#4go-build打包)
  - [七、项目开发全家桶](#七项目开发全家桶)
    - [1、web框架gin](#1web框架gin)
    - [2、配置文件toml](#2配置文件toml)
    - [3、orm框架gorm](#3orm框架gorm)
    - [4、日志框架logrus](#4日志框架logrus)
    - [5、golang调用python](#5golang调用python)
  - [八、rpc框架grpc-go](#八rpc框架grpc-go)
    - [1、protobuffer传输协议](#1protobuffer传输协议)
    - [2、远程调用原理](#2远程调用原理)
    - [3、ssl秘钥鉴权](#3ssl秘钥鉴权)
    - [4、grpc-gateway开启grpc和http双协议](#4grpc-gateway开启grpc和http双协议)
    - [5、grpc流模式](#5grpc流模式)
    - [6、etcd服务注册中心](#6etcd服务注册中心)
  - [参考文章](#参考文章)


## 一、quick-start

### 1、 go环境安装包下载  

https://golang.org/dl/

### 2、 环境变量配置

在path环境变量中配置go的安装位置，例如：

![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20191211161221.png)

### 3、 查看是否安装成功

在cmd控制台输入 go version

![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20191211161152.png)


### 4、第一个go程序


- 新建 main.go文件
  代码如下

  ```go
  package main

  import "fmt"

  func main() {
      /* 第一个go程序 */
    fmt.Println("Hello, yejingxuan!")
  }
  ```

- 运行 main.go文件
  在cmd控制台进入main.go所在文件夹，执行命令
  ```shell
  go run main.go
  ```
  控制台输出：Hello, World!


## 二、内存模型&垃圾回收
### 1、内存模型

### 2、垃圾回收算法

- go采用的是三色标记法
![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/123.png)


## 三、数据结构

### 1、基本数据类型

### 2、go里面的集合

- 数组

- 切片

- Map

### 3、指针类型

- 使用指针的好处


## 四、struct & interface & 面向对象
### 1、struct

### 2、interface

## 五、go的并发

### 1、goroutine协程
- __协程和线程__

- __协程的启动__
  - go method() go关键字+方法名开启协程
    ```go
    func main() {
      fmt.Println("main start")
      go Hello()
      fmt.Println("main end")
    }

    func Hello() {
      fmt.Println("hello everybody , I'm asong")
    }
    ```

  - go func(){} 匿名函数开启协程
    ```go
    func main() {
      fmt.Println("main start")
      go func() {
		    fmt.Println("hello everybody , I'm asong")
      }()
      fmt.Println("main end")
    }
    ```
  - 两种方法输出结果为：
    ```go
    main start
    main end
    ```
  - 一开始很奇怪，我们的协程里的方法为什么没有输出呢，因为主协程执行完后，程序结束了还没有去调度子协程去执行。这点就很坑爹，那怎么保证我们的子协程在主协程执行完之前去执行呢，这个就涉及到协程的调度，详情请看4.2

- __协程的特性__
  
  - 一个协程占用的空间为2kb，所以协程也不是可以永无止境的去创建

### 2、让协程按顺序执行
上面说到协程的调度，那么如何让协程在我们理想的状态下顺序执行呢？:
- __1) sleep方法休眠主协程__

- __2) chanel的使用__

- __3) sync.WaitGroup__
  - sync.WaitGroup它的使用场景是在一个goroutine等待一组goroutine执行完成。
  - sync.WaitGroup拥有一个内部计数器。当计数器等于0时，则Wait()方法会立即返回。否则它将阻塞执行Wait()方法的goroutine直到计数器等于0时为止。
  - 要增加计数器，我们必须使用Add(int)方法。要减少它，我们可以使用Done()（将计数器减1），也可以传递负数给Add方法把计数器减少指定大小，Done()方法底层就是通过Add(-1)实现的


### 3、go中的那些锁
Go 语言在 sync 包中提供了用于同步的一些基本原语，包括常见的互斥锁 Mutex 与读写互斥锁 RWMutex
- __Mutex__

- __RWMutex__

### 4、other
- __线程安全的map——sync.Map__

- __sync.Pool并发池__
  - sync.Pool是一个并发池，负责安全地保存一组对象


- __sync.Once__
  - sync.Once是一个简单而强大的原语，可确保一个函数仅执行一次


- __sync.Cond__



## 六、包管理
### 1、go get 
### 2、go vender
### 3、go mod
- __设置Module环境变量__
  ```shell
  //linux 
  export GO111MODULE=on
  //windows
  set GO111MODULE=on
  ```
- __网络不好的话设置代理__(设置环境变量GOPROXY的值为 https://goproxy.io 或https://athens.azurefd.net)
  ```shell
  //linux 
  export GOPROXY=https://goproxy.io
  //windows
  go env -w GOPROXY=https://goproxy.cn,direct
  ```

- __常用命令__
  ```go
  go mod init //初始化一个mod版本库
  go mod tidy //加载所需要的go包
  go mod vendor //将依赖包复制到项目下的 vendor目录
  go list -m all //显示依赖关系
  go list -m -json all //显示详细依赖关系
  ```

### 4、go build打包
  ```go
  go build [-o output] [-i] [build flags] [packages]
  //示例--编译cmd目录下的main.go，输出到build目录下并命令为service_name
  go build -o build/service_name cmd/main.go
  ```


## 七、项目开发全家桶
### 1、web框架gin
- __项目源码__
  
> github.com/gin-gonic/gin

- __快速开启gin服务__
  ```go
  func main() {
    engine := gin.Default()
    //V1版本接口定义
    v1 := engine.Group("/service/api/v1")
    {
      v1.GET("/healthCheck", search_svr.HealthCheck)
      v1.GET("/searchRecruitInfo", search_svr.SearchRecruitInfo)
    }
    engine.Run(":8082") 
  }
  ```

- __编写自己的API接口__
  通过看gin框架源码得知，vi.get()需传入func(*Context)类型的method
  ```go
  type HandlerFunc func(*Context)
  ```
  所以可以自定义自己的API具体实现方法如下：
  ```go
  package search_svr

  import (
    "context"
    "github.com/gin-gonic/gin"
  )

  func HealthCheck(c *gin.Context) {
    rep := gin.H{
      "message": "ok",
      "code":    200,
    }
    c.JSON(200, rep)
  }
  ```


### 2、配置文件toml
- 项目源码
  
> github.com/BurntSushi/toml

- 配置文件config.toml
  ```toml
  [owner]
  name = "Tom Preston-Werner"
  dob = 1979-05-27T07:32:00-08:00 # 注释可以写在这里

  [database]
  server = "192.168.1.1"
  ports = [ 8001, 8001, 8002 ]
  connection_max = 5000
  enabled = true
  ```

- 编写go的配置信息结构体
  ```go
  package config

  import "time"

  type Config struct {
    DB database `toml:"database"` //可通过配置关系映射来取别名
    Owner ownerInfo //默认忽略首字母大小写
  }

  type ownerInfo struct {
    Name string
    DOB time.Time
  }

  type database struct {
    Server string
    Ports []int
    ConnMax int `toml:"connection_max"`
    Enabled bool
  }
  ```

- 加载配置文件，通过读取config.toml配置，加载配置项到config结构体中
  ```go
  import (
    "github.com/BurntSushi/toml"
  )

  func initConfig() {
    var conf config.Config
    // 通过toml.DecodeFile将toml配置文件的内容，解析到struct对象
    if _, err := toml.DecodeFile("config/config.toml", &conf); err != nil {
      // handle error
      log.Printf("err: %s", err)
    }

    // 可以通过conf读取配置
    log.Printf("conf-info:", conf)
  }
  ```


### 3、orm框架gorm
- __项目源码__
  
> github.com/go-gorm/gorm

- __DB连接池__
  - 在高并发实践中，为了提高数据库连接的使用率，避免重复建立数据库连接带来的性能消耗，会经常使用数据库连接池技术来维护数据库连接。
  - gorm自带了数据库连接池使用非常简单只要设置下数据库连接池参数即可。
  - 注意：使用连接池技术后，千万不要使用完db后调用db.Close关闭数据库连接，这样会导致整个数据库连接池关闭，导致连接池没有可用的连接。
  ```go
  package tools

  //定义全局的db对象，我们执行数据库操作主要通过他实现。
  var _db *gorm.DB

  //包初始化函数，golang特性，每个包初始化的时候会自动执行init函数，这里用来初始化gorm。
  func init() {
      var err error
      _db, err = gorm.Open("mysql", dsn)
      if err != nil {
      panic("连接数据库失败, error=" + err.Error())
  }
    
    //设置数据库连接池参数
    _db.DB().SetMaxOpenConns(100)   //设置数据库连接池最大连接数
    _db.DB().SetMaxIdleConns(20)   //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
  }

  //获取gorm db对象，其他包需要执行数据库查询的时候，只要通过tools.getDB()获取db对象即可。
  //不用担心协程并发使用同样的db对象会共用同一个连接，db对象在调用他的方法的时候会从数据库连接池中获取新的连接
  func GetDB() *gorm.DB {
    return _db
  }
  ```

### 4、日志框架logrus



### 5、golang调用python

```
exec: "gcc": executable file not found in %PATH%
```



```
Python.h: No such file or directory
```





## 八、rpc框架grpc-go

### 1、protobuffer传输协议

### 2、远程调用原理

### 3、ssl秘钥鉴权

### 4、grpc-gateway开启grpc和http双协议

### 5、grpc流模式

### 6、etcd服务注册中心



## 参考文章

- [Go 语言极速入门](https://www.cnblogs.com/java-zhao/p/9942311.html)

- [Go语言sync包的应用详解](https://juejin.im/post/6844904147880263694)

- [grpc之 普通流 、服务端流、 客户端流 、双向流模式](https://www.cnblogs.com/sunlong88/p/12128750.html)

- [Golang序列教程](https://www.tizi365.com/archives/6.html)

- [图文结合，白话Go的垃圾回收原理](https://juejin.im/post/6882206650875248654)

- [聊聊自己学Go一年来的经历与成长](https://juejin.im/post/6863680036407345166)
