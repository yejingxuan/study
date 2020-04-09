# 分布式ID生成

- [分布式ID生成](#%e5%88%86%e5%b8%83%e5%bc%8fid%e7%94%9f%e6%88%90)
  - [分布式id实现分案对比](#%e5%88%86%e5%b8%83%e5%bc%8fid%e5%ae%9e%e7%8e%b0%e5%88%86%e6%a1%88%e5%af%b9%e6%af%94)
  - [snowflake算法](#snowflake%e7%ae%97%e6%b3%95)


## 分布式id实现分案对比

|名称|唯一性|性能|是否有序|内存占用|复杂程度|
|----|---|----|----|----|----|
|数据库自增主键|顺序id,数据库唯一|性能依赖数据库,性能较弱|有序|自然系列，较小|一般|
|UUID方案|系统唯一|代码生成，性能强|没有排序，无法保证趋势递增|36位字符串存储空间比较大，如果是海量数据库，就需要考虑存储量的问题。|简单|
|UUID to Int64|系统唯一|代码生成，性能强|有序|同UUID|简单|
|Redis生成唯一ID|顺序id,数据库唯一|基于内存，性能强|有序|自然系列，较小|复杂|
|zookeeper生成|多步调用API，需要考虑使用分布式锁|需要考虑分布式锁，性能较差|有序|较小|复杂|
|MongoDB的ObjectId方案|系统唯一|基于内存，性能强|有序|较小|复杂|
|snowflake算法|系统唯一|代码生成，性能强|在单机上是递增的，但是由于涉及到分布式环境，每台机器上的时钟不可能完全同步，也许有时候也会出现不是全局递增的情况。|long型id，内存较小|简单|

## snowflake算法

实现代码
```java
public class SnowFlake {

    /**
     * 起始的时间戳
     */
    private final static long START_STMP = 1480166465631L;

    /**
     * 每一部分占用的位数
     */
    private final static long SEQUENCE_BIT = 12; //序列号占用的位数
    private final static long MACHINE_BIT = 5;   //机器标识占用的位数
    private final static long DATACENTER_BIT = 5;//数据中心占用的位数

    /**
     * 每一部分的最大值
     */
    private final static long MAX_DATACENTER_NUM = -1L ^ (-1L << DATACENTER_BIT);
    private final static long MAX_MACHINE_NUM = -1L ^ (-1L << MACHINE_BIT);
    private final static long MAX_SEQUENCE = -1L ^ (-1L << SEQUENCE_BIT);

    /**
     * 每一部分向左的位移
     */
    private final static long MACHINE_LEFT = SEQUENCE_BIT;
    private final static long DATACENTER_LEFT = SEQUENCE_BIT + MACHINE_BIT;
    private final static long TIMESTMP_LEFT = DATACENTER_LEFT + DATACENTER_BIT;

    private long datacenterId;  //数据中心
    private long machineId;     //机器标识
    private long sequence = 0L; //序列号
    private long lastStmp = -1L;//上一次时间戳

    public SnowFlake(long datacenterId, long machineId) {
        if (datacenterId > MAX_DATACENTER_NUM || datacenterId < 0) {
            throw new IllegalArgumentException("datacenterId can't be greater than MAX_DATACENTER_NUM or less than 0");
        }
        if (machineId > MAX_MACHINE_NUM || machineId < 0) {
            throw new IllegalArgumentException("machineId can't be greater than MAX_MACHINE_NUM or less than 0");
        }
        this.datacenterId = datacenterId;
        this.machineId = machineId;
    }

    /**
     * 产生下一个ID
     *
     * @return
     */
    public synchronized long nextId() {
        long currStmp = getNewstmp();
        if (currStmp < lastStmp) {
            throw new RuntimeException("Clock moved backwards.  Refusing to generate id");
        }

        if (currStmp == lastStmp) {
            //相同毫秒内，序列号自增
            sequence = (sequence + 1) & MAX_SEQUENCE;
            //同一毫秒的序列数已经达到最大
            if (sequence == 0L) {
                currStmp = getNextMill();
            }
        } else {
            //不同毫秒内，序列号置为0
            sequence = 0L;
        }

        lastStmp = currStmp;

        return (currStmp - START_STMP) << TIMESTMP_LEFT //时间戳部分
                | datacenterId << DATACENTER_LEFT       //数据中心部分
                | machineId << MACHINE_LEFT             //机器标识部分
                | sequence;                             //序列号部分
    }

    private long getNextMill() {
        long mill = getNewstmp();
        while (mill <= lastStmp) {
            mill = getNewstmp();
        }
        return mill;
    }

    private long getNewstmp() {
        return System.currentTimeMillis();
    }

}
```

调用方式
```java
public static void main(String[] args) {
    SnowFlake snowFlake = new SnowFlake(2, 3);

    long start = System.currentTimeMillis();
    for (int i = 0; i < 1000000; i++) {
        System.out.println(snowFlake.nextId());
    }
    System.out.println(System.currentTimeMillis() - start);
}
```