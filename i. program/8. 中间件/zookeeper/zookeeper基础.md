
- [一、基础功能](#%e4%b8%80%e5%9f%ba%e7%a1%80%e5%8a%9f%e8%83%bd)
- [CAP定理](#cap%e5%ae%9a%e7%90%86)
- [ZAB协议](#zab%e5%8d%8f%e8%ae%ae)
- [leader选举算法和流程](#leader%e9%80%89%e4%b8%be%e7%ae%97%e6%b3%95%e5%92%8c%e6%b5%81%e7%a8%8b)

### 一、基础功能
ZooKeeper(ZK)是一个分布式开源协调服务框架，是Google的Chubby一个开源的实现，是hadoop的一个子项目

主要用来解决分布式系统的一致性问题，封装好了复杂易出错的关键服务，通过简单的接口为外部提供高性能、稳定的服务

Zookeeper可以保证如下分布式一致性特性：

- 顺序一致性：从同一个客户端发起的事务请求，最终将会严格地按照其发起顺序被应用到Zookeeper中去

- 原子性：所有事务请求的处理结果在整个集群中所有的机器上的应用情况是一致的

- 单一视图：无论客户端连接的是哪个Zookeeper服务器，其看到的服务器数据模型都是一致的

- 可靠性：一旦服务端成功地应用了一个事务，并完成对客户端的响应，那么该事务所引起的服务端状态变更将会被一直保留下来，除非有另一个事务又对其进行了变更

- 实时性：在一定的时间内，客户端最终一定能够从服务端上读取到最新的数据状态


根据这些特性，zookeeper可以用来做以下事情：
- 统一命名服务
- 发布/订阅
- 负载均衡
- 分布式配置管理
- 集群管理
- 分布式锁
- 分布式队列
- master选举mkdir


### CAP定理

CAP原则又称CAP定理，指的是在一个分布式系统中， Consistency（一致性）、 Availability（可用性）、Partition tolerance（分区容错性），三者不可得兼。
- 一致性（C）：在分布式系统中的所有数据备份，在同一时刻是否同样的值。（等同于所有节点访问同一份最新的数据副本）
- 可用性（A）：在集群中一部分节点故障后，集群整体是否还能响应客户端的读写请求。（对数据更新具备高可用性）
- 分区容错性（P）：以实际效果而言，分区相当于对通信的时限要求。系统如果不能在时限内达成数据一致性，就意味着发生了分区的情况，必须就当前操作在C和A之间做出选择。

__在分布式系统中CA是不可能同时存在的，而zookeeper采用了CP来保证数据的强一致性__
![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20200409162931.png)

### ZAB协议
Zab协议 的全称是 Zookeeper Atomic Broadcast （Zookeeper原子广播）。
Zookeeper 是通过 Zab 协议来保证分布式事务的最终一致性。

Zab协议要求每个 Leader 都要经历三个阶段：发现，同步，广播。

发现：要求zookeeper集群必须选举出一个 Leader 进程，同时 Leader 会维护一个 Follower 可用客户端列表。将来客户端可以和这些 Follower节点进行通信。

同步：Leader 要负责将本身的数据与 Follower 完成同步，做到多副本存储。这样也是提现了CAP中的高可用和分区容错。Follower将队列中未处理完的请求消费完成后，写入本地事务日志中。

广播：Leader 可以接受客户端新的事务Proposal请求，将新的Proposal请求广播给所有的 Follower。



### leader选举算法和流程