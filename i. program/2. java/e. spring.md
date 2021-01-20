# spring

- [spring](#spring)
  - [一、spring内部构成](#一spring内部构成)
  - [二、spring加载机制](#二spring加载机制)
  - [三、spring特性](#三spring特性)
    - [spring的bean生命周期](#spring的bean生命周期)
    - [spring之IOC](#spring之ioc)
    - [spring之AOP](#spring之aop)
  - [四、spring中涉及的设计模式](#四spring中涉及的设计模式)
    - [4.1、创建型设计模式](#41创建型设计模式)
      - [简单工厂](#简单工厂)
      - [工厂模式](#工厂模式)
      - [单例模式](#单例模式)
    - [4.2、 结构型设计模式](#42-结构型设计模式)
      - [适配器模式](#适配器模式)
      - [装饰模式](#装饰模式)
      - [代理模式](#代理模式)
    - [4.3、行为型设计模式](#43行为型设计模式)
      - [观察者模式](#观察者模式)
      - [策略模式](#策略模式)
      - [模板模式](#模板模式)
  - [五、spring定时器](#五spring定时器)
    - [condition1:单个任务单独执行](#condition1单个任务单独执行)
    - [condition2:单个任务通过子线程运行](#condition2单个任务通过子线程运行)
    - [3:一个任务并发执行](#3一个任务并发执行)

## 一、spring内部构成

Spring核心组件只有Core、Context、Beans三个。core包侧重于帮助类，操作工具，beans包更侧重于bean实例的描述。context更侧重全局控制，功能衍生。

## 二、spring加载机制

首先思考一个main方法中如何启动Spring：
```java
ApplicationContext ctx = new XmlApplicationContext("app.xml");
```
那么Web容器中如何启动Spring呢？

ApplicationContext ctx = new XmlApplicationContext("app.xml");
方法1: 利用Spring自带的Servlet启动，配好Servlet，加载Servlet的时候，就初始化了WebApplicationContext
方法2: 利用Spring自带的Listener启动，配好Listener，加载Listener的时候，就初始化WebApplicationContext

![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20191218175151.png)

容器初始化流程图：
![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20191218174250.png)


## 三、spring特性

### spring的bean生命周期

> [Spring Bean的生命周期](https://www.jianshu.com/p/1dec08d290c1)

Spring什么时候实例化bean，首先要分2种情况：

- 第一：如果你使用BeanFactory作为Spring Bean的工厂类，则所有的bean都是在第一次使用该Bean的时候实例化 
- 第二：如果你使用ApplicationContext作为Spring Bean的工厂类，则又分为以下几种情况： 
  - （1）：如果bean的scope是singleton的，并且lazy-init为false（默认是false，所以可以不用设置），则ApplicationContext启动的时候就实例化该Bean，并且将实例化的Bean放在一个map结构的缓存中，下次再使用该Bean的时候，直接从这个缓存中取 
  - （2）：如果bean的scope是singleton的，并且lazy-init为true，则该Bean的实例化是在第一次使用该Bean的时候进行实例化 
  - （3）：如果bean的scope是prototype的，则该Bean的实例化是在第一次使用该Bean的时候进行实例化 

Spring什么时候销毁bean：


1. 实例化 Instantiation
2. 属性赋值 Populate
3. 初始化 Initialization
4. 销毁 Destruction （是）

```java
protected Object doCreateBean(final String beanName, final RootBeanDefinition mbd, final @Nullable Object[] args) 
    throws BeanCreationException {

    // Instantiate the bean.
    BeanWrapper instanceWrapper = null;
    if (instanceWrapper == null) {
       // 实例化阶段！
      instanceWrapper = createBeanInstance(beanName, mbd, args);
    }

    // Initialize the bean instance.
    Object exposedObject = bean;
    try {
       // 属性赋值阶段！
      populateBean(beanName, mbd, instanceWrapper);
       // 初始化阶段！
      exposedObject = initializeBean(beanName, exposedObject, mbd);
    }
}
```
所有的bean被解析为BeanDefinition对象，并存入BeanDefinitionMap中，由BeanFactory(也就是IOC容器)来进行管理。
__在Spring代码中，BeanFactory只是个接口，并不是IOC容器的具体实现，但是Spring容器给出了很多种实现，如 DefaultListableBeanFactory、XmlBeanFactory、ApplicationContext等，都是附加了某种功能的实现。__




### spring之IOC

IOC分两个过程：bean的解析注册 和 bean的实例化。


### spring之AOP

面向切面编程：




## 四、spring中涉及的设计模式

> [spring中涉及的设计模式](https://blog.csdn.net/caoxiaohong1005/article/details/80039656)

### 4.1、创建型设计模式

#### 简单工厂

- __实现方式__：BeanFactory。 Spring中的BeanFactory就是简单工厂模式的体现，根据传入一个唯一的标识来获得Bean对象，但是否是在传入参数后创建还是传入参数前创建这个要根据具体情况来定。
- __实质__：由一个工厂类根据传入的参数，动态决定应该创建哪一个产品类。
- __实现原理__：
- __优点__：

#### 工厂模式

- __实现方式__：
- __实质__：
- __实现原理__：
- __优点__：

#### 单例模式

- __实现方式__：spring中bean就采取的单例模式
- __实质__：
- __实现原理__：
- __优点__：spring中的单例模式提供了全局的访问点BeanFactory。但没有从构造器级别去控制单例，这是因为spring管理的是任意的java对象。


### 4.2、 结构型设计模式

#### 适配器模式

- __实现方式__：
- __实质__：
- __实现原理__：
- __优点__：

#### 装饰模式

- __实现方式__：
- __实质__：
- __实现原理__：
- __优点__：
  - 动态地给一个对象添加一些额外的职责。
  - 就增加功能来说，Decorator模式相比生成子类更为灵活。

#### 代理模式

- __实现方式__：AOP底层，就是动态代理模式的实现
动态代理：在内存中构建的，不需要手动编写代理类
静态代理：需要手工编写代理类，代理类引用被代理对象
- __实质__：
- __实现原理__：切面在应用运行的时刻被织入。一般情况下，在织入切面时，AOP容器会为目标对象创建动态的创建一个代理对象。SpringAOP就是以这种方式织入切面的。 

  把切面应用到目标对象并创建新的代理对象的过程
- __优点__：


### 4.3、行为型设计模式

#### 观察者模式

- __实现方式__：spring的事件驱动模型使用的是 观察者模式 ，Spring中Observer模式常用的地方是listener的实现
- __实质__：
- __实现原理__：
- __优点__：

#### 策略模式

- __实现方式__：
- __实质__：
- __实现原理__：
- __优点__：

#### 模板模式

- __实现方式__：JdbcTemplate
- __实质__：
- __实现原理__：
- __优点__：


## 五、spring定时器

> [Schedule的源码解析](https://blog.csdn.net/weixin_40318210/article/details/78149692)


>说明：  
    1、使用@Scheduled时要在启动类上增加@EnableScheduling，使用@Async注解时要在启动类上加上@EnableAsync  
    2、下面所说的并行，是指两个不同的调度任务同时执行，并发是指同一个调度任务同时执行（即上次任务还没有执行完，下次任务已经开始执行了）


### condition1:单个任务单独执行

代码片段如下：
```java
@Scheduled(cron = "0/5 * * * * ? ")//每5秒执行一次
public void test1() {
    System.out.println(Thread.currentThread().getName() + ":test1.start" + new Date());
    try {
        Thread.sleep(7000);
    } catch (Exception e) {

    }
    System.out.println(Thread.currentThread().getName() + ":test1.end " + new Date());
}
```

cron表达式为每5秒执行一次，任务中sleep7秒，最终输入为每10秒执行一次。

__结论：上次任务没有执行完，下次任务不会执行，一直等到该次任务执行完且到任务执行时间点，下一次任务才会执行。__

### condition2:单个任务通过子线程运行

代码片段如下：

```java
@Scheduled(cron = "0/5 * * * * ? ")//每5秒执行一次
public void test1() {
    System.out.println("主线程" + Thread.currentThread().getName());
    new Thread(() -> {
        System.out.println("子线程" + Thread.currentThread().getName() + ":test1.start" + new Date());
        try {
            Thread.sleep(7000);
        } catch (Exception e) {

        }
        System.out.println("子线程" + Thread.currentThread().getName() + ":test1.end " + new Date());
    }).start();
}
```

__结论：通过子线程去执行任务，并不会阻塞主线程的任务调度。（按照cron表达式，依旧是5秒执行一次，而且每次都是新开一个线程。）__


### 3:一个任务并发执行

```java
@Scheduled(cron = "0/5 * * * * ? ")//每5秒执行一次
@Async
public void test1() {
    System.out.println(Thread.currentThread().getName() + ":test1.start" + new Date());
    try {
        Thread.sleep(7000);
    } catch (Exception e) {

    }
    System.out.println(Thread.currentThread().getName() + ":test1.end " + new Date());
}
```
__结论：任务异步执行时，且每次执行任务时，都是新建一个线程。__