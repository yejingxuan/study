# 设计模式

## 一、创建型




## 二、行为型

### 责任链模式（Chain Of Responsibility）


### 迭代器模式（Iterator）
- 目的
  - 使调用者无需关心容器内部具体是链表还是数组之类的实现，通过调用容器的createIterator()方法获取一个迭代器，迭代器里实现容器的遍历操作，通过迭代器的hasNext()方法和next()方法，来对容器进行遍历。

- 实现
  ```java
  public interface Iterator<Object> {
      Item next();
      boolean hasNext();
  }

  while (iterator.hasNext()) {
      System.out.println(iterator.next());
  }
  ```

## 三、结构型