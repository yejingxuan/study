# tomcat运行机制以及如何处理http请求

## 一、Tomcat的组成
### (1) Server

服务器元素代表整个catalina servlet容器。是单例模式。

### (2) Service
Service是这样一个集合：它由一个或者多个Connector组成，以及一个Engine，负责处理所有Connector所获得的客户请求。

### (3) Connector

一个Connector将在某个指定端口上侦听客户请求，并将获得的请求交给Engine来处理，从Engine处获得回应并返回客户。TOMCAT有两个典型的Connector，一个直接侦听来自browser的http请求，一个侦听来自其它WebServer的请求

Coyote Http/1.1 Connector 在端口8080处侦听来自客户browser的http请求
Coyote JK2 Connector 在端口8009处侦听来自其它WebServer(Apache)的servlet/jsp代理请求

 

### (4) Engine
　　Engine下可以配置多个虚拟主机Virtual Host，每个虚拟主机都有一个域名，当Engine获得一个请求时，它把该请求匹配到某个Host上，然后把该请求交给该Host来处理，Engine有一个默认虚拟主机，当请求无法匹配到任何一个Host上的时候，将交给该默认Host来处理

 

### (5) Host
　　代表一个Virtual Host，虚拟主机，每个虚拟主机和某个网络域名Domain Name相匹配，每个虚拟主机下都可以部署(deploy)一个或者多个Web App，每个Web App对应于一个Context，有一个Context path，当Host获得一个请求时，将把该请求匹配到某个Context上，然后把该请求交给该Context来处理，匹配的方法是“最长匹配”，所以一个path==""的Context将成为该Host的默认Context，所有无法和其它Context的路径名匹配的请求都将最终和该默认Context匹配

 

### (6) Context
一个Context对应于一个Web Application，一个Web Application由一个或者多个Servlet组成，Context在创建的时候将根据配置文件$CATALINA_HOME/conf/web.xml和$WEBAPP_HOME/WEB-INF/web.xml载入Servlet类，当Context获得请求时，将在自己的映射表(mapping table)中寻找相匹配的Servlet类，如果找到，则执行该类，获得请求的回应，并返回。
