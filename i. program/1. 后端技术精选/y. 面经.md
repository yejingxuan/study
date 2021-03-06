
- [自我介绍](#%e8%87%aa%e6%88%91%e4%bb%8b%e7%bb%8d)
- [面试官问你还有什么要问我的](#%e9%9d%a2%e8%af%95%e5%ae%98%e9%97%ae%e4%bd%a0%e8%bf%98%e6%9c%89%e4%bb%80%e4%b9%88%e8%a6%81%e9%97%ae%e6%88%91%e7%9a%84)
  - [问技术面试官问题如下](#%e9%97%ae%e6%8a%80%e6%9c%af%e9%9d%a2%e8%af%95%e5%ae%98%e9%97%ae%e9%a2%98%e5%a6%82%e4%b8%8b)
  - [问hr问题如下：](#%e9%97%aehr%e9%97%ae%e9%a2%98%e5%a6%82%e4%b8%8b)
- [](#)


### 自我介绍






### 面试官问你还有什么要问我的

#### 问技术面试官问题如下
RD:
1. 我加入的部门业务具体是什么，主要做产品研发还是项目交付工作？
2. 系统架构具体是什么？
3. 中间件、数据库、ORM、前段框架等是否是高并发业务？
4. 工作强度，加班福利？
5. 我在项目组承担什么角色？
6. 开发流程是什么？ 需求讲解、开发 测试？

#### 问hr问题如下：
HR: 
1. 是否是高强度工作，弹性打卡？
2. 试用期考核机制，个人绩效机制？
3. 晋升涨薪机制，个人职业发展路线？
4. 五险一金缴纳比例？
5. 年终奖？
6. 团建，补助，过节费？

### 
1. 我加入的部门和所在岗位具体是哪个，是做第三方项目交付还是公司的产品业务。
2. 我在工作中扮演的角色和所承担的任务具体是什么。
3. 试用期考核机制，个人的绩效KPI机制。
4. 晋升涨薪机制，公司对个人职业发展的规划和方向。
5. 是否是高强度工作，弹性打卡，是否有加班费。
6. 五险一金缴纳比例，公司福利和年终奖机制。


# 面经



## KSKJ面经

- 设计模式

    - spring框架中常用的设计模式
    - 具体业务场景探讨设计模式，谈谈遇到的最复杂的业务场景   

- 多线程

   - 哪几种线程池
   - wait和sleep的区别
   - 常用的线程池   

- springcloud常用组件介绍
   
   - eurka
   - forign
   - ribbon
   - zuul
   - config
   - Hystrix


- 常用消息队列

  - 各类消息队列的吞吐量

  - 消息队列的两种模式：点对点；发布订阅

- 高并发下的限流操作，流量削峰
  
  - 如何利用redis限流
  - 阻塞队列

- spring知识

  - 常用注解介绍


- zookeeper知识

  - 如何实现分布式锁

- docker知识

  - 如何把jar包打成一个镜像

- es基础知识

  - 如何创建索引

- java集合基础知识

  - arraylist和linkedlist区别
  - hashmap底层实现原理

- 如何保证service层的成员变量线程安全

## BFD面经

- spring加载机制

- redis的数据类型和运行机制，集群部署两种方式：主从复制；哨兵机制

- 缓存实现策略

- zookeeper实现原理和使用场景

- 常用设计模式

- springcloud

- mpp的底层实现

- hashmap和hashtable的区别

- currenthashmap的实现原理，和加锁机制，1.8前后有什么区别

- mybatis的#和$有什么区别

- pagehelper的底层实现原理

- jvm虚拟机内存模型



## GSX

- 常用的数据结构
  - list、map、set
  - arraylist和linkedlist的区别，实现原理，ArrayList的扩容机制

- JVM的垃圾回收机制
  - cms、G1等垃圾回收器

- 多线程
  - 线程的创建方式、线程池的参数、线程池的工作流程（队列任务满了如何处理）
  - start 和 run 的区别

- redis
  - redis的基本数据结构、Hash的底层实现

- zookeeper
  - zk的选主策略

- es
  - es的存储结构，索引

- springcloud
  - springcloud和springboot的区别
  - springcloud的常用组件：eurkea、zuul、feign、ribbon、config、Hystrix

- 笔试算法（青蛙跳问题，斐波那契数列）



## HTJY面经
- redis分布式锁实现
- redis缓存、缓存雪崩、缓存穿透
  - 缓存雪崩解决方案
  - 布隆过滤器解决缓存穿透问题
  - 布隆过滤器原理
  - redis如何实现接口限流
- elasticsearch
  - java操作es的常用框架
  - es在生产中如何部署的，分片、副本里的数据如何划分
  - es请求的时候，请求过程。
  - es的match和filter的区别
  - es的score是怎么计算的
- 线程池
  - 线程池参数
  - 常用线程池
  - 线程池工作流程
  - 线程池初始的时候有多少线程
  - 指令重排
  - 线程的运行状态，以及这些状态流程转换的过程
- mysql数据库
  - 为什么索引用b+tree，和b-tree的区别
  - 常用的索引类型，有什么区别
  - sql执行计划，有哪些参数，优化经验
- git
  - git的常用命令
  - 公司git的使用流程，分支规范
- 算法
  - 快速排序、冒泡排序以及他们的时间复杂度


## QTY面经
- final关键字
- 对象引用
- JVM调优
- synchronize锁原理，如何实现可重入
- reentlock如何实现可重入
- synchronize和reentlock区别
- volatile原理和作用
- 线程池参数
- list有哪些实现，底层原理，应用场景
  - 有哪些线程安全的list，Vector过时了，还有没有别的线程安全的list
  - arraylist 和 linkedlist 查找元素的时间复杂度
- set有哪些实现，底层原理，应用场景
  - treeset和linkedhaspset区别
  - linkedhashset如何实现有序，底层是咋实现的
- 讲讲IOC
- spring管理bean的实现类，默认用的哪个
- spring依赖注入有哪些方式
  - 构造函数注入有了解过吗
- spring的bean单例，多例的应用场景
- 高并发下如何去统计接口的访问次数


## STWL
- 讲讲项目架构，springcloud和中台的设计方案
- eureka的作用，服务注册原理，A服务调用B服务的请求调用流程
- 分布式事务
  - 讲讲项目中是咋实现的（本地消息表实现事务最终一致性）
  - 直播场景中如何保证金钱，积分之类的事务强一致性（TCC）
- 高并发
  - 直播场景中如何保证用户频繁刷礼物的状态下，请求的顺序性（可以把某个IP的请求hash到同一台服务器也可以）
  - 直播场景下某个热门主播，如何保证他的礼物排行榜实时刷新
  - 直播场景下如何保证不被别人疯狂刷弹幕
  - 一系列直播，有的忘了
- 分布式限流实现
  - 单位时间段内请求数量
  - 讲讲具体实现，AOP方式无侵入实现
- synchronize和reentlock
  - 不多说，常规问法
- mysql
  - 索引原理b+tree
  - 事务隔离级别，可重复读事务下解决不可重复读的原理
  - 优化经验
  - 组合索引的原理
  - where a=1 and b=2，走的哪个索引，为啥
- es
  - es为什么查询快
  - 倒排索引
- kafka
  - kafka消息消费的原理
- java参数校验如何设计
  - 注解实现
