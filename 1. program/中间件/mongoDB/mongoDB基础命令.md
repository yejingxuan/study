# MongoDB基础命令

## 1、创建命令

- 创建库
    ```
    use dbname
    ```

- 创建表

    ```js
    db.createCollection("table_name")
    ```

- 创建索引
    ```js
    db.getCollection('table_name').createIndex({"colum1":1})
    ```

---

## 2、查询命令

> 参考文章：[mongo数据库的各种查询语句示例](https://blog.csdn.net/qq_27093465/article/details/51700435)

- 查询表数据
    ```js
    //select * from table_name
    db.getCollection('table_name').find({})
    //select colum3, colum4 from table_name where colum1='a' and colum2 = 'b'
    db.getCollection('table_name').find({"colum1":"a","colum2":"b"},{"colum3":1,"colum4":1})
    ```

- 排序(1表示正序，-1表示倒序)
    ```js
    //select * from table_name order by colum1 asc
    db.getCollection('table_name').find({}).sort({"colum1":1})
    ```


- 分组聚合
    ```js
    db.getCollection('table_name').aggregate([{$group : {_id : "$colum1", num_tutorial : {$sum : 1}}}])
    ```

- 分页(skip:跳过多少行数据；limit:查询多少数据[pagesize])
    ```js
    db.getCollection('table_name').find({}).sort({"colum1":-1}).skip(90).limit(10)
    ```

- 去重
    ```js
    //根据column1去重
    db.tablename.distinct('column1')
    //筛选出column2 = column2_value的数据，根据column1去重
    db.tablename.distinct('column1',{"column2" : "column2_value"})
    ```


----

## 3、修改命令




## 4、删除命令

- 删除表
    ```js
    db.getCollection('table_name').drop()
    ```

- 删除表数据
    ```js
    db.getCollection('table_name').remove({ '' : '' });
    ```