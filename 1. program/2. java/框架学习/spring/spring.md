# spring

- [spring](#spring)
  - [一、spring内部构成](#%e4%b8%80spring%e5%86%85%e9%83%a8%e6%9e%84%e6%88%90)
  - [二、spring加载机制](#%e4%ba%8cspring%e5%8a%a0%e8%bd%bd%e6%9c%ba%e5%88%b6)
  - [三、spring特性](#%e4%b8%89spring%e7%89%b9%e6%80%a7)
    - [spring的bean生命周期](#spring%e7%9a%84bean%e7%94%9f%e5%91%bd%e5%91%a8%e6%9c%9f)
    - [spring之IOC](#spring%e4%b9%8bioc)
    - [spring之AOP](#spring%e4%b9%8baop)
  - [四、spring中涉及的设计模式](#%e5%9b%9bspring%e4%b8%ad%e6%b6%89%e5%8f%8a%e7%9a%84%e8%ae%be%e8%ae%a1%e6%a8%a1%e5%bc%8f)
    - [4.1、创建型设计模式](#41%e5%88%9b%e5%bb%ba%e5%9e%8b%e8%ae%be%e8%ae%a1%e6%a8%a1%e5%bc%8f)
      - [简单工厂](#%e7%ae%80%e5%8d%95%e5%b7%a5%e5%8e%82)
      - [工厂模式](#%e5%b7%a5%e5%8e%82%e6%a8%a1%e5%bc%8f)
      - [单例模式](#%e5%8d%95%e4%be%8b%e6%a8%a1%e5%bc%8f)
    - [4.2、 结构型设计模式](#42-%e7%bb%93%e6%9e%84%e5%9e%8b%e8%ae%be%e8%ae%a1%e6%a8%a1%e5%bc%8f)
      - [适配器模式](#%e9%80%82%e9%85%8d%e5%99%a8%e6%a8%a1%e5%bc%8f)
      - [装饰模式](#%e8%a3%85%e9%a5%b0%e6%a8%a1%e5%bc%8f)
      - [代理模式](#%e4%bb%a3%e7%90%86%e6%a8%a1%e5%bc%8f)
    - [4.3、行为型设计模式](#43%e8%a1%8c%e4%b8%ba%e5%9e%8b%e8%ae%be%e8%ae%a1%e6%a8%a1%e5%bc%8f)
      - [观察者模式](#%e8%a7%82%e5%af%9f%e8%80%85%e6%a8%a1%e5%bc%8f)
      - [策略模式](#%e7%ad%96%e7%95%a5%e6%a8%a1%e5%bc%8f)
      - [模板模式](#%e6%a8%a1%e6%9d%bf%e6%a8%a1%e5%bc%8f)

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