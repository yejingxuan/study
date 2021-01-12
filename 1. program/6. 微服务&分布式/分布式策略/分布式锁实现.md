# 分布式锁实现

- [分布式锁实现](#%e5%88%86%e5%b8%83%e5%bc%8f%e9%94%81%e5%ae%9e%e7%8e%b0)
  - [一、概述](#%e4%b8%80%e6%a6%82%e8%bf%b0)
    - [1.1、啥是分布式锁](#11%e5%95%a5%e6%98%af%e5%88%86%e5%b8%83%e5%bc%8f%e9%94%81)
    - [1.2、为什么要实现分布式锁](#12%e4%b8%ba%e4%bb%80%e4%b9%88%e8%a6%81%e5%ae%9e%e7%8e%b0%e5%88%86%e5%b8%83%e5%bc%8f%e9%94%81)
    - [1.3、应用场景](#13%e5%ba%94%e7%94%a8%e5%9c%ba%e6%99%af)
  - [二、实现原理](#%e4%ba%8c%e5%ae%9e%e7%8e%b0%e5%8e%9f%e7%90%86)
  - [三、实现方式](#%e4%b8%89%e5%ae%9e%e7%8e%b0%e6%96%b9%e5%bc%8f)
    - [3.1、基于redis的分布式锁实现](#31%e5%9f%ba%e4%ba%8eredis%e7%9a%84%e5%88%86%e5%b8%83%e5%bc%8f%e9%94%81%e5%ae%9e%e7%8e%b0)
      - [实现流程](#%e5%ae%9e%e7%8e%b0%e6%b5%81%e7%a8%8b)
      - [集成方案](#%e9%9b%86%e6%88%90%e6%96%b9%e6%a1%88)
      - [锁的类型](#%e9%94%81%e7%9a%84%e7%b1%bb%e5%9e%8b)
        - [readlock](#readlock)
      - [注意事项](#%e6%b3%a8%e6%84%8f%e4%ba%8b%e9%a1%b9)
    - [3.2、基于zookeeper的分布式锁实现](#32%e5%9f%ba%e4%ba%8ezookeeper%e7%9a%84%e5%88%86%e5%b8%83%e5%bc%8f%e9%94%81%e5%ae%9e%e7%8e%b0)
      - [实现流程](#%e5%ae%9e%e7%8e%b0%e6%b5%81%e7%a8%8b-1)
      - [集成方案](#%e9%9b%86%e6%88%90%e6%96%b9%e6%a1%88-1)
      - [锁的类型](#%e9%94%81%e7%9a%84%e7%b1%bb%e5%9e%8b-1)
      - [注意事项](#%e6%b3%a8%e6%84%8f%e4%ba%8b%e9%a1%b9-1)
    - [3.3、基于etcd的分布式锁实现](#33%e5%9f%ba%e4%ba%8eetcd%e7%9a%84%e5%88%86%e5%b8%83%e5%bc%8f%e9%94%81%e5%ae%9e%e7%8e%b0)


> 参考文章： [Redis 分布式锁的正确实现方式（ Java 版 ）](http://www.importnew.com/27477.html)

## 一、概述

### 1.1、啥是分布式锁


### 1.2、为什么要实现分布式锁

- 效率：使用分布式锁可以避免不同节点重复相同的工作，这些工作会浪费资源。比如用户付了钱之后有可能不同节点会发出多封短信。

- 正确性：加分布式锁同样可以避免破坏正确性的发生，如果两个节点在同一条数据上面操作，比如多个节点机器对同一个订单操作不同的流程有可能会导致该笔订单最后状态出现错误，造成损失。

- 解决分布式应用在高并发情况下出现的线程安全问题。

### 1.3、应用场景


## 二、实现原理

- 互斥性：和我们本地锁一样互斥性是最基本，但是分布式锁需要保证在不同节点的不同线程的互斥。

- 可重入性：同一个节点上的同一个线程如果获取了锁之后那么也可以再次获取这个锁,而不会发生死锁。

- 锁超时：和本地锁一样支持锁超时，防止死锁。

- 高性能，高可用：加锁和解锁需要高效，同时也需要保证高可用防止分布式锁失效，可以增加降级。

- 支持阻塞和非阻塞：和 ReentrantLock 一样支持 lock 和 trylock 以及 tryLock(long timeOut)。

- 支持公平锁和非公平锁(可选)：公平锁的意思是按照请求加锁的顺序获得锁，非公平锁就相反是无序的。这个一般来说实现的比较少。



|               |redis | zookeeper | etcd
|---            |----  |----       |-----
|互斥性         |√ 单线程  |√ 目录机制  |√ 目录机制
|可重入         |√ 在一个线程获取到锁之后，把当前主机信息和线程信息保存起来，下次再获取之前先检查自己是不是当前锁的拥有者。     |√ 把当前客户端的主机信息和线程信息直接写入到节点中   |√ 把当前客户端的主机信息和线程信息直接写入到节点中
|锁超时         |√ 可设置key值的timout|√  临时节点特性（一旦客户端获取到锁之后突然挂掉（Session连接断开），那么这个临时节点就会自动删除掉。其他客户端就可以再次获得锁。）     |√ lease机制：Etcd 可以为存储的 key-value 对设置租约
|高性能，高可用  |√ 基于内存，可实现集群部署|√ 基于内存,可支持集群部署  |√ 基于内存,可支持集群部署 
|阻塞，非阻塞    |√ 自旋等待锁 |√  在节点上绑定监听器    |√ Watch监听机制
|公平锁，非公平锁|√在线程获取锁之前先把所有等待的线程放入一个队列中，然后按先进先出原则获取锁     |√   创建顺序节点       |√ Revision机制：每个 key 带有一个 Revision 号，每进行一次事务加一，因此它是全局唯一的


## 三、实现方式

### 3.1、基于redis的分布式锁实现
http://developer.51cto.com/art/201812/588335.htm




实现原理：  
- redis的原子性
- redis的单线程操作
- redis的高性能

#### 实现流程

#### 集成方案

#### 锁的类型
第三方jar包redisson
- 可重入锁（Reentrant Lock）
- 公平锁（Fair Lock）
- 联锁（MultiLock）
- 红锁（RedLock）
- 读写锁（ReadWriteLock）

```java
// 可重入锁
RLock lock = redissonClient.getLock(key);

// 公平锁
RLock fairLock = redissonClient.getFairLock(key);

// 联锁
RLock lock1 = redissonClient.getLock("lock1");
RLock lock2 = redissonClient.getLock("lock2");
RLock lock3 = redissonClient.getLock("lock3");
RLock locks = new RedissonMultiLock(lock1, lock2, lock3);

// 红锁
RLock redLock = redissonClient.getRedLock(lock1, lock2);

// 读写锁
RReadWriteLock readWriteLock = redissonClient.getReadWriteLock(key);
```



##### readlock
```
Redlock实现
antirez提出的redlock算法大概是这样的：

在Redis的分布式环境中，我们假设有N个Redis master。这些节点完全互相独立，不存在主从复制或者其他集群协调机制。我们确保将在N个实例上使用与在Redis单实例下相同方法获取和释放锁。现在我们假设有5个Redis master节点，同时我们需要在5台服务器上面运行这些Redis实例，这样保证他们不会同时都宕掉。

为了取到锁，客户端应该执行以下操作:

获取当前Unix时间，以毫秒为单位。
依次尝试从5个实例，使用相同的key和具有唯一性的value（例如UUID）获取锁。当向Redis请求获取锁时，客户端应该设置一个网络连接和响应超时时间，这个超时时间应该小于锁的失效时间。例如你的锁自动失效时间为10秒，则超时时间应该在5-50毫秒之间。这样可以避免服务器端Redis已经挂掉的情况下，客户端还在死死地等待响应结果。如果服务器端没有在规定时间内响应，客户端应该尽快尝试去另外一个Redis实例请求获取锁。
客户端使用当前时间减去开始获取锁时间（步骤1记录的时间）就得到获取锁使用的时间。当且仅当从大多数（N/2+1，这里是3个节点）的Redis节点都取到锁，并且使用的时间小于锁失效时间时，锁才算获取成功。
如果取到了锁，key的真正有效时间等于有效时间减去获取锁所使用的时间（步骤3计算的结果）。
如果因为某些原因，获取锁失败（没有在至少N/2+1个Redis实例取到锁或者取锁时间已经超过了有效时间），客户端应该在所有的Redis实例上进行解锁（即便某些Redis实例根本就没有加锁成功，防止某些节点获取到锁但是客户端没有得到响应而导致接下来的一段时间不能被重新获取锁）。
```


#### 注意事项




### 3.2、基于zookeeper的分布式锁实现

[zookeeper分布式锁](https://baijiahao.baidu.com/s?id=1593258103626631655&wfr=spider&for=pc)


#### 实现流程
```
首先进行可重入的判定：这里的可重入锁记录在 ConcurrentMap<Thread, LockData>threadData 这个 Map 里面。 如果 threadData.get(currentThread)是有值的那么就证明是可重入锁，然后记录就会加 1。 我们之前的 MySQL 其实也可以通过这种方法去优化，可以不需要 count 字段的值，将这个维护在本地可以提高性能。
然后在我们的资源目录下创建一个节点：比如这里创建一个 /0000000002 这个节点，这个节点需要设置为 EPHEMERAL_SEQUENTIAL 也就是临时节点并且有序。
获取当前目录下所有子节点，判断自己的节点是否位于子节点第一个。
如果是第一个，则获取到锁，那么可以返回。
如果不是第一个，则证明前面已经有人获取到锁了，那么需要获取自己节点的前一个节点。 /0000000002 的前一个节点是 /0000000001，我们获取到这个节点之后，再上面注册 Watcher(这里的 Watcher 其实调用的是 object.notifyAll()，用来解除阻塞)。
object.wait(timeout) 或 object.wait()：进行阻塞等待，这里和我们第 5 步的 Watcher 相对应。

```

#### 集成方案


#### 锁的类型
- InterProcessMutex ： 可重入公平锁
- InterProcessSemaphoreMutex： 不可重入锁
- InterProcessReadWriteLock：读写锁
- InterProcessMultiLock：联锁

```java
// 可重入公平锁
InterProcessMutex interProcessMutex = new InterProcessMutex(zkClient, lockPath);

// 不可重入锁
InterProcessSemaphoreMutex processSemaphoreMutex = new InterProcessSemaphoreMutex(zkClient, lockPath);

// 联锁
// 构造方法1
List<InterProcessLock> locks = new ArrayList<>();
InterProcessMutex lock1 = new InterProcessMutex(zkClient, lockPath);
InterProcessMutex lock2 = new InterProcessMutex(zkClient, lockPath);
InterProcessMultiLock interProcessMultiLock1 = new InterProcessMultiLock(locks);
// 构造方法2
List<String> paths = Arrays.asList("/lock/lock1", "/lock/lock2", "/lock/lock3");
InterProcessMultiLock interProcessMultiLock2 = new InterProcessMultiLock(zkClient, paths);

// 读写锁
InterProcessReadWriteLock readWriteLock = new InterProcessReadWriteLock(zkClient, lockPath);

```

#### 注意事项







### 3.3、基于etcd的分布式锁实现

etcd采用Raft算法保证数据的强一致性，某次操作存储到集群中的值必然是全局一致的，所以很容易实现分布式锁。锁服务有两种使用方式，一是保持独占，二是控制时序

https://www.oschina.net/p/jetcd

https://blog.csdn.net/qq_15769369/article/details/82693107


https://segmentfault.com/a/1190000019411892?utm_source=tag-newest

etcd是一个分布式可靠的键值存储系统。它提供了与ZooKeeper相似的功能，但是使用Go语言编写而不是Java语言。Etcd使用Raft协调算法而不是ZooKeeper采用的Paxos算法。在云计算方面，Go是一个大有前景的语言，被誉为云时代的C语言。

对比与ZooKeeper，etcd更轻量级，etc更加关注一下几点：

l简单：curl命令可以调用的API接口（http+JSON）

l保密：可选的SSL客户端认证

l快速：标准检测每个实例每秒1000读写能力

l可靠：恰到地实现分布式协调，采用Raft一致性算法



保持独占，即所有试图获取锁的用户最终只有一个可以得到。etcd为此提供了一套实现分布式锁原子操作CAS（CompareAndSwap）的API。通过设置prevExist值，可以保证在多个节点同时创建某个目录时，只有一个成功，而该用户即可认为是获得了锁。
控制时序，即所有试图获取锁的用户都会进入等待队列，获得锁的顺序是全局唯一的，同时决定了队列执行顺序。etcd为此也提供了一套API（自动创建有序键），对一个目录建值时指定为POST动作，这样etcd会自动在目录下生成一个当前最大的值为键，存储这个新的值（客户端编号）。同时还可以使用API按顺序列出所有当前目录下的键值。此时这些键的值就是客户端的时序，而这些键中存储的值可以是代表客户端的编号。