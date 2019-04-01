# es学习手册

## 1、安装步骤


## 2、基本指令
* 访问es  http://127.0.0.1:9200/  

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
            "fields" : ["_all"]
        }
    },
    "size": 10,
    "from": 0,
    "highlight": {
        "fields" : {
            "name" : {}
        }
    }
}
```


