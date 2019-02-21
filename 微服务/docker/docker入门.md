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

