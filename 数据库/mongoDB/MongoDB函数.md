# MongoDB函数

## 1、mongodb数据库游标的使用

```js
//使用find()返回一个游标：
var cursor = db.XXX.find();
while (cursor.hasNext()) {
  obj = cursor.next();
  print(obj.name);
}

//使用游标的forEach()循环遍历:
cursor.forEach(function(x) {
  print(x.name);
});
```