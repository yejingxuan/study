# springboot集成ElasticSearch

> 参考文章：https://www.toutiao.com/i6648144247281156611/  
 https://blog.csdn.net/jacksonary/article/details/82729556

> github地址：https://github.com/bbossgroups/bboss-elasticsearch



## 一、集成组件






## 二、集成组件spring-data-elasticsearch

> github:https://github.com/spring-projects/spring-data-elasticsearch

> 坑比问题：https://blog.csdn.net/blackhost/article/details/84769317


1. 所需jar包

        <dependency>
            <groupId>org.springframework.data</groupId>
            <artifactId>spring-data-elasticsearch</artifactId>
            <version>3.0.10.RELEASE</version>
        </dependency>

    jar包版本参考
    
    |spring-data-elasticsearch	|elasticsearch|
    |:---:                   	|:---:	      |
    |3.1.x	                    |6.2.2        |
    |3.0.x                  	|5.5.0        |
    |2.1.x                  	|2.4.0        |
    |2.0.x                  	|2.2.0        |
    |1.3.x                     	|1.5.2        |



2. 配置文件配置

        spring:
            data:
                elasticsearch:
                cluster-nodes: localhost:9300 # 配置IP及端口号
                cluster-name: elasticsearch


        spring.data.elasticsearch.clusterNodes=47.98.48.32:9300
        spring.data.elasticsearch.clusterName=elasticsearch
        #禁用Spring boot自身的自动配置类
        spring.autoconfigure.exclude=org.springframework.boot.autoconfigure.data.elasticsearch.ElasticsearchAutoConfiguration


3. 代码集成

    




