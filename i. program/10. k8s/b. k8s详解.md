# k8s入门

- [k8s入门](#k8s入门)
  - [一、概述](#一概述)
    - [1.1、背景&作用](#11背景作用)
    - [1.2、架构](#12架构)
      - [架构图](#架构图)
      - [核心组件](#核心组件)
  - [二、详细](#二详细)
    - [1、pod](#1pod)
    - [2、service](#2service)
    - [3、namespace](#3namespace)
  - [参考文章](#参考文章)

## 一、概述

### 1.1、背景&作用


### 1.2、架构

#### 架构图

![image-20201009191322060](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/image-20201009191322060.png)


#### 核心组件

- ETCD：分布式高性能键值数据库,存储整个集群的所有元数据
- ApiServer:  API服务器,集群资源访问控制入口,提供restAPI及安全访问控制
- Scheduler：调度器,负责把业务容器调度到最合适的Node节点
- Controller Manager：控制器管理,确保集群资源按照期望的方式运行
  - Replication Controller
  - Node controller
  - ResourceQuota Controller
  - Namespace Controller
  - ServiceAccount Controller
  - Tocken Controller
  - Service Controller
  - Endpoints Controller
- kubelet：运行在每运行在每个节点上的主要的“节点代理”个节点上的主要的“节点代理”
  - pod 管理：kubelet 定期从所监听的数据源获取节点上 pod/container 的期望状态（运行什么容器、运行的副本数量、网络或者存储如何配置等等），并调用对应的容器平台接口达到这个状态。
  - 容器健康检查：kubelet 创建了容器之后还要查看容器是否正常运行，如果容器运行出错，就要根据 pod 设置的重启策略进行处理.
  - 容器监控：kubelet 会监控所在节点的资源使用情况，并定时向 master 报告，资源使用数据都是通过 cAdvisor 获取的。知道整个集群所有节点的资源情况，对于 pod 的调度和正常运行至关重要
- kubectl: 命令行接口，用于对 Kubernetes 集群运行命令  https://kubernetes.io/zh/docs/reference/kubectl/ 
- CNI实现: 通用网络接口, 我们使用flannel来作为k8s集群的网络插件, 实现跨节点通信



## 二、详细

### 1、pod
- pod是k8s运行的最小单元；一个pod里可以运行多个docker容器
- k8s控制器决定了创建pod资源的方式和类型，一共有5种控制器类型：
  - Deployment：无状态应用
  - DaemonSet：每台主机部署1个pod（比如守护进程，日志收集进程）
  - StatefulSet：有状态应用（比如MySQL、MongoDB集群）
  - CronJob：定时运行
  - Job：一次性运行
- 无状态pod和有状态pod
  - 但是在实际场景中, 比如:主从关系,主备关系,还有就是数据存储类应用,多个实例通常会在本地磁盘上保存一份数据,而这些实例一旦被杀掉,即使重建出来,实例与数据之间的对应关系也已经丢失,从而导致应用失败.
  - 这种实例之间有不对等关系,或者有依赖关系的应用,被称为"有状态应用"(Stateful Application)
为了能对"有状态应用"做出支持,Kubernetes在Deployment基础上,扩展出了:StatefulSet.
  - RC、Deployment、DaemonSet都是面向无状态的服务，它们所管理的Pod的IP、名字，启停顺序等都是随机的


### 2、service


### 3、namespace

## 参考文章

- [Docker+K8S 集群环境搭建及分布式应用部署](https://blog.csdn.net/ysk_xh_521/article/details/81668631)

- [k8s网络之Flannel网络](https://www.cnblogs.com/goldsunshine/p/10740928.html)

- [深入理解StatefulSet](https://blog.csdn.net/zll_0405/article/details/86683770)