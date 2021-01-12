
# RocketMQ学习

> [分布式开放消息系统(RocketMQ)的原理与实践](https://www.jianshu.com/p/453c6e7ff81c)


## 一、rockerMQ安装

### 1.下载安装包rocketmq-all-4.4.0-bin-release.zip
http://rocketmq.apache.org/


### 2.上传到服务器

### 3.解压安装包
```
unzip rocketmq-all-4.4.0-bin-release.zip
```

### 4.执行bin目录下的启动命令

```linux
#后台启动name：
nohup sh mqnamesrv &
#后台启动broker：
nohup sh mqbroker -n 127.0.0.1:9876 &

#查看日志
tail -f nohup

# 停止broker服务
sh mqshutdown broker
# 定制nameserver服务
sh mqshutdown namesrv
```

> 如果报内存溢出则修改：bin/ 下的服务启动脚本 runserver.sh 、runbroker.sh 中对于内存的限制


## 二、可视化管理

### 1、项目下载
- github下载rocketMQ配套可视化管理项目：
https://github.com/apache/rocketmq-externals/tree/master/rocketmq-console

- 修改配置文件：application.properties
    ```
    rocketmq.config.namesrvAddr=127.0.0.1:9876
    ```

- 执行maven命令进行编译打成jar包
    ```
    mvn clean package -Dmaven.test.skip=true 
    ```
    > 若打包报错，修改下pom里的私服地址，或者更改pom里相关jar包的版本


- 运行jar包
    ```
    nohup java -jar rocketmq-console-ng-1.0.0.jar --server.port=6001 &
    ```   

- 访问 http://ip:6001 进入可视化页面


## 三、