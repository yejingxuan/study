# ThreadLocal知识详解

> 参考文章：https://www.toutiao.com/i6650299389178806791/  
https://www.cnblogs.com/dolphin0520/p/3920407.html


## 1.什么是ThreadLocal

        ThreadLocal，很多地方叫做线程本地变量，也有些地方叫做线程本地存储，其实意思差不多。可能很多朋友都知道ThreadLocal为变量在每个线程中都创建了一个副本，那么每个线程可以访问自己内部的副本变量。
