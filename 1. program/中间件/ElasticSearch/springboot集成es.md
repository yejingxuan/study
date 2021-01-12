# springboot集成ElasticSearch

[TOC]


> 参考文章：https://www.toutiao.com/i6648144247281156611/  
 https://blog.csdn.net/jacksonary/article/details/82729556
> github地址：https://github.com/bbossgroups/bboss-elasticsearch
> github:https://github.com/spring-projects/spring-data-elasticsearch
> 坑比问题：https://blog.csdn.net/blackhost/article/details/84769317


## 一、 所需jar包
本人用的es版本是5.6.9，集成的jar包版本需要根据es实际版本来调整
```xml
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-data-elasticsearch</artifactId>
    <version>2.0.9.RELEASE</version>
</dependency>
```
jar包版本参考：

|spring-data-elasticsearch	|elasticsearch|
|:---:                   	|:---:	      |
|3.1.x	                    |6.2.2        |
|3.0.x                  	|5.5.0        |
|2.1.x                  	|2.4.0        |
|2.0.x                  	|2.2.0        |
|1.3.x                     	|1.5.2        |



## 二、 ElasticSearch配置

### 2.1、配置文件配置

```properties
#es配置
elasticsearch.host=127.0.0.1
elasticsearch.port=9300
elasticsearch.clustername=yjx-es
```

### 2.2、配置类
```java
package com.yjx.homeweb.config;

import com.yjx.common.utils.LogUtil;
import java.net.InetAddress;
import org.elasticsearch.client.Client;
import org.elasticsearch.common.settings.Settings;
import org.elasticsearch.common.transport.TransportAddress;
import org.elasticsearch.transport.client.PreBuiltTransportClient;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.elasticsearch.core.ElasticsearchOperations;
import org.springframework.data.elasticsearch.core.ElasticsearchTemplate;
import org.springframework.data.elasticsearch.repository.config.EnableElasticsearchRepositories;


@Configuration
@EnableElasticsearchRepositories(basePackages = "com.yjx.**.repository")
public class ElasticSearchConfig {

    @Value("${elasticsearch.host}")
    private String esHost;

    @Value("${elasticsearch.port}")
    private int esPort;

    @Value("${elasticsearch.clustername}")
    private String esClusterName;


    @Bean
    public Client client() throws Exception {

        Settings esSettings = Settings.builder()
                .put("cluster.name", esClusterName)
                .put("client.transport.sniff", true)//增加嗅探机制，找到ES集群
                .put("thread_pool.search.size", Integer.parseInt("5"))//增加线程池个数，暂时设为5
                .build();

        return new PreBuiltTransportClient(esSettings)
                .addTransportAddress(new TransportAddress(InetAddress.getByName(esHost), esPort));
    }

    @Bean
    public ElasticsearchOperations elasticsearchTemplateCustom() throws Exception {
        ElasticsearchTemplate elasticsearchTemplate;
        try {
            elasticsearchTemplate = new ElasticsearchTemplate(client());
            return elasticsearchTemplate;
        } catch (Exception e) {
            LogUtil.getLoger().error("初始化ElasticsearchTemplate失败！");
            return new ElasticsearchTemplate(client());
        }

    }
}
```

## 三、 代码集成


model层：
```java
@Document(indexName = "stu", type = "doc")
@Data
public class UserInfoEs implements Serializable {

    /**
     *
     */
    private static final long serialVersionUID = 1L;

    @Id
    //@Field
    private String id;
    // @Field()
    private String stuId;
    // @Field
    private String stuName;
    // @Field
    private String createTime;


}
```


dao层：
```java
@Repository
public interface UserInfoEsRepository extends ElasticsearchRepository<UserInfoEs, String> {

}
```

controller层：
```java
@RestController
public class EsController {

    @Autowired
    private UserInfoEsRepository userInfoEsRepository;

    @PutMapping(value = "/save")
    public String insertUserInfo(UserInfoEs userInfoEs){
        UserInfoEs save = userInfoEsRepository.save(userInfoEs);
        return "success";
    }

    @GetMapping(value = "/getUser")
    public Iterable<UserInfoEs> getUserInfo(){
        Iterable<UserInfoEs> all = userInfoEsRepository.findAll();
        return all;
    }
}
```

为了省事省略service层




## 四、 源码地址
https://github.com/yejingxuan/springboot-component-demo/tree/master/yjx-demo-es