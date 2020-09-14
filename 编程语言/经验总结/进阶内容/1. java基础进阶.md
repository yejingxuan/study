# java基础进阶

- [java基础进阶](#java基础进阶)
  - [一、java数据结构](#一java数据结构)
    - [1.1、List](#11list)
    - [1.2、Set](#12set)
    - [1.3、Map](#13map)
    - [1.4、红黑树](#14红黑树)
    - [1.5、List-set-Map比较](#15list-set-map比较)
    - [1.6、tree](#16tree)
  - [二、java类](#二java类)
    - [2.1、java内部类](#21java内部类)
    - [2.2、抽象类和接口的区别](#22抽象类和接口的区别)
    - [2.3、为什么java不能多继承](#23为什么java不能多继承)
  - [三、java知识](#三java知识)
    - [3.1、深拷贝 && 浅拷贝 && 延迟拷贝](#31深拷贝--浅拷贝--延迟拷贝)
    - [3.2、序列化](#32序列化)
    - [3.3、java是值传递还是引用传递？](#33java是值传递还是引用传递)
    - [3.4、i++问题](#34i问题)
  - [四、spring](#四spring)
    - [4.1、IOC](#41ioc)
    - [4.2、AOP](#42aop)
    - [4.3、spring bean容器](#43spring-bean容器)
    - [4.4、spring依赖注入方式](#44spring依赖注入方式)
    - [4.5、动态代理](#45动态代理)
    - [4.6、spring事务](#46spring事务)
    - [4.7、springMVC](#47springmvc)
  - [参考文章](#参考文章)


## 一、java数据结构

### 1.1、List
- __吐血汇总__

  |       |ArrayList|LinkedList|Vector（过时）|CopyOnWriteArrayList|
  |:-----:|:-------:|:--------:|:----:|:----:|
  |实现原理|Object动态数组|双向链表|Object动态数组(同ArrayList)|Object动态数组|
  |线程安全|否|否|是，所有方法加synchronize锁|是，reentlock+CopyOnWrite机制|
  |插入速度|慢|快|慢|慢|
  |查询速度|快 O(1)|慢 O(n)|快 O(1)|快 O(1)|

- __ArrayList__
  - 实现原理
    - ArrayList是List接口的实现类，插入的元素是有序的，其顺序就是插入时的先后顺序，ArrayList底层实现为动态数组，每次添加元素时会检查容量是否超过出示容量，如果超出则进行扩容（int newCapacity = (oldCapacity * 3)/2 + 1;）
  - 增删机制
    - 删除元素需要调用 System.arraycopy() 将 index+1 后面的元素都复制到 index 位置上，该操作的时间复杂度为 O(N)，ArrayList 删除元素的代价是非常高的。
  - 扩容机制
    - arraylist的默认初始容量为0，当第一次添加元素时，扩容到10
    - 然后后面每次添加元素时会检查容量是否超过出示容量，超出容量后进行扩容
    - 把原来的数组复制到另一个是自己1.5倍的数组，把新元素添加到扩容以后的数组中

- __Vector__
  - 实现原理
    - 实现原理几乎同ArrayList，只是每个方法加了synchronize锁，效率低。
  - 增删机制
    - 同ArrayList
  - 扩容机制
    - Vector可以指定增长因子，如果该增长因子指定了，那么扩容的时候会每次新的数组大小会在原数组的大小基础上加上增长因子；如果不指定增长因子，那么就给原数组大小*2

- __LinkedList__
  - 实现原理
    - LinkedList是一个继承于AbstractSequentialList的双向链表。它也可以被当做堆栈、队列或双端队列进行使用，出LinkedList是一个无界链表，不存在容量不足的问题
    - 使用 Node 存储链表节点信息，Node包含next和prev指针，每个链表又存储了 first 和 last 指针
  - 特性
    - 双向链表，在首部和尾部添加、删除元素效率高效，在中间添加、删除元素效率较低。


- __synchronizedList__
  - 实现原理
    - 同ArrayList
  - 特性
    - 无论是读取还是写入，它都会进行加锁，当我们并发级别特别高，线程之间在任何操作上都会进行等待，因此在某些场景中它不是最好的选择

- __CopyOnWriteArrayList__
  - 实现原理
    - 同ArrayList，基于动态数组来实现，只是写操作在一个复制的数组上进行，读操作还是在原始数组中进行，读写分离，互不影响。
  - 特性
    - 读读之间不互斥 并且 读写也不互斥
    - add操作中使用了重入锁，但是此锁只针对写-写操作
    - 读写之间不用互斥，关键就在于添加值的操作并不是直接在原有数组中完成，而是使用原有数组复制一个新的数组，然后将值插入到新的数组中，最后使用新数组替换旧数组，这样插入就完成了。在add的过程中旧数组没有得到修改，因此写入操作不影响读取操，另外，数组定义private transient volatile Object[] array，其中采用volatile修饰，保证内存可见性

### 1.2、Set

- __吐血整理__
  |       |HashSet|LinkedHashSet|TreeSet|CopyOnWriteArraySet|
  |:-----:|:-------:|:--------:|:----:|:----:|
  |实现原理|基于HashMap|基于LinkedHashMap|二叉树|CopyOnWriteArrayList数据结构|
  |线程安全|否|否|否|是，reentlock+CopyOnWrite机制|
  |可否存储空值|允许|允许|不允许|是|
  |插入速度|快|快|快|慢|
  |查询速度|快 O()|快 O()|慢 O()|快|
  |是否有序|无|是|是|是|

- __HashSet__
  - 实现原理
    - 底层基于HashMap来实现的。HashSet的元素实际上是存储在底层HashMap的key上的
  - 特性
    - HashSet不存入重复元素的规则.使用hashcode和equals


- __LinkedHashSet__
  - 实现原理
    - 它的底层是一个LinkedHashMap，元素的所有操作都是由LinkedHashMap来维护，内部由双向链表来记录顺序，保证有序性

- __TreeSet__
  - 实现原理
    - TreeSet的底层是通过TreeMap实现的。TreeSet并不是根据插入的顺序来排序，而是根据实际的值的大小来排序

- __CopyOnWriteArraySet__



### 1.3、Map

- __吐血整理__
  |       |HashMap|LinkedHashMap|ConcurrentHashMap|TreeMap|
  |:-----:|:-------:|:--------:|:----:|:----:|
  |实现原理|数组+链表/红黑树|数组+链表|数组+链表/红黑树|红黑树
  |是否有null的key|是|是|是|视比较器来判断，默认不为空|
  |是否有序的key|否|是|否|是|
  |线程安全|否|是|是|否|
  |插入速度|快|慢|稍快|稍快|


- __HashMap__
  - 实现原理
    - 底层基于数组+链表来实现，链表元素超过8个后转为红黑树，红黑树的时间复杂度O（logn）
  - 特性
    - Hash算法优化：hashcode计算后右移16位，并和原始hashcode值进行异或运算，减少Hash碰撞
    - 寻址算法优化：(n-1) & hash  ==== hash取模的效果和 hash&(n-1)结果是一样的，但是性能高一些。
    - HashMap 允许插入键为 null 的键值对。但是因为无法调用 null 的 hashCode() 方法，也就无法确定该键值对的桶下标，只能通过强制指定一个桶下标来存放。HashMap 使用第 0 个桶存放键为 null 的键值对。
  - 扩容原理
    - 默认情况下HashMap的容量是16，如果通过构造函数指定了一个数字作为容量，那么会选择大于该数字的第一个2的幂作为容量。(3->4、7->8、9->16)。设置16是因为是2的幂，符合内部计算的机制，而且这个值，不大也不小，太小了就有可能频繁发生扩容，影响效率。太大了又浪费空间,这一特点也能够极大降低重新计算桶下标操作的复杂度
    - 当一个map填满了75%的bucket（数组）时候，和其它集合类(如ArrayList等)一样，将会创建原来HashMap大小的两倍的bucket数组，来重新调整map的大小，并将原来的对象放入新的bucket（数组）中
    - 而加载因子0.75的是为了提高空间利用率和减少查询成本的折中，0.75的话碰撞最小
    - Resize步骤
       - 扩容：创建一个新的Entry空数组，长度是原数组的2倍。
       - ReHash：遍历原Entry数组，把所有的Entry重新Hash到新数组。为什么要重新Hash呢？因为长度扩大以后，Hash的规则也随之改变。
       - hash公式：index = HashCode（Key） & （Length - 1）

- __LinkedHashMap__
  - 实现原理
    - 底层基于HashMap来实现，只不过，在每个元素中添加了头尾指针，指向前一个和后一个元素，从而构造一条双向链表，实现有序性

- __ConcurrentHashMap__
  - 实现原理
    - 基本同HashMap，将HashMap差分为多个
  - 特性
    - 1.8之前采用分段锁来保证线程安全，默认采用16个HashMap分段，支持16个线程并发
    - 1.8之后不再使用分段的思想，还是使用一个数组，对数组中的每个数组位置进行CAS操作，对数组上的链表元素进行synchronize加锁处理

- __TreeMap__
  - 原理
    - TreeMap底层是红黑树，其容器内所有的K-V键值对对象都是这个红黑树上的一个节点, 能够实现该Map集合有序
  - 特性
    - 若比较器为空则key一定不能为null，若比较器不为空则key可以为null由TreeMap其比较器而定
    - 排序特性：put方法里要么根据treemap自身比较器，若其比较器为null，就用key的compareTo方法

- __WeakHashMap__ 
  - entry对象是用的弱引用，被 WeakReference 关联的对象在下一次垃圾回收时会被回收。通过使用 WeakHashMap 来引用缓存对象，由 JVM 对这部分缓存进行回收


### 1.4、红黑树

### 1.5、List-set-Map比较

### 1.6、tree
- 二叉树遍历
![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20200629173221.png)

## 二、java类

### 2.1、java内部类

- __内部类的作用__
  - 内部类方法可以访问该类定义所在作用域中的数据，包括被 private 修饰的私有数据
  - 内部类可以对同一包中的其他类隐藏起来
  - 内部类可以实现 java 单继承的缺陷
  - 当我们想要定义一个回调函数却不想写大量代码的时候我们可以选择使用匿名内部类来实现

- __java从语法上来讲把内部类分为了__
  1. 成员内部类
     - 成员内部类就像一个实例变量。它可以访问它的外部类的所有成员变量和方法，不管是静态的还是非静态的都可以。
     - 生成方式
        ```java
        // 创建成员内部类的对象
        // 需要先创建外部类的实例
        MemberInner.Inner2 inner = new MemberInner().new Inner2();
        ```
  2. 静态内部类
     - 只能使用 外部类的静态成员和静态方法
     - 生成静态内部类对象的方式为：
        ```java
        OuterClass.InnerClass inner = new OuterClass.InnerClass()
        ```
    
  3. 局部内部类
     - 只能访问方法中定义的final类型的局部变量。局部内部类在方法中定义，所以只能在方法中使用，即只能在方法当中生成局部内部类的实例并且调用其方法。
  
  4. 匿名内部类
     - 匿名内部类隐式地继承了一个父类或者实现了一个接口,匿名内部类使用得比较多，通常是作为一个方法参数
        ```java
        new Thread(new Runnable() {
            @Override
            public void run() {

            }
        }).run();
        ```


### 2.2、抽象类和接口的区别

- 抽象类只能继承，接口只能实现
- 接口只能做方法声明，抽象类中可以作方法声明，也可以做方法实现
- 接口中的所有成员变量 为public static final， 静态不可修改
- 接口强调的是功能，抽象类强调的是所属关系


### 2.3、为什么java不能多继承
- 防止多继承的时候子类调用父类方法的时候，出现多个父类都拥有该方法，不知道该调用哪个情况



## 三、java知识

### 3.1、深拷贝 && 浅拷贝 && 延迟拷贝
- __浅拷贝__
  - 浅拷贝是按位拷贝对象，它会创建一个新对象，这个对象有着原始对象属性值的一份精确拷贝。
  - 如果属性是基本类型（如String int等等），拷贝的就是基本类型的值；如果属性是内存地址（引用类型），拷贝的就是内存地址 ，因此如果其中一个对象改变了这个地址，就会影响到另一个对象。
  - __集合、数组的复制默认都是浅拷贝__

- __深拷贝__
  - 深拷贝会拷贝所有的属性,并拷贝属性指向的动态分配的内存。__当对象和它所引用的对象一起拷贝时即发生深拷贝__。深拷贝相比于浅拷贝速度较慢并且花销较大

- __延迟拷贝__
  - 当最开始拷贝一个对象时，会使用速度较快的浅拷贝，还会使用一个计数器来记录有多少对象共享这个数据。当程序想要修改原始的对象时，它会决定数据是否被共享（通过检查计数器）并根据需要进行深拷贝。
  - 读取数据时进行浅拷贝，修改数据时进行浅拷贝。



### 3.2、序列化

- 可以序列化是干什么的?它将整个对象图写入到一个持久化存储文件中并且当需要的时候把它读取回来, 这意味着当你需要把它读取回来时你需要整个对象图的一个拷贝。这就是当你深拷贝一个对象时真正需要的东西。请注意，当你通过序列化进行深拷贝时，必须确保对象图中所有类都是可序列化的。



### 3.3、java是值传递还是引用传递？

- __值传递（pass by value）__
  - 是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。

- __引用传递（pass by reference）__
  - 是指在调用函数时将实际参数的地址直接传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。



- 大部分的认知里，传递的参数如果是普通类型，那就是值传递，如果是对象，那就是引用传递。
  ```java
  public static void main(String[] args) {
    ParamTest pt = new ParamTest();

    User hollis = new User();
    hollis.setName("Hollis");
    hollis.setGender("Male");
    pt.pass(hollis);
    System.out.println("print in main , user is " + hollis);
  }

  public void pass(User user) {
    user = new User();
    user.setName("hollischuang");
    user.setGender("Male");
    System.out.println("print in pass , user is " + user);
  }

  //打印结果如下
  print in pass , user is User{name='hollischuang', gender='Male'}
  print in main , user is User{name='Hollis', gender='Male'}
  ```

- 所以上面的参数其实是值传递，只是浅拷贝的情况下把实参里对象引用的地址当做值传递给了形式参数。
- 这里是把实际参数的引用的地址复制了一份，传递给了形式参数。所以，上面的参数其实是值传递，把实参对象引用的地址当做值传递给了形式参数。


### 3.4、i++问题

- i++是一个复合操作，可分为三个阶段：
  1. 读值，从内存到寄存器
  2. +1，寄存器自增
  3. 写值，写回内存

- volatile只能保证可见性和顺序性，不能保证原子性，所以volatile对i++没卵用，可以用AtmoicInteger原子类，底层基于cas



## 四、spring

### 4.1、IOC


### 4.2、AOP


### 4.3、spring bean容器

- __bean的生命周期__
  - 创建->初始化->使用->销毁
  1. 实例化bean：spring启动时根据配置和反射、动态代理来创建bean,
  2. 设置对象属性（依赖注入）：去查找这个bean依赖了哪些bena,然后也创建出来，进行注入
  3. 处理Aware接口：即把容器注入给bean,如果bean已经实现了ApplicationContextAware接口，spring容器会调用我们的bean的setApplicationContext(ApplicationContext)方法，传入spring的上下文，把spring容器传给这个bean
  4. BeanPostProcessor：bean实例构建好了后，可以对bean进行一些自定义的处理，让Bean实现BeanPostProcessor接口
  5. InitializingBean和init-method：执行配置的初始化方法
  6. beanPostProcessor：接口会调用postProcessorAfterInitialization方法， 在bean初始化完了后执行，可以用于一些缓存技术。
  7. disposableBean：当bean不再需要时，会经过清理阶段，如果bean实现了该接口，会调用他的destroy方法
  8. destroy-method：如果配置了destroy-method可以在bean销毁后执行该方法。


- __bean的作用域，通过scope属性来实现__
  - singleton:单例模式
    - 应用场景
  - prototype:多例模式，为每次请求创建一个对象
    - 应用场景：
  - request：对每次网络请求创建一次对象
  - session:对每个会话创建一次

### 4.4、spring依赖注入方式



- @Autowired:自动装配

- 接口注入

- set方法输入

- 构造函数注入
  ```java
  public class ServiceA{
    private MyDao myDao;
    public ServiceA(MyDao myDao){
      this.myDao = myDao
    }
  }
  ```


### 4.5、动态代理

- JDK动态代理
  - springAOP使用的就是JDK动态代理， jdk动态代理是在你的类有接口的时候来使用

- cglib动态代理
  - 当要代理的类没有接口时，AOP会使用cglib来生成动态代理。对代理的类生成子类，动态生成字节码，覆盖掉一些方法，在方法里加入增强代码

### 4.6、spring事务

- __spring事务流程__
  - 通过@transaction注解，spring会使用AOP的思想，在方法执行之前开启事务，在执行完毕后，根据方法是否报错，来决定是回滚还是提交事务。

- __事务传播机制__
  - __propagation_required__：spring默认设置————如果当前没有事务，就创建一个事务；如果当前存在事务，就加入该事务。

  - __propagation_requires_new__：无论当前有没有事务，都会创建新事务

  - __propagation_nested__：嵌套事务，外层事务回滚，则内层事务也会回滚，如果内层事务回滚，则只回滚内层事务

  - propagation_supports：如果当前没有事务，就以非事务机制执行；如果当前存在事务，就加入该事务；

  - propagation_mandatory：如果当前没有事务，就抛出异常；如果当前存在事务，就加入该事务；

  - propagation_not_support：以非事务的方法执行，如果当前存在事务，就把当前事务挂起

  - propagation_never：以非事务执行，存在事务就抛出异常

### 4.7、springMVC
![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20200530212411.png)

![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20200530212641.png)


## 参考文章

- [并发容器(二)—线程安全的List](https://blog.csdn.net/p_programmer/article/details/86027076)
- [搞懂 JAVA 内部类](https://juejin.im/post/5a903ef96fb9a063435ef0c8)