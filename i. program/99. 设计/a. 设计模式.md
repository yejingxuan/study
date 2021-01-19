- [一、六种创建型设计模式](#一六种创建型设计模式)
  - [1、单例模式](#1单例模式)
  - [2、简单工厂](#2简单工厂)
  - [3、工厂模式](#3工厂模式)
  - [4、抽象工厂](#4抽象工厂)
  - [5、原型模式（Prototype）](#5原型模式prototype)
  - [6、建造者模式](#6建造者模式)
- [二、七种结构型设计模式](#二七种结构型设计模式)
  - [1、适配器模式(Adapter)](#1适配器模式adapter)
  - [2、装饰模式(Decorator)](#2装饰模式decorator)
  - [3、代理模式(Proxy)](#3代理模式proxy)
  - [4、外观模式(Facade)](#4外观模式facade)
  - [5、桥接模式(Bridge)](#5桥接模式bridge)
  - [6、组合模式](#6组合模式)
  - [7、享元模式](#7享元模式)
- [三、十一种行为型设计模式](#三十一种行为型设计模式)
  - [1、策略模式(strategy)](#1策略模式strategy)
  - [2、模板方法模式(Template)](#2模板方法模式template)
  - [3、观察者模式(Template)](#3观察者模式template)
  - [4、迭代器模式(Iterator)](#4迭代器模式iterator)
  - [5、责任链模式(Chain of Responsibility)](#5责任链模式chain-of-responsibility)
  - [6、命令模式(Command)](#6命令模式command)
  - [7、备忘录模式(Memento)](#7备忘录模式memento)
  - [8、状态模式(State)](#8状态模式state)
- [参考文章](#参考文章)

## 一、六种创建型设计模式

### 1、单例模式

__完美单例：__

```java
public class SingletonTest {
 
    private static SingletonTest instance = null;
 
    private SingletonTest() {
    }
 
    private static synchronized void syncInit() {
        if (instance == null) {
            instance = new SingletonTest();
        }
    }
 
    public static SingletonTest getInstance() {
        if (instance == null) {
            syncInit();
        }
        return instance;
    }
}
```
完美单例并不完美，只想理想状态下的完美。
- 1、反射可以跨过private权限，获取到构造函数，来生成新的对象——解决方案：在构造函数里再加上一层判断
    ```java
    private SingletonTest() {
        if (instance != null) {
            return instance;
        }
    }
    ```
- 2、序列化可以通过流的方式生成新的对象——解决方案：
重写readUnshared()方法，反序列化时会调用该方法，返回对象引用即可
    ```java
    private  Object readResolve(){
        return instance;
    }
    ```

### 2、简单工厂




### 3、工厂模式




### 4、抽象工厂




### 5、原型模式（Prototype）

>原型模式虽然是创建型的模式，但是与工程模式没有关系，从名字即可看出，该模式的思想就是将一个对象作为原型，对其进行复制、克隆，产生一个和原对象类似的新对象


```java

public class Prototype implements Cloneable, Serializable {
 
    private static final long serialVersionUID = 1L;
    private String string;
 
    private SerializableObject obj;
 
    /* 浅复制 */
    public Object clone() throws CloneNotSupportedException {
        Prototype proto = (Prototype) super.clone();
        return proto;
    }
 
    /* 深复制 */
    public Object deepClone() throws IOException, ClassNotFoundException {
 
        /* 写入当前对象的二进制流 */
        ByteArrayOutputStream bos = new ByteArrayOutputStream();
        ObjectOutputStream oos = new ObjectOutputStream(bos);
        oos.writeObject(this);
 
        /* 读出二进制流产生的新对象 */
        ByteArrayInputStream bis = new ByteArrayInputStream(bos.toByteArray());
        ObjectInputStream ois = new ObjectInputStream(bis);
        return ois.readObject();
    }
 
    public String getString() {
        return string;
    }
 
    public void setString(String string) {
        this.string = string;
    }
 
    public SerializableObject getObj() {
        return obj;
    }
 
    public void setObj(SerializableObject obj) {
        this.obj = obj;
    }
 
}

```






### 6、建造者模式


> 工厂类模式提供的是创建单个类的模式，而建造者模式则是将各种产品集中起来进行管理，用来创建复合对象，所谓复合对象就是指某个类具有不同的属性，

建造者模式将很多功能集成到一个类里，这个类可以创造出比较复杂的东西。

所以与工程模式的区别就是：工厂模式关注的是创建单个产品，而建造者模式则关注创建符合对象，多个部分。因此，是选择工厂模式还是建造者模式，依实际情况而定。

```java
public class Builder {
    
    private List<Sender> list = new ArrayList<Sender>();
    
    public void produceMailSender(int count){
        for(int i=0; i<count; i++){
            list.add(new MailSender());
        }
    }
    
    public void produceSmsSender(int count){
        for(int i=0; i<count; i++){
            list.add(new SmsSender());
        }
    }
}
```





## 二、七种结构型设计模式

### 1、适配器模式(Adapter)

>适配器模式将某个类的接口转换成客户端期望的另一个接口表示，目的是消除由于接口不匹配所造成的类的兼容性问题。主要分为三类：类的适配器模式、对象的适配器模式、接口的适配器模式


- 类的适配器


- 对象适配器


- 接口适配器





### 2、装饰模式(Decorator)

>装饰模式就是给一个对象增加一些新的功能，而且是动态的，要求装饰对象和被装饰对象实现同一个接口，装饰对象持有被装饰对象的实例



### 3、代理模式(Proxy)

>代理模式就是多一个代理类出来，替原对象进行一些操作，比如我们在租房子的时候回去找中介，为什么呢？因为你对该地区房屋的信息掌握的不够全面，希望找一个更熟悉的人去帮你做，此处的代理就是这个意思。再如我们有的时候打官司，我们需要请律师，因为律师在法律方面有专长，可以替我们进行操作，表达我们的想法



- __代理模式和装饰模式的区别__

  PS：两个设计模式看起来很像。对装饰器模式来说，装饰者（decorator）和被装饰者（decoratee）都实现同一个 接口。对代理模式来说，代理类（proxy class）和真实处理的类（real class）都实现同一个接口。此外，不论我们使用哪一个模式，都可以很容易地在真实对象的方法前面或者后面加上自定义的方法。

  然而，实际上，在装饰器模式和代理模式之间还是有很多差别的。装饰器模式关注于在一个对象上动态的添加方法，然而代理模式关注于控制对对象的访问。换句话 说，用代理模式，代理类（proxy class）可以对它的客户隐藏一个对象的具体信息。因此，当使用代理模式的时候，我们常常在一个代理类中创建一个对象的实例。并且，当我们使用装饰器模 式的时候，我们通常的做法是将原始对象作为一个参数传给装饰者的构造器。

  我们可以用另外一句话来总结这些差别：使用代理模式，代理和真实对象之间的的关系通常在编译时就已经确定了，而装饰者能够在运行时递归地被构造。



### 4、外观模式(Facade)

> 为子系统中的一组接口提供一个统一的入口。外观模式定义了一个高层接口，这个接口使得这一子系统更加容易使用。外观模式就是将他们的关系放在一个Facade类中，降低了类类之间的耦合度，该模式中没有涉及到接口

外观模式又称为门面模式，它是一种对象结构型模式。__外观模式是迪米特法则的一种具体实现__，通过引入一个新的外观角色可以降低原有系统的复杂度，同时降低客户类与子系统的耦合度。


```java
public class CPU {    
    public void startup(){
        System.out.println("cpu startup!");
    }
}

public class Disk {    
    public void startup(){
        System.out.println("disk startup!");
    }
}

public class Computer {
    private CPU cpu;
    private Disk disk;
    
    public Computer(){
        cpu = new CPU();
        disk = new Disk();
    }
    
    public void startup(){
        System.out.println("start the computer!");
        cpu.startup();
        disk.startup();
        System.out.println("start computer finished!");
    }
}

```



### 5、桥接模式(Bridge)

>桥接模式就是把事物和其具体实现分开，使他们可以各自独立的变化。桥接的用意是：将抽象化与实现化解耦，使得二者可以独立变化，像我们常用的JDBC桥DriverManager一样，JDBC进行连接数据库的时候，在各个数据库之间进行切换，基本不需要动太多的代码，甚至丝毫不用动，原因就是JDBC提供统一接口，每个数据库提供各自的实现，用一个叫做数据库驱动的程序来桥接就行了

```java
public interface Sourceable {
    public void method();
}

public class SourceSub1 implements Sourceable {
    @Override
    public void method() {
        System.out.println("this is the first sub!");
    }
}

public class SourceSub2 implements Sourceable {
 
    @Override
    public void method() {
        System.out.println("this is the second sub!");
    }
}

public abstract class Bridge {
    private Sourceable source;
 
    public void method(){
        source.method();
    }
    
    public Sourceable getSource() {
        return source;
    }
 
    public void setSource(Sourceable source) {
        this.source = source;
    }
}

public class MyBridge extends Bridge {
    public void method(){
        getSource().method();
    }
}

public class BridgeTest {
    
    public static void main(String[] args) {
        
        Bridge bridge = new MyBridge();
        
        /*调用第一个对象*/
        Sourceable source1 = new SourceSub1();
        bridge.setSource(source1);
        bridge.method();
        
        /*调用第二个对象*/
        Sourceable source2 = new SourceSub2();
        bridge.setSource(source2);
        bridge.method();
    }
}

```


### 6、组合模式

> 组合模式有时又叫部分-整体模式在处理类似树形结构的问题时比较方便



### 7、享元模式

> 享元模式的主要目的是实现对象的共享，即共享池，当系统中对象多的时候可以减少内存的开销，通常与工厂模式一起使用。

* 建立一个池的概念，实现对象的共享，来减少大量创建对象的内存开销



## 三、十一种行为型设计模式

- 十一种设计模式关系图：
  - 第一类：通过父类与子类的关系进行实现。
  - 第二类：两个类之间。
  - 第三类：类的状态。
  - 第四类：通过中间类

  ![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20191219105430.png)

------

### 1、策略模式(strategy)

> 策略模式定义了一系列算法，并将每个算法封装起来，使他们可以相互替换，且算法的变化不会影响到使用算法的客户。需要设计一个接口，为一系列实现类提供统一的方法，多个实现类实现该接口，设计一个抽象类（可有可无，属于辅助类），提供辅助函数

> 定义一系列算法类，将每一个算法封装起来，并让它们可以相互替换，策略模式让算法独立于使用它的客户而变化，也称为政策模式(Policy)。策略模式是一种对象行为型模式。

- _策略模式的决定权在用户，系统本身提供不同算法的实现，新增或者删除算法，对各种算法做封装。因此，策略模式多用在算法决策系统中，外部用户只需要决定用哪个算法即可。_

- _PS:工厂模式和策略模式的侧重点不同， 工厂模式侧重于生成某个基类的多个子类， 策略模式侧重于客户类以何种方式使用使用多种子类完成自己需要实现的功能。_


### 2、模板方法模式(Template)

> 解释一下模板方法模式，就是指：一个抽象类中，有一个主方法，再定义1...n个方法，可以是抽象的，也可以是实际的方法，定义一个类，继承该抽象类，重写抽象方法，通过调用抽象类，实现对子类的调用。


----




___包括这个模式在内的接下来的四个模式，都是类和类之间的关系，不涉及到继承，___

### 3、观察者模式(Template)

> 当一个对象变化时，其它依赖该对象的对象都会收到通知，并且随着变化！对象之间是一种一对多的关系。



### 4、迭代器模式(Iterator)

> 迭代器模式就是顺序访问聚集中的对象，一般来说，集合中非常常见，如果对集合类比较熟悉的话，理解本模式会十分轻松。这句话包含两层意思：一是需要遍历的对象，即聚集对象，二是迭代器对象，用于对聚集对象进行遍历访问。


### 5、责任链模式(Chain of Responsibility)

> 责任链模式，有多个对象，每个对象持有对下一个对象的引用，这样就会形成一条链，请求在这条链上传递，直到某一对象决定处理该请求。但是发出者并不清楚到底最终那个对象会处理该请求，所以，责任链模式可以实现，在隐瞒客户端的情况下，对系统进行动态的调整。


### 6、命令模式(Command)

> 命令模式很好理解，举个例子，司令员下令让士兵去干件事情，从整个事情的角度来考虑，司令员的作用是，发出口令，口令经过传递，传到了士兵耳朵里，士兵去执行。这个过程好在，三者相互解耦，任何一方都不用去依赖其他人，只需要做好自己的事儿就行，司令员要的是结果，不会去关注到底士兵是怎么实现的。


-----

### 7、备忘录模式(Memento)

> 主要目的是保存一个对象的某个状态，以便在适当的时候恢复对象，个人觉得叫备份模式更形象些，通俗的讲下：假设有原始类A，A中有各种属性，A可以决定需要备份的属性，备忘录类B是用来存储A的一些内部状态，类C呢，就是一个用来存储备忘录的，且只能存储，不能修改等操作。做个图来分析一下：



### 8、状态模式(State)

> 当对象的状态改变时，同时改变其行为，很好理解！就拿QQ来说，有几种状态，在线、隐身、忙碌等，每个状态对应不同的操作，而且你的好友也能看到你的状态，所以，状态模式就两点：1、可以通过改变状态来获得不同的行为。2、你的好友能同时看到你的变化。





## 参考文章
- https://blog.csdn.net/zhangerqing/article/details/8194653
- https://blog.csdn.net/zhangerqing/article/details/8239539 
- https://blog.csdn.net/lovelion/article/details/17517213?tdsourcetag=s_pctim_aiomsg
- https://blog.csdn.net/zhangerqing/article/details/8243942  
- https://blog.csdn.net/zhangerqing/article/details/8245537