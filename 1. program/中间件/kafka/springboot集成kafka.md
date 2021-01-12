# springboot集成kafka

[TOC]


## 一、所需jar包

此处springboot版本为2.1.0.RELEASE，不同的springboot版本需要引用不同版本的spring-kafka版本，此处需注意。

```xml
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-web</artifactId>
</dependency>

<dependency>
    <groupId>org.springframework.kafka</groupId>
    <artifactId>spring-kafka</artifactId>
    <version>2.2.0.RELEASE</version>
</dependency>
```


## 二、 kafka配置

### 2.1、配置文件配置

```properties
#kafka相关配置
spring.kafka.bootstrap-servers=127.0.0.1:9092

#指定生产者key-value序列化反序列化
spring.kafka.producer.key-serializer=org.apache.kafka.common.serialization.StringSerializer
spring.kafka.producer.value-serializer=org.apache.kafka.common.serialization.StringSerializer
spring.kafka.producer.batch-size=65536
spring.kafka.producer.buffer-memory=524288
#生产者发送的消息大小限制
spring.kafka.producer.properties.max.request.size=4096000
spring.kafka.producer.properties.buffer.memory=4096000

spring.kafka.consumer.group-id=0
#指定消费者key-value序列化反序列化
spring.kafka.consumer.key-deserializer=org.apache.kafka.common.serialization.StringDeserializer
spring.kafka.consumer.value-deserializer=org.apache.kafka.common.serialization.StringDeserializer
#轮询消费者时使用的超时（以毫秒为单位）
spring.kafka.listener.poll-timeout=30000

#topic的默认分片数量
spring.kafak.topic.numPartitions=${TOPIC_NUMPARTITIONS:1}
#topic的副本数量
spring.kafak.topic.replicationFactor=${TOPIC_REPLICATIONFACTOR:1}
```

### 2.2、配置类
```java
@Configuration
public class KafkaConfig {

    @Autowired
    private ApplicationConfig config;

    @Bean
    public KafkaAdmin kafkaAdmin() {
        Map<String, Object> props = new HashMap<>();
        //配置Kafka实例的连接地址
        props.put(AdminClientConfig.BOOTSTRAP_SERVERS_CONFIG, config.getKafkaServer());
        KafkaAdmin admin = new KafkaAdmin(props);
        return admin;
    }

    @Bean
    public AdminClient adminClient() {
        return AdminClient.create(kafkaAdmin().getConfig());
    }
}
```


## 三、 代码集成

### 3.1、创建topic
```java
@Autowired
private KafkaTemplate kafkaTemplate;

@Autowired
private KafkaConfig kafkaConfig;

@ApiOperation(value = "创建topic")
@PutMapping(value = "createTopic")
public String createTopic(@RequestParam(value = "topicName") String topicName){
    List<NewTopic> topics = new ArrayList();
    NewTopic topic = new NewTopic(topicName, 1, Short.valueOf("1"));
    Map<String, String> configs = new HashMap<>();
    //设置topic能接收的最大消息体，默认为一兆
    configs.put("max.message.bytes","4096000");
    topic.configs(configs);

    topics.add(topic);
    kafkaConfig.adminClient().createTopics(topics);
    return "success";
}
```

### 3.2、发送kafka消息
```java
@ApiOperation(value = "发送kafka消息")
@PutMapping(value = "sendMsg")
public String sendMsg(UserInfo userInfo, @RequestParam(value = "topicName") String topicName){
    String msg = JSONObject.toJSONString(userInfo);
    log.info(String.valueOf(msg.length()));
    kafkaTemplate.send(topicName, msg);
    return "success";
}
```


### 3.3、接收kafka消息
```java
@Slf4j
@Component
public class KafkaReceive {
    @KafkaListener(topics = {"yjx-user"})
    public void listenUser(ConsumerRecord message){
        log.info(message.value().toString());
    }
}
```



## 四、 源码地址
https://github.com/yejingxuan/springboot-component-demo/tree/master/yjx-demo-kafka