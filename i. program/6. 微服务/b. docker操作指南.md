# docker操作指南

- [docker操作指南](#docker操作指南)
  - [一、quick-start](#一quick-start)
    - [1、docker安装-linux](#1docker安装-linux)
    - [2、docker国内镜像源配置](#2docker国内镜像源配置)
    - [3、docker目录](#3docker目录)
    - [4、docker基础命令](#4docker基础命令)
  - [二、构建镜像](#二构建镜像)
    - [1、构建镜像有三种方式](#1构建镜像有三种方式)
    - [2、Dockerfile文件编写(把jar包制作成镜像)](#2dockerfile文件编写把jar包制作成镜像)
    - [3、把tar包中的镜像导入docker中](#3把tar包中的镜像导入docker中)
  - [三、扩展](#三扩展)
    - [1、替换容器里的jar包](#1替换容器里的jar包)
    - [2、查看容器日志](#2查看容器日志)
    - [3、目录挂载](#3目录挂载)
    - [4、一个运行中的Docker容器怎么修改执行run命令时的环境变量](#4一个运行中的docker容器怎么修改执行run命令时的环境变量)
  - [四、docker原理](#四docker原理)
    - [1、docker镜像分层原理-layer](#1docker镜像分层原理-layer)
    - [2、docker镜像瘦身](#2docker镜像瘦身)
  - [五、docker UI管理平台](#五docker-ui管理平台)
    - [1、简介 shipyard](#1简介-shipyard)
    - [2、shipyard 安装教程](#2shipyard-安装教程)
  - [六、docker与代码的集成，实现CI、CD](#六docker与代码的集成实现cicd)
    - [1、springboot整合docker](#1springboot整合docker)

## 一、quick-start

### 1、docker安装-linux

```shell
# 1、更新yum源
yum clean all
yum update

# 2、yum安装docker
# 设置yum源
yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
# 查看所有仓库中所有docker版本
yum list docker-ce --showduplicates | sort -r
# 安装指定版本
yum install docker-ce-18.03.1.ce-1.el7.centos

# 3、启动并加入开机启动
systemctl start docker
systemctl enable docker

# 4、验证安装是否成功(有client和service两部分表示docker安装启动都成功了)
docker version

# 5、安装docker-compose
curl -L "https://github.com/docker/compose/releases/download/1.23.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

chmod +x /usr/local/bin/docker-compose

```

### 2、docker国内镜像源配置

- 在 Linux 环境下：  
  我们可以通过修改 /etc/docker/daemon.json ( 如果文件不存在，你可以直接创建它 ) 这个 Docker 服务的配置文件达到效果。
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
- window & mac 环境下：
  ![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20210115145700.png)

### 3、docker目录
- 修改docker镜像文件存储的位置：ln -s /data/tools/docke /var/lib/docker
  

### 4、docker基础命令
- docker服务启动命令
    ```
    service docker start
    ```

- docker镜像拉取命令
    ```docker
    docker search 镜像名称
    docker pull 镜像名称
    ```
- docker镜像删除命令
    ```
    docker rmi 镜像id
    ```

- docker容器启动命令

    ```linux
    docker run   --name DIYname  -d -p 7002:7002 imagename:版本号 /bin/bash
    ```

    - i: 以交互模式运行容器，通常与 -t 同时使用；  
    - -d: 后台运行容器，并返回容器ID；  
    - --name="nginx-lb": 为容器指定一个名称；  
    - -p: 端口映射，格式为：主机(宿主)端口:容器端口  
    - -m :设置容器使用内存最大值；

- docker容器信息获取
    ```docker
    docker inspect 容器id
    ```

- docker容器内进程查看
    ```docker
    docker top 容器id
    ```


- docker容器停止命令
    ```docker
    docker stop 容器id
    docker kill 容器id
    ```

- docker容器删除命令
    ```docker
    #删除前必须先停止容器
    docker rm 容器id
    ```


## 二、构建镜像

### 1、构建镜像有三种方式
- 使用docker build命令 和 Dockerfile文件
- 把tar包中的镜像导入
- docker commit 命令

### 2、Dockerfile文件编写(把jar包制作成镜像)

1. 在文件夹下创建Dockerfile
    ```
    touch Dockerfile
    ```


2. Dockerfile内容模板
    ```Dockerfile
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
    #设置时区
    ENV TZ Asia/Shanghai
    #默认运行的命令
    CMD java -jar demo.jar
    ```

    Dockerfile各关键字含义
    
    - FROM：基于哪个镜像来制作
    - MAINTAINER：标注作者
    - ADD：添加本地文件到镜像中
    - EXPOSE：暴露端口
    - WORKDIR： 为RUN 、 CMD 、 ENTRYPOINT 等指令配置工作目
    - CMD：设置容器执行时的默认命令
    - RUN：每条指令将在当前镜像基础上执行，并提交为新的镜像
    - ENV：指定环境变量，会被RUN指令使用，并在容器运行时保存
    - COPY：复制本地主机的 <src> （ 为 Dockerfile 所在目录的相对路径）到容器中的 <dest> （当使用本地目录为源目录时，推荐使用 COPY）
    - VOLUME：创建一个可以从本地主机或其他容器挂载的挂载点，一般用来存放数据库和需要保持的数据等


3. Dockerfile编写完成后执行镜像构建命令

    ```cmd
    docker build -t imagename .
    # 指定dockerfile进行镜像构建
    docker build -f Dockerfile01 -t xxx-image:v2.0.0 .
    ```

4. 运行镜像
    ```cmd
    docker run -d -p 7002:7002 imagename
    ```


### 3、把tar包中的镜像导入docker中

1. 导出镜像
    ```
    docker save -o /jxye/DIYimagename.tar imagename
    docker export -o /jxye/DIYimagename.tar imagename
    ```
    - 其中-o和>表示输出到文件
    - export命令是从容器（container）中导出tar文件，而save命令则是从镜像（images）中导出

2. 导入镜像
    ```
    docker load -i DIYimagename.tar
    docker import /jxye/DIYimagename.tar imagename
    ```
    - 其中-i和<表示从文件输入。会成功导入镜像及相关元数据，包括tag信息


## 三、扩展

### 1、替换容器里的jar包

1. 进入容器内部
    ```
    docker exec -it  容器id /bin/bash
    ```

2. 到容器内把build进去的jar包删除掉
    ```
    rm test-1.0.jar
    ```

3. 将jar包拷贝到容器内
    ```
    docker cp /home/test-1.0.jar  容器id:/opt/project
    ```

4. 提交镜像
    ```
    docker commit  容器id  docker.io/yjx-testjar:新容器id
    ```



### 2、查看容器日志
```
docker logs --tail=200 -f 容器id
```


### 3、目录挂载





### 4、一个运行中的Docker容器怎么修改执行run命令时的环境变量

进入到容器中docker exec -it 容器id /bin/bash，再使用export修改就可以了
export ENV='value'

export -p 查看所有环境变量

## 四、docker原理

### 1、docker镜像分层原理-layer


### 2、docker镜像瘦身
>https://blog.csdn.net/kurisu_/article/details/100052611


## 五、docker UI管理平台

### 1、简介 shipyard
shipyard是一个开源的docker管理平台，其特性主要包括：

- 支持节点动态集群，可扩展节点的规模（swarm、etcd方案）
- 支持镜像管理、容器管理、节点管理等功能
- 可视化的容器管理和监控管理
- 在线容器console终端

### 2、shipyard 安装教程

> 参考文章: [shipyard中文版安装教程](https://www.fcwys.cc/archives/145.html)

a、下载所需镜像
```shell
docker pull rethinkdb
docker pull microbox/etcd
docker pull shipyard/docker-proxy
docker pull swarm
docker pull dockerclub/shipyard
```

b、修改原安装脚本为中文版安装脚本
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

c、设置web访问端口(根据需要修改)
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

d、安装与删除
```shell
sh deploy                                //安装
cat deploy | ACTION=remove bash          //删除
```

e、使用
```shell
浏览器输入:http://主机IP:8080
默认账号:admin
默认密码:shipyard
```

f、增加节点
```shell
curl https://shipyard-project.com/deploy | ACTION=node DISCOVERY=etcd://主服务器IP:4001 bash 
若下载失败请使用
curl -sSL  https://raw.githubusercontent.com/shipyard/shipyard-project.com/master/site/themes/shipyard/static/deploy | ACTION=node DISCOVERY=etcd://主节点IP:4001 bash -s
```


## 六、docker与代码的集成，实现CI、CD

### 1、springboot整合docker

a、 添加maven插件
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

b、 编写Dockerfile文件，放置项目根目录
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

c、 执行命令
```shell
mvn package -DskipTests

mvn dockerfile:build
```