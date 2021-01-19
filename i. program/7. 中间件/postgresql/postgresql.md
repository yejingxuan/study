
- [一、简介](#%e4%b8%80%e7%ae%80%e4%bb%8b)
- [二、功能特点](#%e4%ba%8c%e5%8a%9f%e8%83%bd%e7%89%b9%e7%82%b9)
- [三、部署方式](#%e4%b8%89%e9%83%a8%e7%bd%b2%e6%96%b9%e5%bc%8f)
- [四、扩展](#%e5%9b%9b%e6%89%a9%e5%b1%95)
    - [fdw功能](#fdw%e5%8a%9f%e8%83%bd)
    - [MPP介绍](#mpp%e4%bb%8b%e7%bb%8d)

### 一、简介
PostgreSQL的Slogan是“世界上最先进的开源关系型数据库”

它是一款一专多长的全栈数据库：在可观的规模内，都能做到一招鲜吃遍天。        


### 二、功能特点

成熟的应用可能会用到许许多多的数据组件（功能）：
- 缓存，
- OLTP，
- OLAP/批处理/数据仓库，
- 流处理/消息队列，
- 搜索索引，
- NoSQL/文档数据库，
- 地理数据库，
- 空间数据库，
- 时序数据库，
- 图数据库


传统架构选型可能会组合使用多种组件，典型的如：Redis + MySQL + Greenplum/Hadoop + Kafuka/Flink + ElasticSearch。在这里MySQL只能扮演OLTP关系型数据库的角色，但如果是PostgreSQL，就可以身兼多职，One hold them all，

比如：
- OLTP：事务处理是PostgreSQL的本行
- OLAP：citus分布式插件，ANSI SQL兼容，窗口函数，CTE，CUBE等高级分析功能，任意语言写UDF
- 流处理：PipelineDB扩展，Notify-Listen，物化视图，规则系统，灵活的存储过程与函数编写
- 时序数据：timescaledb时序数据库插件，分区表，BRIN索引
- 空间数据：PostGIS扩展（杀手锏），内建的几何类型支持，GiST索引。
- 搜索索引：全文搜索索引足以应对简单场景；丰富的索引类型，支持函数索引，条件索引
- NoSQL：JSON，JSONB，XML，HStore原生支持，至NoSQL数据库的外部数据包装器
- 数据仓库：能平滑迁移至同属Pg生态的GreenPlum，DeepGreen，HAWK等，使用FDW进行ETL
- 图数据：递归查询
- 缓存：物化视图

在探探的实践中，整个系统就是围绕PostgreSQL设计的。几百万日活，几百万全局DB-TPS，几百TB数据的量级下，数据组件只用了PostgreSQL。直到接近千万日活，才开始进行架构调整引入独立的数仓，消息队列和缓存。这只是验证过的规模量级，进一步压榨PG是完全可行的。


### 三、部署方式


### 四、扩展


##### fdw功能
PostgreSQL的fdw实现的功能是各个postgresql数据库及远程数据库之间的跨库操作，功能和Oracle的dblink一样。 本文中的环境如下图所示：


##### MPP介绍
MPP(大规模并行处理)简介