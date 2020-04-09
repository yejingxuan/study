# docker常用镜像部署

[TOC]

## 概述

1. 由于docker默认使用的官方镜像源去拉去镜像，速度太慢，难以忍受，这里推荐使用国内镜像库，DaoCloud Hub,网址：https://hub.daocloud.io， 基本上常用镜像都可以找到。

2. 由于使用docker run命令部署太过蛋疼，部署以后过段时间根本不知道自己创建的容器配置是什么，目录挂载在哪。所以此处均采用docker-compose


## docker部署mysql

- 拉取镜像
  这里使用的是mysql5.6，也可以拉取其他版本
    ```shell
    docker pull daocloud.io/library/mysql:5.6
    ```

- 创建容器编排文件docker-compose.yml
  ```yml
  version: '3'
  services:
    yjx-mysql:
      image: daocloud.io/library/mysql:5.6
      ports:
        - "3306:3306"
      # 目录挂载
      volumes:
        - /jxye/docker/mysql/data:/var/lib/mysql
        - /jxye/docker/mysql/conf:/etc/mysql/conf.d
        - /jxye/docker/mysql/logs:/logs
      # 密码设置
      environment:
        MYSQL_ROOT_PASSWORD: 123456
      container_name: yjx-mysql
  ```

- 使用docker-compose启动容器，在docker-compose.yml目录执行下面命令
  ```shell
  docker-compose up -d 
  ```

- 启动后执行命令,查看容器是否启动成功
  ```shell
  docker ps 
  ```
- 若启动失败，查看容器日志定位原因
  ```shell
  docker logs --tail -f 容器id
  ```
- 附加：使用docker run 直接启动容器的方法
  ```docker
  docker run --name yjx-mysql -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password 镜像id
  ```

## docker部署postgresql


- 拉取镜像
  ```shell
  docker pull docker.io/postgres:11 
  ```

- 创建容器编排文件docker-compose.yml
  ```yml
  version: '3'
  services:
    yjx-postgres:
      image: docker.io/postgres:11
      ports:
        - "5432:5432"
      volumes:
        - /jxye/docker/postgres/data:/var/lib/postgresql/data/pgdata
      environment:
        POSTGRES_PASSWORD: 123456
      container_name: yjx-postgres
  ```

- 使用docker-compose启动容器，在docker-compose.yml目录执行下面命令
  ```shell
  docker-compose up -d 
  ```

- 启动后执行命令,查看容器是否启动成功
  ```shell
  docker ps 
  ```
- 若启动失败，查看容器日志定位原因
  ```shell
  docker logs --tail -f 容器id
  ```

- 附加：使用docker run 直接启动容器的方法
  ```docker
  docker run -d --name yjx-pg -p 5432:5432 -e POSTGRES_PASSWORD=123456 docker.io/postgres:11
  ```


## docker部署rabbitMQ

- 拉取镜像
  这里使用的是rabbitmq的3.7-management版本，也可以拉取其他版本，（踩坑记录：只有management版本附带有web管理）
    ```shell
    docker pull daocloud.io/library/rabbitmq:3.7-management
    ```
- 创建容器编排文件docker-compose.yml
  ```yml
  version: '3'
  services:
    yjx-rabbit2:
      image: daocloud.io/library/rabbitmq:3.7-management
      hostname: yjx-rabbit2
      ports:
        - "15671:15671"
        - "15672:15672"
        - "5672:5672"
      environment:
        # web管理页面账号密码设置
        RABBITMQ_DEFAULT_USER: admin
        RABBITMQ_DEFAULT_PASS: admin
      container_name: yjx-rabbit2

  ```

- 使用docker-compose启动容器，在docker-compose.yml目录执行下面命令
  ```shell
  docker-compose up -d 
  ```

- 启动后执行命令,查看容器是否启动成功
  ```shell
  docker ps 
  ```
- 若启动失败，查看容器日志定位原因
  ```shell
  docker logs --tail -f 容器id
  ```
- 附加：使用docker run 直接启动容器的方法
  ```docker
    docker run -d -p 15671:15671 -p 15672:15672 -p 5671:5671 -p 5672:5672 -p 4369:4369 --name yjx-rabbit --hostname yjx-rabbit 镜像id
  ```
## docker部署elasticsearch

- 拉取镜像
  ```shell
  docker pull docker.io/elasticsearch:7.0.1
  ```

- 创建容器编排文件docker-compose.yml
  ```yml
  version: '3'
  services:
    yjx-elasticsearch:
      image: docker.io/elasticsearch:7.0.1
      ports:
        - "9200:9200"
        - "9300:9300"
      volumes:
        - /jxye/docker/elasticsearch/data:/usr/share/elasticsearch/data
        - /home/jxye/docker-servers/elasticsearch7/plugins:/usr/share/elasticsearch/plugins
      environment:
        # 设置es占用的jvm内存大小
        ES_JAVA_OPTS: "-Xms256m -Xmx256m"
        discovery.type: "single-node"
      container_name: yjx-elasticsearch
  ```

- 使用docker-compose启动容器，在docker-compose.yml目录执行下面命令
  ```shell
  docker-compose up -d 
  ```

- 启动后执行命令,查看容器是否启动成功
  ```shell
  docker ps 
  ```
- 若启动失败，查看容器日志定位原因
  ```shell
  docker logs --tail -f 容器id
  ```
- 常见问题
    ```
    StartupException: ElasticsearchException[failed to bind service];
    ```
    解决方案：对外挂的目录执行可读写权限
    ```shell
    chmod 777 data/
    ```


- 附加：使用docker run 直接启动容器的方法
  ```shell
  docker run -d --name yjx-es -p 9200:9200 -p 9300:9300 -e ES_JAVA_OPTS="-Xms256m -Xmx256m" -v /server/yjx-es/config/elasticsearch.yml:/usr/share/elasticsearch/config/elasticsearch.yml -v /server/yjx-es/data:/usr/share/elasticsearch/data 镜像id
  ```

## docker部署redis

- 拉取镜像
    ```shell
    docker pull docker.io/redis:latest
    ```

- 创建容器编排文件docker-compose.yml
  ```yml
  version: '3'
  services:
    yjx-redis1:
      image: docker.io/redis:latest
      ports:
        - "6379:6379"
      restart: always
      # 目录挂载
      volumes:
        - /jxye/docker/redis/data:/data
        - /jxye/docker/redis/conf:/usr/local/etc/redis
      # 设置密码
      command: redis-server --requirepass qwerty123
      container_name: yjx-redis1
  ```

- 使用docker-compose启动容器，在docker-compose.yml目录执行下面命令
  ```shell
  docker-compose up -d 
  ```

- 启动后执行命令,查看容器是否启动成功
  ```shell
  docker ps 
  ```
- 若启动失败，查看容器日志定位原因
  ```shell
  docker logs --tail -f 容器id
  ```



## docker部署kafka（单节点部署）

- 拉取镜像
    ```shell
    docker pull docker.io/zookeeper:3.4
    docker pull docker.io/wurstmeister/kafka:latest
    ```

- 创建容器编排文件docker-compose.yml
  ```yml
  version: '3'
  
  services:
    # zookeeper编排
    zoo1:
      image: docker.io/zookeeper:3.4
      restart: unless-stopped
      hostname: zoo1
      ports:
        - "2181:2181"
      container_name: yjx-zookeeper
  
    # kafka编排
    kafka1:
      image: docker.io/wurstmeister/kafka:latest
      ports:
        - "9092:9092"
      environment:
        KAFKA_ADVERTISED_HOST_NAME: 127.0.0.1
        KAFKA_ZOOKEEPER_CONNECT: "zoo1:2181/kafka"
        KAFKA_BROKER_ID: 1
        KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1
        KAFKA_CREATE_TOPICS: "stream-in:1:1,stream-out:1:1"
        KAFKA_HEAP_OPTS: "-Xmx216M -Xms216M"
      depends_on:
        - zoo1
      container_name: yjx-kafka
  ```


- 使用docker-compose启动容器，在docker-compose.yml目录执行下面命令
  ```shell
  docker-compose up -d 
  ```

- 启动后执行命令,查看容器是否启动成功
  ```shell
  docker ps 
  ```
- 若启动失败，查看容器日志定位原因
  ```shell
  docker logs --tail -f 容器id
  ```



## docker部署tomcat

- 拉取镜像

- 根据镜像生成容器并启动
  ```shell
  docker run -d -p 8080:8080 --name yjx-tomcat 镜像id
  ```

- 使用docker-compose启动容器，在docker-compose.yml目录执行下面命令
  ```shell
  docker-compose up -d 
  ```

- 启动后执行命令,查看容器是否启动成功
  ```shell
  docker ps 
  ```
- 若启动失败，查看容器日志定位原因
  ```shell
  docker logs --tail -f 容器id
  ```