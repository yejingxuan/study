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