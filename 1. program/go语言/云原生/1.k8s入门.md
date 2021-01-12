# k8s入门



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



## 二、部署

### 2.1、单机搭建

1.安装docker

```shell
yum install docker
#启动服务
systemctl start docker.service
systemctl enable docker.service
#测试
docker version123456
```

2.安装etcd

```shell
yum install etcd -y
#启动etcd
systemctl start etcd
systemctl enable etcd
#输入如下命令查看 etcd 健康状况
etcdctl -C http://localhost:2379 cluster-health
#安装 Kubernetes
yum install kubernetes -y
```

3.安装k8s

```shell
#安装 Kubernetes
yum install kubernetes -y
```

4.然后分别启动以下程序（Master）

```shell
systemctl start kube-apiserver
systemctl enable kube-apiserver
systemctl start kube-controller-manager
systemctl enable kube-controller-manager
systemctl start kube-scheduler
systemctl enable kube-scheduler
```

5.接下来启动 Node 节点的程序：

```shell
systemctl start kubelet
systemctl enable kubelet
systemctl start kube-proxy
systemctl enable kube-proxy
```

6.这样，一个简单的 K8S 集群环境就已经搭建完成了，我们可以运行以下命令来查看集群状态。

```shell
kubectl get no
```

7.创建覆盖网络 flannel对集群中 pod 的网络进行统一管理

```shell
##1.安装 flannel：
yum install flannel -y1

##2.编辑文件 /etc/sysconfig/flanneld，增加以下代码,其中 –iface 对应的是网卡的名字。
--logtostderr=false --log_dir=/var/log/k8s/flannel/ --etcd-prefix=/atomic.io/network  --etcd-endpoints=http://localhost:2379 --iface=enp0s31

##3.配置 etcd 中关于 flanneld 的 key,flannel 使用 etcd 进行配置，来保证多个 flannel 实例之间的配置一致性，所以需要在 etcd 上进行如下配置：/atomic.io/network/config 这个 key 与上文 /etc/sysconfig/flannel 中的配置项 FLANNEL_ETCD_PREFIX 是相对应的，错误的话启动就会出错）Network 是配置网段，不能和物理机 IP 冲突，可以随便定义，尽量避开物理机 IP 段。
etcdctl mk /atomic.io/network/config '{ "Network": "10.0.0.0/16" }'

##4.启动修改后的 flannel ，并依次重启 docker、kubernete：
systemctl enable flanneld  
systemctl start flanneld
service docker restart
systemctl restart kube-apiserver
systemctl restart kube-controller-manager
systemctl restart kube-scheduler
systemctl enable flanneld
systemctl start flanneld
service docker restart
systemctl restart kubelet
systemctl restart kube-proxy1234567891011
```

这样，我们将应用部署到 Docker 容器中时，就可以通过物理IP访问到容器了。



## 参考文章

- [Docker+K8S 集群环境搭建及分布式应用部署](https://blog.csdn.net/ysk_xh_521/article/details/81668631)

- [k8s网络之Flannel网络](https://www.cnblogs.com/goldsunshine/p/10740928.html)