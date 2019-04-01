# docker入门教程

## 一、docker基础命令

###  docker服务启动命令
```
service docker start
```

###  docker镜像拉取命令
```docker
docker search 镜像名称
docker pull 镜像名称
``` 

###  docker容器启动命令

```linux
docker run   --name DIYname  -d imagename:版本号 /bin/bash
```

- i: 以交互模式运行容器，通常与 -t 同时使用；  
- -d: 后台运行容器，并返回容器ID；  
- --name="nginx-lb": 为容器指定一个名称；  
- -p: 端口映射，格式为：主机(宿主)端口:容器端口  
- -m :设置容器使用内存最大值；

### docker容器信息获取
```docker
docker inspect 容器id
```

### docker容器内进程查看
```docker
docker top 容器id
```


### docker容器停止命令
```docker
docker stop 容器id
docker kill 容器id
``` 

### docker容器删除命令
```docker
#删除前必须先停止容器
docker rm 容器id
``` 


## 二、构建镜像

构建镜像有两种方式
- 使用docker commit命令
- 使用docker build命令 和 Dockerfile文件

### 2.1、DockerFile文件编写

```DockerFile
#FROM统一镜像
FROM docker.io/fabric8/java-jboss-openjdk8-jdk
#标注Dockerfile作者
MAINTAINER yjx<jingxuan.ye@qq.com>

#添加自己的项目到$PRO_PATH
ADD demo.jar/ $PRO_PATH/

#暴露端口
EXPOSE 8080

#更改自己的工作目录
WORKDIR $PRO_PATH/

#默认运行的命令
CMD java -jar demo.jar

```




## 三、扩展

### 3.1 替换容器里的jar包

1> 进入镜像
```
docker exec -it  容器id /bin/bash
```

2> 到镜像内把我们的jar包删除掉
```
rm test-1.0.jar
```

3> 将jar包拷贝到容器内
```
docker cp /home/test-1.0.jar  容器id:/opt/project
```

4> 提交镜像
```
docker commit  容器id  docker.io/yjx-testjar:新容器id
```
