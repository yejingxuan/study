# es学习手册

1.安装步骤


2.基本指令
* 访问es  http://47.98.48.32:9200/  

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

* 新增索引 

        PUT  http://47.98.48.32:9200/userinfo

* 基本查询  

        GET  http://47.98.48.32:9200/userinfo


