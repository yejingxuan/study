

- [一、redis特性](#%e4%b8%80redis%e7%89%b9%e6%80%a7)
- [二、redis部署](#%e4%ba%8credis%e9%83%a8%e7%bd%b2)
- [三、redis基础指令](#%e4%b8%89redis%e5%9f%ba%e7%a1%80%e6%8c%87%e4%bb%a4)
  - [1. 查询指令](#1-%e6%9f%a5%e8%af%a2%e6%8c%87%e4%bb%a4)
  - [2. 增加指令](#2-%e5%a2%9e%e5%8a%a0%e6%8c%87%e4%bb%a4)
  - [3. 删除指令](#3-%e5%88%a0%e9%99%a4%e6%8c%87%e4%bb%a4)
  - [4. 更改指令](#4-%e6%9b%b4%e6%94%b9%e6%8c%87%e4%bb%a4)
- [参考文章](#%e5%8f%82%e8%80%83%e6%96%87%e7%ab%a0)


### 一、redis特性

### 二、redis部署

### 三、redis基础指令
#### 1. 查询指令

* 获取所有键：keys *

* 获取键值（string）：get key



* 获取键值（list）：
        
        1. 两边获取语法：lrange key start stop 
        2. 获取所有：lrange key 0 -1
        3. 获取键值长度（list）:llen key

* 获取键值（set）：

        1. 获取键值：smembers key
        2. 获取键值长度：scard key

* 获取键值（hash）:

        1. 获取hash的所有键值：hgetall key
        2. 获取hash的某个键的值：hget key key

* 获取键总数：dbsize

* 查询键是否存在：exists key [key …]  

        查询查询多个，返回存在的个数。

* 查询键类型：type key

* 查询键生命周期：pttl(ttl) key 
        
        ttl：秒语法
        pttl：毫秒语法
        -1：永远不过期
        -2：键不存在

    __设置key的失效时间：expire key milliseconds__ 

    __设置key永不过期：persist key__



#### 2. 增加指令

* 新增key(string)：set key value
* 新增key(list)
        
        1. 将一个或多个值插入到列表头部：lpush key value1 value2
        2. 在尾部插入值：rpush key value1 value2

* 新增key(set)

        1. 在头部新增值：sadd key value1 value2



#### 3. 删除指令

* 删除key：del key1 key2




#### 4. 更改指令

* 重命名键值：rename key newkey



### 参考文章
> redis参考文章：https://www.toutiao.com/i6616478418596790798/
> jedis参考文章：https://blog.csdn.net/zhangguanghui002/article/details/78770071