# es学习手册

## 1、安装步骤


## 2、基本指令
* 访问es  http://127.0.0.1:9200/  
```json
{
"name" : "BvAiR-X",
"cluster_name" : "elasticsearch",
"cluster_uuid" : "LorchJMjTOKvvR2fv-H8oA",
"version" : {
        "number" : "5.5.0",
        "build_hash" : "260387d",
        "build_date" : "2017-06-30T23:16:05.735Z",
        "build_snapshot" : false,
        "lucene_version" : "6.6.0"
},
"tagline" : "You Know, for Search"
}
```

## 3、新增索引

```js
PUT  127.0.0.1:9200/userinfo
```


## 4、基本查询  

### 参数放于url中

```js
//查询总数
127.0.0.1:9200/userinfo/_count

//查询全部数据
127.0.0.1:9200/userinfo/_search

//搜索name里有叶菁烜的数据
127.0.0.1:9200/picture/_search?q=name:叶菁烜
```

### 参数放于请求体（body）里

```js
127.0.0.1:9200/picture/_search

{
    "query": {
        "multi_match" : {
            "query" : "张三",
            "fields" : ["name", "addr"]
        }
    },
    "size": 10,
    "from": 0,
    "highlight": {
    	"pre_tags": ["<tag1>"],
    	"post_tags": ["</tag2>"],
        "fields": {
            "name": {}
        }
    }
}
```

- query 里表示查询条件
    - "fields" : ["_all"] 表示查询所有字段

- size 表示分页的pagesize

- from 表示从第几行开始查询

- highlight 表示高亮字段


