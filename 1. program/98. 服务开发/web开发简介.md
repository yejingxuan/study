# web开发简介

## 1. http和https协议

> HTTP是什么？HTTP是基于TCP/IP的关于数据如何在万维网中如何通信的协议。


## 2. get和post的区别

> 参考文章：[GET和POST两种基本请求方法的区别](https://www.cnblogs.com/logsharing/p/8448446.html)

- 最直观的区别就是GET把参数包含在URL中，POST通过request body传递参数。

- GET在浏览器回退时是无害的，而POST会再次提交请求。

- GET请求只能进行url编码，而POST支持多种编码方式。

- 对参数的数据类型，GET只接受ASCII字符，而POST没有限制。

- GET比POST更不安全，因为参数直接暴露在URL上，所以不能用来传递敏感信息。

> HTTP的底层是TCP/IP。所以GET和POST的底层也是TCP/IP，也就是说，GET/POST都是TCP链接。GET和POST能做的事情是一样一样的。你要给GET加上request body，给POST带上url参数，技术上是完全行的通的。 


#### __GET和POST还有一个重大区别__

- 简单的说：GET产生一个TCP数据包；POST产生两个TCP数据包。

- 详细的说：对于GET方式的请求，浏览器会把http header和data一并发送出去，服务器响应200（返回数据）；而对于POST，浏览器先发送header，服务器响应100 continue，浏览器再发送data，服务器响应200 ok（返回数据）。


- 因为POST需要两步，时间上消耗的要多一点，看起来GET比POST更有效。因此Yahoo团队有推荐用GET替换POST来优化网站性能。但这是一个坑！跳入需谨慎。为什么？

    1. GET与POST都有自己的语义，不能随便混用。

    2. 据研究，在网络环境好的情况下，发一次包的时间和发两次包的时间差别基本可以无视。而在网络环境差的情况下，两次包的TCP在验证数据包完整性上，有非常大的优点。

    3. 并不是所有浏览器都会在POST中发送两次包，Firefox就只发送一次。



## 3. 什么是restful


