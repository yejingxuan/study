# idea远程调试java代码


- [idea远程调试java代码](#idea%e8%bf%9c%e7%a8%8b%e8%b0%83%e8%af%95java%e4%bb%a3%e7%a0%81)
  - [概述](#%e6%a6%82%e8%bf%b0)
  - [1、在idea中打开edit configurations选项](#1%e5%9c%a8idea%e4%b8%ad%e6%89%93%e5%bc%80edit-configurations%e9%80%89%e9%a1%b9)
  - [2、新建Remote选项](#2%e6%96%b0%e5%bb%baremote%e9%80%89%e9%a1%b9)
  - [3、配置Remote](#3%e9%85%8d%e7%bd%aeremote)
  - [4、修改远端项目的启动命令，并重新启动](#4%e4%bf%ae%e6%94%b9%e8%bf%9c%e7%ab%af%e9%a1%b9%e7%9b%ae%e7%9a%84%e5%90%af%e5%8a%a8%e5%91%bd%e4%bb%a4%e5%b9%b6%e9%87%8d%e6%96%b0%e5%90%af%e5%8a%a8)
  - [5、debug启动idea中的项目](#5debug%e5%90%af%e5%8a%a8idea%e4%b8%ad%e7%9a%84%e9%a1%b9%e7%9b%ae)

## 概述
由于在开发过程中，某些问题或者功能需要远程调试服务器端代码去分析解决。

## 1、在idea中打开edit configurations选项

![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20191220135658.png)

## 2、新建Remote选项

![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20191220135730.png)

## 3、配置Remote

- ___name___: 可随意填写，建议填写有实际意义的名称
- ___debugger mode___: 选择 attach to remote jvm选项
- ___transport___: 选择socket
- ___host___: 填写远端服务器的ip地址
- ___port___: 端口号可自定义（必须为本地服务和远端服务都没有被占用的端口，防止端口冲突）
- ___command lin arguments for reomte jvm___: idea自动补全,若没有自动补全可手动输入
- ___use module classpath___: 选择需要远端调试的项目

![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20191220140332.png)

## 4、修改远端项目的启动命令，并重新启动

编辑项目的运行命令，加入命令：-jar -Xdebug -Xrunjdwp:server=y,transport=dt_socket,address=5005,suspend=n 

例如springboot项目启动命令为
```shell
java -jar springboot-project.jar
```
更改为
```shell
java -jar -Xdebug -Xrunjdwp:server=y,transport=dt_socket,address=5005,suspend=n -jar springboot-project.jar
```

点击保存，并重启该项目的服务。

ps："address=5005": 5005为在idea中配置remote port时相同的端口。


## 5、debug启动idea中的项目
在idea中打入需要调试的代码断点。选择刚才配置好的remoe（warming）,执行debug启动方式。此时访问远端服务器上部署的项目的相关接口时，就会进入本地idea中的断点调试。

![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20191220140610.png)