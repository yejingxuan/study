# springboot2集成elasticsearch7

- [springboot2集成elasticsearch7](#springboot2%e9%9b%86%e6%88%90elasticsearch7)
  - [一、elasticsearch7搭建](#%e4%b8%80elasticsearch7%e6%90%ad%e5%bb%ba)
    - [1.1、docker-compose编排文件](#11docker-compose%e7%bc%96%e6%8e%92%e6%96%87%e4%bb%b6)
  - [二、springboot2.X集成es7](#%e4%ba%8cspringboot2x%e9%9b%86%e6%88%90es7)
    - [2.1、jar包集成](#21jar%e5%8c%85%e9%9b%86%e6%88%90)
    - [2.2、代码集成](#22%e4%bb%a3%e7%a0%81%e9%9b%86%e6%88%90)
    - [2.3、常用API](#23%e5%b8%b8%e7%94%a8api)
      - [创建索引](#%e5%88%9b%e5%bb%ba%e7%b4%a2%e5%bc%95)
      - [检测索引是否存在](#%e6%a3%80%e6%b5%8b%e7%b4%a2%e5%bc%95%e6%98%af%e5%90%a6%e5%ad%98%e5%9c%a8)
      - [删除索引](#%e5%88%a0%e9%99%a4%e7%b4%a2%e5%bc%95)
      - [添加数据](#%e6%b7%bb%e5%8a%a0%e6%95%b0%e6%8d%ae)
      - [查询数据](#%e6%9f%a5%e8%af%a2%e6%95%b0%e6%8d%ae)

## 一、elasticsearch7搭建

### 1.1、docker-compose编排文件
```yml
version: '3'
services:
  yjx-elasticsearch7:
    image: docker.io/elasticsearch:7.0.1
    ports:
      - "9200:9200"
      - "9300:9300"
    volumes:
      - /home/jxye/docker-servers/elasticsearch7/plugins:/usr/share/elasticsearch/plugins
      - /home/jxye/docker-servers/elasticsearch7/data:/usr/share/elasticsearch/data
    environment:
      ES_JAVA_OPTS: "-Xms256m -Xmx256m"
      discovery.type: "single-node"
    container_name: yjx-elasticsearch7
```

一共挂载了两个目录
data目录 ：持久化es的数据，外挂到系统盘中
plugins目录 ：外挂es的插件到系统盘中，方便灵活配置。例如IK分词插件


## 二、springboot2.X集成es7


### 2.1、jar包集成
在es7以后，spring-data已经不再支持，es官方也推荐我们使用官方的java API——elasticsearch-rest-high-level-client
```xml
<!--集成es7-->
<dependency>
    <groupId>org.elasticsearch.client</groupId>
    <artifactId>elasticsearch-rest-high-level-client</artifactId>
    <version>7.2.0</version>
</dependency>

<dependency>
    <groupId>org.elasticsearch</groupId>
    <artifactId>elasticsearch</artifactId>
    <version>7.2.0</version>
</dependency>
```


### 2.2、代码集成

配置文件
```properties
elasticsearch.host=47.98.48.32
elasticsearch.port=9200
elasticsearch.clustername=docker-cluster
```


代码配置
```java
@Configuration
public class ElasticSearchSevenConfig {

    @Autowired
    private ApplicationConfig config;

    private static final String HTTP_SCHEME = "http";

    @Bean
    public RestClientBuilder restClientBuilder() {
        HttpHost hosts = new HttpHost(config.getEsHost(), config.getEsPort(), HTTP_SCHEME);
        return RestClient.builder(hosts);
    }

    @Bean(name = "highLevelClient")
    public RestHighLevelClient highLevelClient(RestClientBuilder restClientBuilder) {
        return new RestHighLevelClient(restClientBuilder);
    }

}
```

使用es,注入highLevelClient
```java
@Autowired
private RestHighLevelClient highLevelClient;
```



### 2.3、常用API

#### 创建索引

```java
public String createEmployIndex() {
    try {
        //初始化mapping
        Resource resource = new ClassPathResource("es_mappings/employ_mapping.json");
        byte[] template = IOUtils.toByteArray(resource.getInputStream());
        String index = new String(template);

        //初始化索引
        String indexName = "employ_index";
        CreateIndexRequest indexRequest = new CreateIndexRequest(indexName)
                .source(index, XContentType.JSON);

        //创建索引
        highLevelClient.indices().create(indexRequest, RequestOptions.DEFAULT);
    } catch (Exception e) {
        log.error("create error", e);
        return "error";
    }
    return "success";
}
```

mapping文件，放置于resource/es_mappings目录下
analyzer、search_analyzer配置IK分词器

```json
{
  "properties": {
    "title": {
      "type": "text",
      "analyzer": "ik_max_word",
      "search_analyzer": "ik_smart"
    },
    "href": {
      "type": "text",
      "analyzer": "ik_max_word",
      "search_analyzer": "ik_smart"
    },
    "deadlines": {
      "type": "keyword"
    },
    "header_count": {
      "type": "keyword"
    }
  }
}
```

#### 检测索引是否存在

```java
public Boolean checkIndex(String indexName) {
    GetIndexRequest request = new GetIndexRequest(indexName);
    Boolean res = false;
    try {
        res = highLevelClient.indices().exists(request, RequestOptions.DEFAULT);
    } catch (IOException e) {
        log.error("checkIndex异常", e);
    }
    return res;
}
```


#### 删除索引

```java
public Boolean deleteIndex(String indexName) {
    DeleteIndexRequest request = new DeleteIndexRequest(indexName);
    Boolean res = false;
    try {
        AcknowledgedResponse rep = highLevelClient.indices()
                .delete(request, RequestOptions.DEFAULT);
        res = rep.isAcknowledged();
    } catch (IOException e) {
        log.error("deleteIndex异常", e);
    }
    return res;
}
```

#### 添加数据

```java
public Boolean addEmployDocuments() {
    Boolean res = false;
    String indexName = "employ_index";
    try {
        EmployEs info1 = new EmployEs("华中农业大学2020年度辅导员招聘公告（博士学位辅导员实行年薪制）",
                "http://www.gaoxiaojob.com/zhaopin/fudaoyuan/20191209/415572.html",
                "2019年12月24日", "若干");
        EmployEs info2 = new EmployEs("武昌工学院2020年各类教师引进计划",
                "http://www.gaoxiaojob.com/zhaopin/gaoxiaojiaoshi/20191206/415305.html",
        "详见正文", "若干");
        String item1 = JSONObject.toJSONString(info1);
        String item2 = JSONObject.toJSONString(info2);

        BulkRequest bulkRequest = new BulkRequest();
        bulkRequest.add(new IndexRequest(indexName).source(item1, XContentType.JSON));
        bulkRequest.add(new IndexRequest(indexName).source(item2, XContentType.JSON));

        BulkResponse bulk = highLevelClient.bulk(bulkRequest, RequestOptions.DEFAULT);
        res = bulk.hasFailures();
    } catch (IOException e) {
        log.error("addDocuments异常", e);
    }
    return !res;
}
```

#### 查询数据

```java
public SearchHits searchEmployDocuments(String title) {
    String indexName = "employ_index";

    SearchSourceBuilder searchSourceBuilder = new SearchSourceBuilder();
    //searchSourceBuilder.query(QueryBuilders.matchAllQuery());
    searchSourceBuilder.query(QueryBuilders.matchQuery("title", title));
    searchSourceBuilder.from(0);
    searchSourceBuilder.size(10);

    HighlightBuilder highlightBuilder = new HighlightBuilder();
    highlightBuilder.preTags("<tag1>");
    highlightBuilder.postTags("</tag1>");
    highlightBuilder.field("title");
    searchSourceBuilder.highlighter(highlightBuilder);

    SearchRequest searchRequest = new SearchRequest().indices(indexName).source(searchSourceBuilder);
    SearchHits hits = null;
    try {
        SearchResponse search = highLevelClient.search(searchRequest, RequestOptions.DEFAULT);
        hits = search.getHits();
        log.info(hits.getAt(0).getSourceAsMap().toString());
        log.info(hits.getAt(0).getHighlightFields().toString());
    } catch (IOException e) {
        e.printStackTrace();
    }

    return hits;

}
```