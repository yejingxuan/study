# 多线程&&锁进阶

- [多线程&&锁进阶](#多线程锁进阶)
  - [一、多线程基础](#一多线程基础)
    - [1.1、多线程内存模型简介](#11多线程内存模型简介)
    - [1.2、指令重排](#12指令重排)
    - [1.3、内存屏障](#13内存屏障)
    - [1.4、创建线程的方式](#14创建线程的方式)
    - [1.5、线程上下文切换](#15线程上下文切换)
  - [二、线程状态](#二线程状态)
    - [2.1、线程的运行状态](#21线程的运行状态)
    - [2.2、线程的中断](#22线程的中断)
    - [2.3、wait和sleep的区别](#23wait和sleep的区别)
    - [2.4、start和run的区别](#24start和run的区别)
    - [2.5、线程的顺序执行](#25线程的顺序执行)
    - [2.5、线程之间数据共享](#25线程之间数据共享)
  - [三、线程池](#三线程池)
    - [3.1、线程池的优点](#31线程池的优点)
    - [3.2、线程池创建](#32线程池创建)
    - [3.3、工作流程](#33工作流程)
    - [3.4、阻塞队列](#34阻塞队列)
    - [3.5、forkjoin](#35forkjoin)
  - [四、锁 && 线程安全](#四锁--线程安全)
    - [4.1、java中的各种锁](#41java中的各种锁)
    - [4.2、synchronize同步锁](#42synchronize同步锁)
    - [4.3、synchronize锁升级过程](#43synchronize锁升级过程)
    - [4.4、AQS && ReentantLock](#44aqs--reentantlock)
    - [4.5、ReentrantLock和synchronize区别](#45reentrantlock和synchronize区别)
    - [4.5、threadlocal](#45threadlocal)
    - [4.6、volatile](#46volatile)
    - [4.6、Atomic](#46atomic)
  - [参考文章](#参考文章)

## 一、多线程基础

### 1.1、多线程内存模型简介
![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20200526163314.png)

![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20200530162123.png)


- __可见性__
  - 是值每个线程在自己工作内存中操作的数据对于其他结果是不可见的，可以使用
  - volatile关键字来保证变量的可见性

- __有序性__
  - 指对于代码，cpu会有一个指令重排的过程，导致执行顺序没有按照代码的顺序执行。
  - volatile关键字来保证变量有序性

- __原子性__
  - 线程A执行操作的时候，线程B不能中途插入，即线程A的read->load->use->assign->store->write操作是整体执行的。
  - Atmoic可以保障线程的原子性

### 1.2、指令重排

- 指令重排是指JVM在编译Java代码的时候，或者CPU在执行JVM字节码的时候，对现有的指令顺序进行重新排序。
- 指令重排的目的是为了在不改变程序执行结果的前提下，优化程序的运行效率。需要注意的是，这里所说的不改变执行结果，指的是不改变单线程下的程序执行结果。
- 指令重排是一把双刃剑，虽然优化了程序的执行效率，但是在某些情况下，会影响到多线程的执行结果。
- __在什么情景下需要禁止指令重排__
  1. 锁定规则，加锁的时候，lock先于unlock
  2. 按顺序执行的传递规则，例如执行A->B操作，不能反过来
  3. 线程启动规则，线程的start操作需要先发生于线程的其他操作
  4. 线程中断规则
   

### 1.3、内存屏障

- 内存屏障又叫happen-before，使用内存屏障可以避免指令重排的问题。


### 1.4、创建线程的方式

1. 集成Thread类
2. 实现runnable接口
3. 实现callable接口,创建futureTask
4. 继承RecursiveTask类，通过ForkJoinPool来创建
5. 通过线程池ThreadPoolExecutor来创建

### 1.5、线程上下文切换
- 对于单核CPU来说（对于多核CPU，此处就理解为一个核），CPU在一个时刻只能运行一个线程，当在运行一个线程的过程中转去运行另外一个线程，这个叫做线程上下文切换（对于进程也是类似）。

- 由于可能当前线程的任务并没有执行完毕，所以在切换时需要保存线程的运行状态，以便下次重新切换回来时能够继续切换之前的状态运行。举个简单的例子：比如一个线程A正在读取一个文件的内容，正读到文件的一半，此时需要暂停线程A，转去执行线程B，当再次切换回来执行线程A的时候，我们不希望线程A又从文件的开头来读取


## 二、线程状态

### 2.1、线程的运行状态
![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20200509171056.png)

- __new__ :
  - 创建后尚未启动的线程。


- __runnable__ 
  - 就绪状态，可运行状态，调用了线程的start方法，等待CPU调度


- __blocked__ 
  - 阻塞状态下, 是在多个线程有同步操作的场景, 比如正在等待另一个线程的synchronized 块的执行释放, 或者可重入的 synchronized块里别人调用wait() 方法, 也就是这里是线程在等待进入临界区

- __waiting__  
  - 无限期等待状态，它们要等其他线程显示的唤醒。
  - 进入wait状态的条件
    - 没有设置 Timeout 参数的 Object.wait() 方法。
    - 没有设置 Timeout 参数的 Thread.join() 方法。
    - LookSupport.park() 方法。
  - 阻塞和等待的区别
    - 阻塞是被动的，它是在等待获取 monitor lock；等待是主动的，通过调用 Object.wait() 等方法进入。


- __timed_waiting__
  - 限期等待，无需等待其它线程显式地唤醒，在一定时间之后会被系统自动唤醒。
  - 进入timed_waiting状态的条件
    - Thread.sleep() 方法，时间结束后自动退出
    - 设置了 Timeout 参数的 Object.wait() 方法；等时间结束 / Object.notify() / Object.notifyAll()会退出
    - 设置了 Timeout 参数的 Thread.join() 方法；等时间结束 / 被调用的线程执行完毕会会退出

- __terminated__ 
  - 这个状态下表示 该线程的run方法已经执行完毕了, 基本上就等于死亡了(当时如果线程被持久持有, 可能不会被回收)


### 2.2、线程的中断
- __sleep__
  - Thread.sleep(millisec) 方法会休眠当前正在执行的线程
- __wait与notify__
  - Thread.wait()方法会让线程进入等待状态，直到在其他线程调用此对象的notify()方法或notifyAll()方法将其唤醒。
  - 在调用wait()之前，线程必须要获得该对象的对象级别锁，因此只能在同步方法或同步块中调用wait()方法。
- __yield()__
  - yield()让当前运行线程回到可运行状态,声明了当前线程已经完成了生命周期中最重要的部分，可以切换给其它线程来执行。该方法只是对线程调度器的一个建议，而且也只是建议具有相同优先级的其它线程可以运行
- __interrupt()__
  - 通过调用一个线程的 interrupt() 来中断该线程，如果该线程处于阻塞、限期等待或者无限期等待状态，那么就会抛出 InterruptedException，从而提前结束该线程。但是不能中断 I/O 阻塞和 synchronized 锁阻塞。
  - 如果一个线程的 run() 方法执行一个无限循环，并且没有执行 sleep() 等会抛出 InterruptedException 的操作，那么调用线程的 interrupt() 方法就无法使线程提前结束。
  - 但是调用 interrupt() 方法会设置线程的中断标记，此时调用 interrupted() 方法会返回 true。因此可以在循环体中使用 interrupted() 方法来判断线程是否处于中断状态，从而提前结束线程。

- __shutdown()__
  - 调用 Executor 的 shutdown() 线程池的状态则立刻变成SHUTDOWN状态。此时，则不能再往线程池中添加任何任务，否则将会抛出RejectedExecutionException异常。但是，此时线程池不会立刻退出，直到添加到线程池中的任务都已经处理完成，才会退出
  
- __shutdownNow()__
  - 相当于调用每个线程的 interrupt() 方法

- __cancel(true)__
  - 如果只想中断 Executor 中的一个线程，可以通过使用 submit() 方法来提交一个线程，它会返回一个 Future<?> 对象，通过调用该对象的 cancel(true) 方法就可以中断线程。


### 2.3、wait和sleep的区别
  - 对于sleep()方法，我们首先要知道该方法是属于Thread类中的。而wait()方法，则是属于Object类中的。
  - sleep()方法的过程中，线程进入TIMED_WAITING时间，并不会释放锁，在设定时间到或被interrupt后抛出InterruptedException后进入RUNNABLE状态。而当调用wait()方法的时候，线程会放弃对象锁，进入waiting状态，只有针对此对象调用notify()方法后本线程才进入blocked状态，拿到锁内进入runnable状态
  - wait，notify和notifyAll只能在同步控制方法或者同步控制块里面使用，而sleep可以在任何地方使用
  - sleep必须捕获异常，而wait，notify和notifyAll不需要捕获异常。

### 2.4、start和run的区别
  - 线程的任务处理逻辑可以在Tread类的run实例方法中直接实现或通过该方法进行调用，因此run()相当于线程的任务处理逻辑的入口方法，它由Java虚拟机在运行相应线程时直接调用，而不是由应用代码进行调用。
  - 而start()的作用是启动相应的线程。启动一个线程实际是请求Java虚拟机运行相应的线程，而这个线程何时能够运行是由线程调度器决定的。start()调用结束并不表示相应线程已经开始运行，这个线程可能稍后运行，也可能永远也不会运行。

### 2.5、线程的顺序执行

- __join__
  - 在当前线程中，如果调用某个thread的join方法，那么当前线程就会被阻塞，直到thread线程执行完毕，当前线程才能继续执行。
  - join的原理是，不断的检查join加入的thread是否存活，如果存活，那么让当前线程一直wait，直到oin加入的thread线程终止，当前线程的this.notifyAll 就会被调用

- __CountDownLatch__
  - CountDownLatch中我们主要用到两个方法一个是await()方法，调用这个方法的线程会被阻塞，另外一个是countDown()方法，调用这个方法会使计数器减一，当计数器的值为0时，因调用await()方法被阻塞的线程会被唤醒，继续执行
  - 构造函数
    ```java
    final CountDownLatch latch = new CountDownLatch(2);
    ```

- __CyclicBarrier__
  - 字面意思回环栅栏，通过它可以实现让一组线程等待至某个状态之后再全部同时执行
  - 主要方法
    ```java
    //用来挂起当前线程，直至所有线程都到达 barrier 状态再同时执行后续任务；
    public int await();
    // 让这些线程等待至一定的时间，如果还有线程没有到达 barrier 状态就直接让到达 barrier 的线程执行后续任务
    public int await(long timeout, TimeUnit unit);
    ```

- __Semaphore__
  - Semaphore 可以控制同时访问的线程个数， 通过acquire() 获取一个许可， 使用release()释放一个许可
  - 主要方法
    ```java
    Semaphore semaphore = new Semaphore(5); //机器数目
    for(int i=0;i<N;i++)
      new Worker(i,semaphore).start();
    }

    class Worker extends Thread{
      private int num;
      private Semaphore semaphore;
      public Worker(int num, Semaphore semaphore){
        this.num = num;
        this.semaphore = semaphore;
      }
      @Override
      public void run() {
        try {
          semaphore.acquire();
          System.out.println("工人"+this.num+"占用一个机器在生产...");
          Thread.sleep(2000);
          System.out.println("工人"+this.num+"释放出机器");
          semaphore.release();
        } catch (InterruptedException e) {
          e.printStackTrace();
        } 
    }
    ```

- __CountDownLatch 和 join的区别__
  - 调用join方法需要等待thread执行完毕才能继续向下执行
  - 而CountDownLatch只需要检查计数器的值为零就可以继续向下执行，相比之下，CountDownLatch更加灵活一些


- __CountDownLatch和CyclicBarrier区别__
  - CountDownLatch 是一次性的，CyclicBarrier 是可循环利用的
  - CyclicBarrier适用于用于多线程计算数据，最后合并计算结果的场景，而CountDownLatch更合适用于线程通知

### 2.5、线程之间数据共享
- 通过构造函数传入共享数据
- 通过匿名内部类传入共享数据


## 三、线程池

### 3.1、线程池的优点
  - 第一：降低资源消耗。通过重复利用已创建的线程降低线程创建和销毁造成的消耗。
  - 第二：提高响应速度。当任务到达时，任务可以不需要等到线程创建就能立即执行。
  - 第三：提高线程的可管理性。线程是稀缺资源，如果无限制地创建，不仅会消耗系统资源，还会降低系统的稳定性，使用线程池可以进行统一分配、调优和监控。但是，要做到合理利用线程池，必须对其实现原理了如指掌

### 3.2、线程池创建
- __ThreadPoolExecutor构造方法详解：__
  |参数名 |作用|
  |:--:|:--:|
  |corePoolSize   |核心线程池大小|
  |maximumPoolSize|最大线程池大小|
  |keepAliveTime  |线程池中超过corePoolSize数目的空闲线程最大存活时间|
  |TimeUnit       |keepAliveTime时间单位|
  |workQueue      |阻塞任务队列|
  |threadFactory  |线程工厂，用于创建线程，一般用默认的即可|
  |handler        |拒绝策略，当任务太多来不及处理，如何拒绝任务|

- __major方法__
  - __allowCoreThreadTimeOut(true)：__
    - 线程池回收线程只会发生在当前线程池中线程数量大于corePoolSize参数的时候；当线程池中线程数量小于等于corePoolSize参数的时候，回收过程就会停止。
    - allowCoreThreadTimeOut设置项可以要求线程池：将包括“核心线程”在内的，没有任务分配的任何线程，在等待keepAliveTime时间后全部进行回收
  - __prestartAllCoreThreads()：__
    - 可以在线程池创建，但还没有接收到任何任务的情况下，先行创建符合corePoolSize参数值的线程数

- __构造函数__
  - __newSingleExecutor__
    ```java
    public static ExecutorService newSingleThreadExecutor() {
        return new FinalizableDelegatedExecutorService
            (new ThreadPoolExecutor(1, 1,
                                    0L, TimeUnit.MILLISECONDS,
                                    new LinkedBlockingQueue<Runnable>()));
    }
    ```
    - 因为LinkedBlockingQueue是长度为Integer.MAX_VALUE的队列，可以认为是无界队列，因此往队列中可以插入无限多的任务，在资源有限的时候容易引起OOM异常，同时因为无界队列，maximumPoolSize和keepAliveTime参数将无效，压根就不会创建非核心线程

  - __newFixedThreadPool__
    ```java
    public static ExecutorService newFixedThreadPool(int nThreads) {
        return new ThreadPoolExecutor(nThreads, nThreads,
                                      0L, TimeUnit.MILLISECONDS,
                                      new LinkedBlockingQueue<Runnable>());
    }

    ```
    - 允许的请求队列长度为Integer.MAX_VALUE，可能会堆积大量的请求，从而引起OOM异常


  - __newCachedThreadPool__
    ```java
    public static ExecutorService newCachedThreadPool() {
        return new ThreadPoolExecutor(0, Integer.MAX_VALUE,
                                      60L, TimeUnit.SECONDS,
                                      new SynchronousQueue<Runnable>());
    }
    ```
    - 当一个任务提交时，corePoolSize为0不创建核心线程，SynchronousQueue是一个不存储元素的队列，可以理解为队里永远是满的，因此最终会创建非核心线程来执行任务。
  
    - 对于非核心线程空闲60s时将被回收。因为Integer.MAX_VALUE非常大，可以认为是可以无限创建线程的，在资源有限的情况下容易引起OOM异常
  
  - __总结__
    - FixedThreadPool和SingleThreadExecutor => 允许的请求队列长度为Integer.MAX_VALUE，可能会堆积大量的请求，从而引起OOM异常
  
    - CachedThreadPool => 允许创建的线程数为Integer.MAX_VALUE，可能会创建大量的线程，从而引起栈OOM异常

### 3.3、工作流程

- __线程工作流程__

  ![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/线程池.jpg)

  1. 判断线程池里的核心线程是否都在执行任务，如果不是（核心线程空闲或者还有核心线程没有被创建）则创建一个新的工作线程来执行任务。如果核心线程都在执行任务，则进入下个流程。

  2. 线程池判断工作队列是否已满，如果工作队列没有满，则将新提交的任务存储在这个工作队列里。如果工作队列满了，则进入下个流程。

  3. 判断线程池里的线程是否都处于工作状态，如果没有，则创建一个新的工作线程来执行任务。如果已经满了，则交给饱和策略来处理这个任务。

  - __线程池初始化时，核心线程数量为0__

- __RejectedExecutionHandler：饱和策略__
  - 1、AbortPolicy（默认策略）：直接抛出异常
  - 2、CallerRunsPolicy：调用所在的主线程运行任务
  - 3、DiscardOldestPolicy：丢弃队列里最近的一个任务，并执行当前任务。
  - 4、DiscardPolicy：不处理，丢弃掉。


### 3.4、阻塞队列


### 3.5、forkjoin


## 四、锁 && 线程安全

### 4.1、java中的各种锁

- __乐观锁__
    - 乐观锁是一种乐观思想，即认为读多写少，遇到并发写的可能性低，每次去拿数据的时候都认为别人不会修改，所以不会上锁，但是在更新的时候会判断一下在此期间别人有没有去更新这个数据，采取在写时先读出当前版本号，然后加锁操作（比较跟上一次的版本号，如果一样则更新），如果失败则要重复读-比较-写的操作
    - __CAS:CAS是英文单词Compare And Swap的缩写，翻译过来就是比较并替换__。
  
- __悲观锁__
  - 悲观锁是就是悲观思想，即认为写多，遇到并发写的可能性高，每次去拿数据的时候都认为别人会修改，所以每次在读写数据的时候都会上锁，这样别人想读写这个数据就会 block 直到拿到锁。java中的悲观锁就是 Synchronized

- __自旋锁__
  - 自旋锁原理非常简单， 如果持有锁的线程能在很短时间内释放锁资源，那么那些等待竞争锁的线程就不需要做内核态和用户态之间的切换进入阻塞挂起状态，它们只需要等一等（自旋），等持有锁的线程释放锁后即可立即获取锁，这样就避免用户线程和内核的切换的消耗。

- __可重入锁__
  - 当一个线程拥有锁的时候，下次可以直接不用获得锁，直接执行。

- __公平锁__
  - 线程竞争锁的时候进入一个队列里，先进先出，先来的先获取锁


### 4.2、synchronize同步锁

- __synchronize是java中的一个关键字，它可以对方法、代码块进行加锁操作__
  - 同步普通方法，锁的是当前对象。
  - 同步静态方法，锁的是当前 Class 对象。
  - 同步代码块，锁的是 {} 中的对象。
- __synchronized的原理__
  - JVM 是通过进入、退出对象监视器( Monitor )来实现对方法、同步块的同步的。

  - 具体实现是在编译之后在同步方法调用前加入一个 monitor.enter 指令，在退出方法和异常处插入 monitor.exit 的指令。

  - 其本质就是对一个对象监视器( Monitor )进行获取，而这个获取过程具有排他性从而达到了同一时刻只能一个线程访问的目的。

  - 而对于没有获取到锁的线程将会阻塞到方法入口处，直到获取锁的线程 monitor.exit 之后才能尝试继续获取锁。


### 4.3、synchronize锁升级过程

- __什么是锁升级__
  - synchronize锁一开始的时候认为线程之间没有发生资源竞争关系，就默认给对象设置的偏向锁，保证单个线程实现锁的可重入和轻量级判断。而后随着各种线程资源竞争的情况发生，不断的对锁进行升级。
  - __升级过程为：偏向锁 -> 轻量级锁（无锁、自旋锁） -> 重量级锁__

- __锁升级流程__
  - 刚开始加上偏向锁，意思为偏向第一个申请锁的线程，在对象头中记录持有锁的线程的指针，来实现可重入。
  - 有资源开始竞争升级为轻量级锁，在对象头中记录线程栈的lock record指针，采用自旋的方式进行竞争。
  - 资源竞争较大（自旋10次）升级为重量级锁，未获得锁的线程进入阻塞队列中等待

- __锁降级__
  - 只有在GC的情况下可能会发生

- __锁消除__
  - 例如springbuffer中的add方法加了synchronize锁，当只有一个线程执行add方法时，就会自动消除锁

- __锁粗化__
  - 例如循环100次进行springbuffer的add方法，要进行100次的加锁，解锁，jvm就会将加锁的范围粗化到这一连串的循环操作外部，使得只需加一次锁


### 4.4、AQS && ReentantLock

- __AQS（抽象队列同步器）__
  - AQS全程 AbstractQueuedSynchronizer ，翻译过来就是抽象队列同步器的意思，其实就是一个多线程队列同步的抽象类，实现这个类，并按照AQS的约定规范来进行加锁，解锁。
  
  - __AQS的核心参数__
    - state:资源状态，记录锁的状态，并通过CAS来修改状态
    - FIFO队列：实现资料的排队
    ![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20200528165641.png)


- __ReentantLock实现流程__
  ![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20200530105547.png)

  - ReentantLock是AQS的具体实现

  - ReentantLock 继承接口 Lock 并实现了接口中定义的方法， 他是一种可重入锁， 除了能完成 synchronized 所能完成的所有工作外，还提供了诸如可响应中断锁、可轮询锁请求、定时锁等避免多线程死锁的方法。

- __ReentrantLock实现可重入分析__
  根据state状态 和 getExclusiveOwnerThread()获取当前持有锁的线程

  ```java
  final boolean nonfairTryAcquire(int acquires) {
      // 获取当前线程
      final Thread current = Thread.currentThread();
      // 获取当前同步状态
      int c = getState();
      // 如果为0代表没加锁
      if (c == 0) {
          // 直接尝试获取锁
          if (compareAndSetState(0, acquires)) {
              setExclusiveOwnerThread(current);
              return true;
          }
      }
      // 判断当前线程是否已经获取了该锁，如果是实现可重入
      else if (current == getExclusiveOwnerThread()) {
          // 实现可重入，将同步状态 + 1
          int nextc = c + acquires;
          if (nextc < 0) // overflow
              throw new Error("Maximum lock count exceeded");
          setState(nextc);
          return true;
      }
      return false;
  }
  ```

- __公平锁 / 非公平锁实现分析__
  ```java
  //Nonfair-lock
  final void lock() {
	    if (compareAndSetState(0, 1))
	        setExclusiveOwnerThread(Thread.currentThread());
	    else
	        acquire(1);
	}

  //fair-lock
  final void lock() {
	    acquire(1);
	}
  ```
  - 公平锁中，每一次的tryAcquire都会检查CLH队列中是否仍有前驱的元素，如果仍然有那么继续等待，通过这种方式来保证先来先服务的原则。
  - 非公平锁，首先是检查并设置锁的状态，新的线程可能会抢占已经在排队的线程的锁，这样就无法保证先来先服务，但是已经等待的线程们是仍然保证先来先服务的


- __可中断锁__
  - ReentrantLock中的lockInterruptibly()方法使得线程可以在被阻塞时响应中断
  - 比如一个线程t1通过lockInterruptibly()方法获取到一个可重入锁，并执行一个长时间的任务，另一个线程通过interrupt()方法就可以立刻打断t1线程的执行，来获取t1持有的那个可重入锁
  - 原理：
    - 是如果监测到中断就直接throw new InterruptedException();抛出来了，所以可以响应中断。而非中断锁只是把中断状态记录了下来等得到同步状态后再处理

### 4.5、ReentrantLock和synchronize区别

1. ReentrantLock 显示的获得、释放锁， synchronized 隐式获得释放锁
2. ReentrantLock 可响应中断，synchronized 是不可以响应中断的
3. ReentrantLock 可以实现公平锁
4. ReentrantLock 是 API 级别的， synchronized 是 JVM 级别的
5. Lock 是一个接口，而 synchronized 是 Java 中的关键字
6. synchronized 在发生异常时，会自动释放线程占有的锁，因此不会导致死锁现象发生；而 Lock 在发生异常时，如果没有主动通过 unLock()去释放锁，则很可能造成死锁现象，因此使用 Lock 时需要在 finally 块中释放锁
7. 通过 Lock 可以知道有没有成功获取锁，而 synchronized 却无法办到
8. Lock 可以提高多个线程进行读操作的效率，既就是实现读写锁等

### 4.5、threadlocal

- __原理__
  - 每个线程Thread对象都包含一个成员变亮，threadlocalMap，threadlocald的put，get方法实际操作的是线程独有的threadlocalMap变量
  - threadLocalMap将当前线程对象t作为key，要存储的对象作为value存到map里面去。如果该Map不存在，则初始化一个。

- __内存泄露问题__
  ![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20200522181031.png)
  - 通过源码知道threadLocalMap是threadlocal的一个静态内部类，主要用entry存储数据，而entry还是弱引用
    ![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20200522181629.png)
  - 为什么entry是弱引用呢，主要为了解决内存泄露问题。假如entry是强引用，当threadlocal的值在运行结束后被设置为null，但是threadlocalmap中还保存这threadlocal的强引用，导致threadlocal无法被回收。产生内存泄露，而当是虚引用的时候，则完全可以回收掉threadlocal。
  - 但是当threadlocal作为map的key被回收后，value却无法被回收，所以我们最好用完threadlocal后手动remove掉key（jdk在每次put和get的时候都回去remove key）。


### 4.6、volatile

- __一个变量被定义为 volatile 的特性：__
    1. 保证此变量对所有线程的可见性
       - 每个线程操作数据的时候会把数据从主内存读取到自己的工作内存，__如果他操作了数据并且写了，其他已经读取的线程的变量副本就会失效了__，需要都数据进行操作又要再次去主内存中读取了。

    2. 禁止指令重排序优化。
       - 通过插入内存屏障保证一致性。volatile写是在前面和后面分别插入内存屏障，而volatile读操作是在后面插入两个内存屏障。
        ```
        1、在每个volatile写操作前插入StoreStore屏障，在写操作后插入StoreLoad屏障。
        2、在每个volatile读操作前插入LoadLoad屏障，在读操作后插入LoadStore屏障。
        ```

- __MESI（缓存一致性协议）__
  - 当CPU写数据时，如果发现操作的变量是共享变量，即在其他CPU中也存在该变量的副本，会发出信号通知其他CPU将该变量的缓存行置为无效状态，因此当其他CPU需要读取这个变量时，发现自己缓存中缓存该变量的缓存行是无效的，那么它就会从内存重新读取。


- __DCL(double check lock)问题__

- __volatile应用场景__
  - volatile修饰符适用于以下场景：某个属性被多个线程共享，其中有一个线程修改了此属性，其他线程可以立即得到修改后的值，比如booleanflag;或者作为触发器，实现轻量级同步。
  - volatile属性的读写操作都是无锁的，它不能替代synchronized，因为它没有提供原子性和互斥性。因为无锁，不需要花费时间在获取锁和释放锁上，所以说它是低成本的。
  - volatile可以在单例双重检查中实现可见性和禁止指令重排序，从而保证安全性。

- __volatile原子性问题__
  - volatile并不能保证原子性操作，使用的时候需要注意这一点。

### 4.6、Atomic

- __Atomic子类__
  - 标量类：AtomicBoolean，AtomicInteger，AtomicLong，AtomicReference
  - 数组类：AtomicIntegerArray，AtomicLongArray，AtomicReferenceArray
  - 更新器类：AtomicLongFieldUpdater，AtomicIntegerFieldUpdater，AtomicReferenceFieldUpdater
  - 复合变量类：AtomicMarkableReference，AtomicStampedReference

- __实现原理__
  - 基于volatile + CAS原理实现。
    ```java
    public final int incrementAndGet() {
      for (;;) {
          int current = get();
          int next = current + 1;
          if (compareAndSet(current, next))
              return next;
      }
    }

    public final boolean compareAndSet(int expect, int update) {
        return unsafe.compareAndSwapInt(this, valueOffset, expect, update);
    }
    ```
  - 为先获取到当前的 value 属性值，然后将 value 加 1，赋值给一个局部的 next 变量，然而，这两步都是非线程安全的，但是内部有一个死循环，不断去做compareAndSet操作
  - compareAndSwapInt 是一个native方法，基于的是CPU 的 CAS指令来实现的。所以基于 CAS 的操作可认为是无阻塞的


## 参考文章
- [大白话聊聊Java并发面试问题之volatile到底是什么？【石杉的架构笔记】](https://mp.weixin.qq.com/s?__biz=MzU0OTk3ODQ3Ng==&mid=2247484058&idx=1&sn=d5c1533204ea655e65947ec57f924799&chksm=fba6ea99ccd1638f945c585cf3b2df6f4d4112b17ea3648730d50fdb5508555d5f30316f4186&mpshare=1&scene=1&srcid=0608cDtcDBaGNgIU9v4zxS3f%23rd)
- [漫画：volatile对指令重排的影响
](https://www.itcodemonkey.com/article/1725.html)
- [ReentrantLock中中断锁和非中断锁源码分析](https://blog.csdn.net/qq_37859539/article/details/85697007?utm_medium=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-1.nonecase&depth_1-utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-1.nonecase)
- [CountDownLatch与thread.join()的区别](https://www.jianshu.com/p/795151ac271b)