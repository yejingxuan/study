# LRU算法

- [LRU算法](#lru%e7%ae%97%e6%b3%95)
  - [一、LRU是什么？](#%e4%b8%80lru%e6%98%af%e4%bb%80%e4%b9%88)

## 一、LRU是什么？

LRU是什么？按照英文的直接原义就是Least Recently Used,最近最久未使用法。

它是按照一个非常著名的计算机操作系统基础理论得来的：最近使用的页面数据会在未来一段时期内仍然被使用,已经很久没有使用的页面很有可能在未来较长的一段时间内仍然不会被使用。基于这个思想,会存在一种缓存淘汰机制，每次从内存中找到最久未使用的数据然后置换出来，从而存入新的数据！它的主要衡量指标是使用的时间，附加指标是使用的次数。在计算机中大量使用了这个机制，它的合理性在于优先筛选热点数据，所谓热点数据，就是最近最多使用的数据！因为，利用LRU我们可以解决很多实际开发中的问题，并且很符合业务场景。