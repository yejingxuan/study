# MongoDB基础命令

## 1、创建命令

- 创建表
```
use dbname
```

- 创建表

```js
db.createCollection("table_name")
```

---

## 2、查询命令

- 查询表数据
```js
db.getCollection('table_name').find({})

db.getCollection('table_name').find({"colum1":""},{"colum2":""})
```

- 排序(1表示正序，-1表示倒序)
```js
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
db.tablename.distinct('column1')
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