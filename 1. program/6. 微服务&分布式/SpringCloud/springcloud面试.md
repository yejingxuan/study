# springcloud面试题


> 参考文章：https://www.toutiao.com/a6602104153906872836/


## 一. 什么是微服务

微服务架构是一种架构模式或者说是一种架构风格，它提倡将单一应用程序划分为一组小的服务，每个服务运行在其独立的自己的进程中，服务之间相互协调、互相配合，为用户提供最终价值。服务之间采用轻量级的通信机制互相沟通（通常是基于HTTP的RESTful API）,每个服务都围绕着具体的业务进行构建，并且能够被独立的构建在生产环境、类生产环境等。另外，应避免统一的、集中式的服务管理机制，对具体的一个服务而言，应根据业务上下文，选择合适的语言、工具对其进行构建，可以有一个非常轻量级的集中式管理来协调这些服务，可以使用不同的语言来编写服务，也可以使用不同的数据存储。

## 二. 微服务的缺点

### 1. 运维的新挑战

### 2. 接口的一致性



### 3. 分布式的复杂性
与分布式系统相关的复杂性-这种开销包括网络问题，延迟开销，带宽问题，安全问题。



## 三. dubbo和springcloud的区别

>参考文章：https://blog.csdn.net/anningzhu/article/details/76599875

通过上面再几个环节上的分析，相信大家对Dubbo和Spring Cloud有了一个初步的了解。就我个人对这两个框架的使用经验和理解，打个不恰当的比喻：使用Dubbo构建的微服务架构就像组装电脑，各环节我们的选择自由度很高，但是最终结果很有可能因为一条内存质量不行就点不亮了，总是让人不怎么放心，但是如果你是一名高手，那这些都不是问题；而Spring Cloud就像品牌机，在Spring Source的整合下，做了大量的兼容性测试，保证了机器拥有更高的稳定性，但是如果要在使用非原装组件外的东西，就需要对其基础有足够的了解。


|           |Dubbo         |	Spring Cloud|
|:---:      |:---:	       |:---:	      |
|服务注册中心| Zookeeper    |	Spring Cloud Netflix Eureka|
|服务调用方式|	RPC         |	REST API|
|服务网关   |	无          | Spring Cloud Netflix Zuul
|断路器     |	不完善      |Spring Cloud Netflix Hystrix
|分布式配置 |	无          |Spring Cloud Config
|服务跟踪   |	无          |	Spring Cloud Sleuth
|消息总线   |	无          |Spring Cloud Bus
|数据流     |	无          |Spring Cloud Stream
|批量任务   |	无          |	Spring Cloud Task