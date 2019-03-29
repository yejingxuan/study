# spring定时器

> [Schedule的源码解析](https://blog.csdn.net/weixin_40318210/article/details/78149692)


>说明：  
    1、使用@Scheduled时要在启动类上增加@EnableScheduling，使用@Async注解时要在启动类上加上@EnableAsync  
    2、下面所说的并行，是指两个不同的调度任务同时执行，并发是指同一个调度任务同时执行（即上次任务还没有执行完，下次任务已经开始执行了）


## condition1:单个任务单独执行

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

## condition2:单个任务通过子线程运行

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


## 3:一个任务并发执行

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