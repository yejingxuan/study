>参考文章：https://blog.csdn.net/zhou870498/article/details/80723697


## 一、集成activeMQ

> 参考文章：https://blog.csdn.net/AnselLyy/article/details/81044320

### 1. 所需jar包

```xml
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-activemq</artifactId>
</dependency>

<dependency>  
    <groupId>org.apache.activemq</groupId>  
    <artifactId>activemq-pool</artifactId>  
</dependency>
```

### 2. application.properties 配置

```properties
# 整合 jms
spring.activemq.broker-url=tcp://10.11.124.102:61616

spring.activemq.user=admin
spring.activemq.password=admin

spring.activemq.pool.enabled=true
spring.activemq.pool.max-connections=100
```

