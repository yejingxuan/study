# logback基本配置

> 参考文章：https://www.cnblogs.com/warking/p/5710303.html  
https://www.jianshu.com/p/1ded57f6c4e3

## 1. jar包集成

        <dependency>
            <groupId>org.slf4j</groupId>
            <artifactId>slf4j-api</artifactId>
            <version>1.7.5</version>
        </dependency>
        <dependency>
            <groupId>ch.qos.logback</groupId>
            <artifactId>logback-core</artifactId>
            <version>1.0.11</version>
        </dependency>
        <dependency>
            <groupId>ch.qos.logback</groupId>
            <artifactId>logback-classic</artifactId>
            <version>1.0.11</version>
        </dependency>

## 2. 基本配置

* 根节点<configuration>，包含下面三个属性：

        scan: 当此属性设置为true时，配置文件如果发生改变，将会被重新加载，默认值为true。
        scanPeriod: 设置监测配置文件是否有修改的时间间隔，如果没有给出时间单位，默认单位是毫秒。当scan为true时，此属性生效。默认的时间间隔为1分钟。
        debug: 当此属性设置为true时，将打印出logback内部日志信息，实时查看logback运行状态。默认值为false。

        <configuration scan="true" scanPeriod="60 seconds" debug="false"> 
         <!--其他配置省略-->             
        </configuration>　
