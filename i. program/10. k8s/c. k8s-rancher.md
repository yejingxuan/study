# racnher

- [racnher](#racnher)
  - [一、rancher简介](#一rancher简介)
  - [二、rancher部署](#二rancher部署)
    - [1、docker部署rancher](#1docker部署rancher)
    - [2、racher镜像管理](#2racher镜像管理)
  - [三、rancher实践](#三rancher实践)
    - [1、rancher的结构组成](#1rancher的结构组成)
    - [2、添加项目和命名空间](#2添加项目和命名空间)

## 一、rancher简介

## 二、rancher部署
### 1、docker部署rancher
- 运行命令
  ```shell
  docker run --privileged -d -p 4433:443 -p 30080:30080  --name=myrancher rancher/rancher:v2.5-head
  ```
- 访问路径
  > https://localhost:4433/

### 2、racher镜像管理
- racher内部使用的k3s来管理镜像和容器，使用命令为k3s crictl ps，和docker的二级命令差不多
- k3s加载镜像命令([具体用法](https://blog.csdn.net/tongzidane/article/details/114587138)) 
  ```shell
  k3s ctr -n k8s.io i import pause.tar
  ```
- k3s导出镜像
  ```shell
  ctr -n k8s.io i export pause.tar k8s.gcr.io/pause:3.2
  ```

## 三、rancher实践
### 1、rancher的结构组成
- rancher的基本结构可分为
  - 集群
  - 项目
  - 命令空间
  - pods  

  ![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20210407162024.png)

### 2、添加项目和命名空间
![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20210407161705.png)