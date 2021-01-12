
- [docker简介](#docker%e7%ae%80%e4%bb%8b)
  - [什么是docker](#%e4%bb%80%e4%b9%88%e6%98%afdocker)
- [docker原理](#docker%e5%8e%9f%e7%90%86)
  - [docker镜像分层原理-layer](#docker%e9%95%9c%e5%83%8f%e5%88%86%e5%b1%82%e5%8e%9f%e7%90%86-layer)
  - [docker镜像瘦身](#docker%e9%95%9c%e5%83%8f%e7%98%a6%e8%ba%ab)
- [docker基本操作](#docker%e5%9f%ba%e6%9c%ac%e6%93%8d%e4%bd%9c)
  - [国内镜像源配置](#%e5%9b%bd%e5%86%85%e9%95%9c%e5%83%8f%e6%ba%90%e9%85%8d%e7%bd%ae)
- [docker-compose容器编排工具](#docker-compose%e5%ae%b9%e5%99%a8%e7%bc%96%e6%8e%92%e5%b7%a5%e5%85%b7)
- [docker UI管理平台](#docker-ui%e7%ae%a1%e7%90%86%e5%b9%b3%e5%8f%b0)
  - [简介 shipyard](#%e7%ae%80%e4%bb%8b-shipyard)
  - [shipyard 安装教程](#shipyard-%e5%ae%89%e8%a3%85%e6%95%99%e7%a8%8b)
- [springboot整合docker](#springboot%e6%95%b4%e5%90%88docker)

## docker简介

### 什么是docker

```
docker
英 [ˈdɒkə(r)]  美 [ˈdɑːkər] 
n. 码头工人、搬运工、物件
```

## docker原理

### docker镜像分层原理-layer


### docker镜像瘦身
>https://blog.csdn.net/kurisu_/article/details/100052611


## docker基本操作

### 国内镜像源配置

在 Linux 环境下，我们可以通过修改 /etc/docker/daemon.json ( 如果文件不存在，你可以直接创建它 ) 这个 Docker 服务的配置文件达到效果。
```json
{
    "registry-mirrors": [
        "https://registry.docker-cn.com"
    ]
}
```

配置后重启 docker 让配置生效
```shell
sudo systemctl restart docker
```

通过 docker info 来查阅当前注册的镜像源列表，验证我们配置的镜像源是否生效
```shell
sudo docker info
```


## docker-compose容器编排工具



## docker UI管理平台

### 简介 shipyard
shipyard是一个开源的docker管理平台，其特性主要包括：

- 支持节点动态集群，可扩展节点的规模（swarm、etcd方案）
- 支持镜像管理、容器管理、节点管理等功能
- 可视化的容器管理和监控管理
- 在线容器console终端

### shipyard 安装教程

> 参考文章: [shipyard中文版安装教程](https://www.fcwys.cc/archives/145.html)

1、下载所需镜像
```shell
docker pull rethinkdb
docker pull microbox/etcd
docker pull shipyard/docker-proxy
docker pull swarm
docker pull dockerclub/shipyard
```

2、修改原安装脚本为中文版安装脚本
```shell
#下载官方脚本
wget https://shipyard-project.com/deploy
若下载失败请使用
wget https://raw.githubusercontent.com/shipyard/shipyard- 
project.com/master/site/themes/shipyard/static/deploy

#替换官方脚本
grep -n shipyard:latest deploy
sed -i 's/shipyard\/shipyard:latest/dockerclub\/shipyard:latest/g' deploy
```

3、设置web访问端口(根据需要修改)
```shell
#检查8080端口是否被占用，若占用需修改端口
yum install -y net-tools    //安装net-tools工具包，若已安装可跳过此步骤
netstat -tlnp | grep 8080   //查看宿主机8080端口是否被占用

#配置修改
grep -n 'PORT:-8080' deploy
SHIPYARD_PORT=${PORT:-8080}
修改为
SHIPYARD_PORT=${PORT:-指定端口}
```

4、安装与删除
```shell
sh deploy                                //安装
cat deploy | ACTION=remove bash          //删除
```

5、使用
```shell
浏览器输入:http://主机IP:8080
默认账号:admin
默认密码:shipyard
```

6、增加节点
```shell
curl https://shipyard-project.com/deploy | ACTION=node DISCOVERY=etcd://主服务器IP:4001 bash 
若下载失败请使用
curl -sSL  https://raw.githubusercontent.com/shipyard/shipyard-project.com/master/site/themes/shipyard/static/deploy | ACTION=node DISCOVERY=etcd://主节点IP:4001 bash -s
```



## springboot整合docker


添加maven插件
```xml
<!--docker插件-->
<plugin>
    <groupId>com.spotify</groupId>
    <artifactId>dockerfile-maven-plugin</artifactId>
    <version>1.4.0</version>
    <configuration>
        <repository>${docker.registry}/${docker.image.prefix}/${project.artifactId}</repository>
        <tag>${project.version}</tag> <!-- 不指定tag默认为latest -->
        <buildArgs>
            <JAR_FILE>${project.build.finalName}.jar</JAR_FILE>
        </buildArgs>
    </configuration>
</plugin>
```

编写Dockerfile文件，放置项目根目录
```Dockerfile
#FROM统一镜像
FROM docker.io/fabric8/java-jboss-openjdk8-jdk

#标注Dockerfile作者
MAINTAINER yjx<jingxuan.ye@qq.com>

#设置国内时区
ENV TZ Asia/Shanghai

#添加设置的参数
ARG JAR_FILE

#添加自己的项目到$PRO_PATH
ADD target/${JAR_FILE}/ $PRO_PATH/app.jar

#暴露端口
EXPOSE 52002

#更改自己的工作目录
WORKDIR $PRO_PATH/

#默认运行的命令
CMD ["java", "-jar", "app.jar"]
```

执行命令
```shell
mvn package -DskipTests

mvn dockerfile:build
```