# springcloud入门

- [springcloud入门](#springcloud入门)
  - [一、什么是微服务](#一什么是微服务)
  - [二、微服务的缺点](#二微服务的缺点)
    - [1、运维的新挑战](#1运维的新挑战)
    - [2、接口的一致性](#2接口的一致性)
    - [3、分布式的复杂性](#3分布式的复杂性)
  - [三、dubbo和springcloud的区别](#三dubbo和springcloud的区别)
  - [四、springcloud组件](#四springcloud组件)
    - [1、springcloud提供两种服务调用方法ribbon和feign](#1springcloud提供两种服务调用方法ribbon和feign)

## 一、什么是微服务

微服务架构是一种架构模式或者说是一种架构风格，它提倡将单一应用程序划分为一组小的服务，每个服务运行在其独立的自己的进程中，服务之间相互协调、互相配合，为用户提供最终价值。服务之间采用轻量级的通信机制互相沟通（通常是基于HTTP的RESTful API）,每个服务都围绕着具体的业务进行构建，并且能够被独立的构建在生产环境、类生产环境等。另外，应避免统一的、集中式的服务管理机制，对具体的一个服务而言，应根据业务上下文，选择合适的语言、工具对其进行构建，可以有一个非常轻量级的集中式管理来协调这些服务，可以使用不同的语言来编写服务，也可以使用不同的数据存储。

## 二、微服务的缺点

### 1、运维的新挑战

### 2、接口的一致性



### 3、分布式的复杂性
与分布式系统相关的复杂性-这种开销包括网络问题，延迟开销，带宽问题，安全问题。



## 三、dubbo和springcloud的区别

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


## 四、springcloud组件


### 1、springcloud提供两种服务调用方法ribbon和feign

  Ribbon和Feign都是用于调用其他服务的，不过方式不同。

  1.启动类使用的注解不同，Ribbon用的是@RibbonClient，Feign用的是@EnableFeignClients。

  2.服务的指定位置不同，Ribbon是在@RibbonClient注解上声明，Feign则是在定义抽象方法的接口中使用@FeignClient声明。

  3.调用方式不同，Ribbon需要自己构建http请求，模拟http请求然后使用RestTemplate发送给其他服务，步骤相当繁琐。
  Feign则是在Ribbon的基础上进行了一次改进，采用接口的方式，将需要调用的其他服务的方法定义成抽象方法即可，不需要自己构建http请求。不过要注意的是抽象方法的注解、方法签名要和提供服务的方法完全一致。

  从实践上看，采用feign的方式更优雅（feign内部也使用了ribbon做负载均衡）。


