# docker入门教程


## 一、docker基础命令

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

### 构建镜像有三种方式
- 使用docker build命令 和 Dockerfile文件
- 把tar包中的镜像导入
- docker commit 命令

### 2.1、Dockerfile文件编写(把jar包制作成镜像)

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
    ```

4. 运行镜像
    ```cmd
    docker run -d -p 7002:7002 imagename
    ```


### 2.2、把tar包中的镜像导入docker中

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

### 3.1、替换容器里的jar包

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



### 3.2、查看容器日志
```
docker logs --tail=200 -f 容器id
```


### 3.3、目录挂载





### 3.4、一个运行中的Docker容器怎么修改执行run命令时的环境变量

进入到容器中docker exec -it 容器id /bin/bash，再使用export修改就可以了
export ENV='value'

export -p 查看所有环境变量