# java基础进阶

- [java基础进阶](#java基础进阶)
  - [一、java集合](#一java集合)
    - [1.1、List](#11list)
    - [1.2、Set](#12set)
    - [1.3、Map](#13map)
    - [1.4、List-set-Map比较](#14list-set-map比较)
  - [二、java类](#二java类)
  - [三、java关键字](#三java关键字)
  - [参考文章](#参考文章)


## 一、java集合

### 1.1、List
- __吐血汇总__

  |       |ArrayList|LinkedList|Vector（过时）|CopyOnWriteArrayList|
  |:-----:|:-------:|:--------:|:----:|:----:|
  |实现原理|Object动态数组|双向链表|Object动态数组(同ArrayList)|Object动态数组|
  |线程安全|否|否|是，所有方法加synchronize锁|是，reentlock+CopyOnWrite机制|
  |插入速度|慢|快|慢|慢|
  |查询速度|快 O(1)|慢 O(n)|快 O(1)|快 O(1)

- __ArrayList__
  - 实现原理
    - ArrayList是List接口的实现类，插入的元素是有序的，其顺序就是插入时的先后顺序，ArrayList底层实现为动态数组，默认大小为10，每次添加元素时会检查容量是否超过出示容量，如果超出则进行扩容（int newCapacity = (oldCapacity * 3)/2 + 1;）
  - 扩容机制
    - 每次添加元素时会检查容量是否超过出示容量，超出容量后进行扩容
    - 把原来的数组复制到另一个是自己1.5倍的数组，把新元素添加到扩容以后的数组中

- __LinkedList__
  - 实现原理
    - LinkedList是一个继承于AbstractSequentialList的双向链表。它也可以被当做堆栈、队列或双端队列进行使用，出LinkedList是一个无界链表，不存在容量不足的问题
  - 特性
    - 双向链表，在首部和尾部添加、删除元素效率高效，在中间添加、删除元素效率较低  
  
- __Vector__
  - 实现原理
    - 实现原理几乎同ArrayList，只是每个方法加了synchronize锁，效率低。
  - 扩容机制
    - Vector可以指定增长因子，如果该增长因子指定了，那么扩容的时候会每次新的数组大小会在原数组的大小基础上加上增长因子；如果不指定增长因子，那么就给原数组大小*2

- __synchronizedList__
  - 实现原理
    - 同ArrayList
  - 特性
    - 无论是读取还是写入，它都会进行加锁，当我们并发级别特别高，线程之间在任何操作上都会进行等待，因此在某些场景中它不是最好的选择

- __CopyOnWriteArrayList__
  - 实现原理
    - 同ArrayList
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
    - 底层基于数组+链表来实现，链表元素超过8个后转为红黑树
  - 特性

  - 扩容原理

- __LinkedHashMap__
  - 实现原理
    - 底层基于HashMap来实现，只不过，在每个元素中添加了头尾指针，指向前一个和后一个元素，从而构造一条双向链表，实现有序性

- __ConcurrentHashMap__
  - 实现原理
    - 基本同HashMap，将HashMap差分为多个
  - 特性
    - 采用分段锁来保证线程安全，默认采用八个HashMap分段，支持8个线程并发

- __TreeMap__
  - 原理
    - TreeMap底层是红黑树，能够实现该Map集合有序
  - 特性
    - 若比较器为空则key一定不能为null，若比较器不为空则key可以为null由TreeMap其比较器而定
    - 排序特性：put方法里要么根据treemap自身比较器，若其比较器为null，就用key的compareTo方法


### 1.4、List-set-Map比较



## 二、java类







## 三、java关键字

















## 参考文章

- [并发容器(二)—线程安全的List](https://blog.csdn.net/p_programmer/article/details/86027076)