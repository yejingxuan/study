# k8s部署

- [k8s部署](#k8s部署)
  - [一、linux部署](#一linux部署)
  - [二、windows && mac部署](#二windows--mac部署)
    - [1、需安装 Docker Desktop 的 Mac 或者 Windows 版本](#1需安装-docker-desktop-的-mac-或者-windows-版本)
    - [2、开启 Kubernetes](#2开启-kubernetes)
    - [3、下载 k8s-for-docker-desktop 的安装脚本](#3下载-k8s-for-docker-desktop-的安装脚本)
    - [4、安装k8s所需镜像](#4安装k8s所需镜像)
    - [5、配置 Kubernetes dashboard 控制台](#5配置-kubernetes-dashboard-控制台)
    - [6、istio 部署](#6istio-部署)
      - [](#)
  - [三、常用指令](#三常用指令)
    - [1、命名空间](#1命名空间)
    - [2、pods](#2pods)
  - [四、运维](#四运维)
    - [1、nginx服务部署](#1nginx服务部署)

## 一、linux部署


## 二、windows && mac部署

### 1、需安装 Docker Desktop 的 Mac 或者 Windows 版本
![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20210224135516.png)

### 2、开启 Kubernetes
为 Docker daemon 配置镜像加速，参考[阿里云镜像服务](https://cr.console.aliyun.com/cn-hangzhou/instances/mirrors) 或中科大镜像加速地址```https://docker.mirrors.ustc.edu.cn```
![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20210224135644.png)

### 3、下载 k8s-for-docker-desktop 的安装脚本
- git地址：https://github.com/AliyunContainerService/k8s-for-docker-desktop

### 4、安装k8s所需镜像
- 在 Mac 上执行如下脚本
  ```bash
  ./load_images.sh
  ```
- 在Windows上，使用 PowerShell
  ```shell
  .\load_images.ps1
  ```
- 如果因为安全策略无法执行 PowerShell 脚本，请在 “以管理员身份运行” 的 PowerShell 中执行 ```Set-ExecutionPolicy RemoteSigned``` 命令


### 5、配置 Kubernetes dashboard 控制台
- 安装dashboard、可在 二、3步骤中获取 kubernetes-dashboard.yaml
  ```shell
  kubectl apply -f kubernetes-dashboard.yaml
  ```

- 检查 kubernetes-dashboard 应用状态
  ```shell
  kubectl get pod -n kubernetes-dashboard
  ```

- 开启 API Server 访问代理
  ```shell
  kubectl proxy
  ```

- 通过如下 URL 访问 Kubernetes dashboard
  http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/

- 配置控制台访问令牌
  - 对于Mac环境
    ```shell
    TOKEN=$(kubectl -n kube-system describe secret default| awk '$1=="token:"{print $2}')
    kubectl config set-credentials docker-for-desktop --token="${TOKEN}"
    echo $TOKEN
    ```

  - 对于Windows环境
    ```shell
    $TOKEN=((kubectl -n kube-system describe secret default | Select-String "token:") -split " +")[1]
    kubectl config set-credentials docker-for-desktop --token="${TOKEN}"
    echo $TOKEN
    ```

- 登录dashboard的时候
  ![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20210224140417.png)


### 6、istio 部署

- windows & mac 手动部署
  - 下载地址 https://github.com/istio/istio/releases
  - 解压文件夹
  - 配置istio的bin目录到环境变量，使 istioctl 命令生效
  - 安装istio
    ```
    istioctl manifest apply --set profile=demo
    ```
  - 检查 Istio 状态
    ```
    kubectl get pods -n istio-system
    ```
  - 为 ```default``` 名空间开启自动 sidecar 注入
    ```shell
    kubectl label namespace default istio-injection=enabled
    ```
  - 查看各个namespace的istio-injection状态
    ```shell
    kubectl get namespace -L istio-injection
    ```
    ![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20210311111258.png)

#### 

- 通过helm部署
  ```shell
  # 创建istio-system命名空间
  kubectl create namespace istio-system
  # 创建istio CRD
  helm template install/kubernetes/helm/istio-init --namespace istio-system | kubectl apply -f -
  # 检查CRD是否部署完成
  kubectl -n istio-system wait --for=condition=complete job --all
  # 部署istio
  helm template install/kubernetes/helm/istio --namespace istio-system | kubectl apply -f -
  ```

## 三、常用指令

```shell
# 查看service、pod、deploy
kubectl get svc,pod,deploy

# 查看路由、VirtualService
kubectl get gateway,VirtualService

```

### 1、命名空间
- 查看命名空间
  >  kubectl get namespace
- 创建命名空间
  > kubectl create namespace default-mem-example

### 2、pods
- 列出所有命名空间下的所有 pod
  > kubectl get pods --all-namespaces

- 列出指定命名空间的 pods  
  > kubectl get pods -n kubernetes-dashboard


## 四、运维

### 1、nginx服务部署
- 创建nginx deployment & pods
  ```shell
  kubectl create deployment my-nginx --image=nginx
  # 查看已创建pod的yaml详细信息
  kubectl get pod pod_name -o yaml
  ```
  ![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20210305114023.png)

- 创建nginx service 并暴露端口出来
  ```shell
  kubectl expose deployment my-nginx --port=80 --type=NodePort
  ```