- [微服务方案](#%e5%be%ae%e6%9c%8d%e5%8a%a1%e6%96%b9%e6%a1%88)
  - [1、RPC框架](#1rpc%e6%a1%86%e6%9e%b6)
  - [2、服务注册发现](#2%e6%9c%8d%e5%8a%a1%e6%b3%a8%e5%86%8c%e5%8f%91%e7%8e%b0)
  - [3、API网关路由](#3api%e7%bd%91%e5%85%b3%e8%b7%af%e7%94%b1)
  - [4、负载均衡](#4%e8%b4%9f%e8%bd%bd%e5%9d%87%e8%a1%a1)
  - [5、配置中心](#5%e9%85%8d%e7%bd%ae%e4%b8%ad%e5%bf%83)
  - [6、熔断限流](#6%e7%86%94%e6%96%ad%e9%99%90%e6%b5%81)
  - [7、链路追踪](#7%e9%93%be%e8%b7%af%e8%bf%bd%e8%b8%aa)
  - [9、任务调度](#9%e4%bb%bb%e5%8a%a1%e8%b0%83%e5%ba%a6)
  - [10、分布式事务](#10%e5%88%86%e5%b8%83%e5%bc%8f%e4%ba%8b%e5%8a%a1)
  - [11、分布式锁](#11%e5%88%86%e5%b8%83%e5%bc%8f%e9%94%81)
  - [12、分布式ID生成](#12%e5%88%86%e5%b8%83%e5%bc%8fid%e7%94%9f%e6%88%90)
  - [13、一致性HASH算法](#13%e4%b8%80%e8%87%b4%e6%80%a7hash%e7%ae%97%e6%b3%95)
  - [14、消息总线](#14%e6%b6%88%e6%81%af%e6%80%bb%e7%ba%bf)
  - [15、认证授权](#15%e8%ae%a4%e8%af%81%e6%8e%88%e6%9d%83)
  - [16、日志监控](#16%e6%97%a5%e5%bf%97%e7%9b%91%e6%8e%a7)
  - [17、告警](#17%e5%91%8a%e8%ad%a6)
  - [18、CI/CD](#18cicd)
- [微服务常用架构](#%e5%be%ae%e6%9c%8d%e5%8a%a1%e5%b8%b8%e7%94%a8%e6%9e%b6%e6%9e%84)

### 微服务方案

#### 1、RPC框架
- springcloud
- dubbo
- grpc
- httpclient


#### 2、服务注册发现
- zookeeper
- etcd
- eureka
- nacos

#### 3、API网关路由
- kong
- zuul
- gateway

#### 4、负载均衡
- ribbon
- nginx
- HAProxy


#### 5、配置中心
- spring cloud config
- apollo
- nacos

#### 6、熔断限流

- hystrix
- kong
- sentinel
- ratelimit

#### 7、链路追踪
- zipkin
- sleuth
- skyWalking

#### 9、任务调度
- quartz
- schedule
- elastic-job
- xxl-job

#### 10、分布式事务
- 阿里seata

#### 11、分布式锁
- redisson框架

#### 12、分布式ID生成
- 雪花算法

#### 13、一致性HASH算法
- 一致性hash算法，hash倾斜问题解决等等

#### 14、消息总线
- spring cloud bus

#### 15、认证授权
- spring security/spring cloud security
- oauth2:
- OIDC: OIDC=(Identity, Authentication)



#### 16、日志监控
- ELK：ElasticSearch+Logstash+Kibana

#### 17、告警
- sentry

#### 18、CI/CD
- k8s+docker
- jenkins

### 微服务常用架构
