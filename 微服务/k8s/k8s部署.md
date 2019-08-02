# kubernetes部署

[TOC]


## 安装kubernetes



## 安装dashboard可视化页面


## 卸载kubernetes

- 首先清理运行到k8s群集中的pod，使用，
  ```shell
  kubectl delete node --all
  ```
- 使用脚本停止所有k8s服务
  ```shell
  yum -y remove kubernetes #if it's registered as a service
  ```