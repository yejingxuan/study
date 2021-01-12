# Django开发指南

- [Django开发指南](#django开发指南)
  - [一、Django是什么](#一django是什么)
  - [二、Django 特点](#二django-特点)
  - [三、Django 模型](#三django-模型)
  - [四、入门](#四入门)
    - [4.1、安装django](#41安装django)
    - [4.2、初始化django项目](#42初始化django项目)
    - [4.3、启动django项目](#43启动django项目)
  - [参考文章](#参考文章)

## 一、Django是什么
Python下有许多不同的 Web 框架。Django是重量级选手中最有代表性的一位。许多成功的网站和APP都基于Django。
Django是一个开放源代码的Web应用框架，由Python写成。
Django遵守BSD版权，初次发布于2005年7月, 并于2008年9月发布了第一个正式版本1.0 。
Django采用了MVC的软件设计模式，即模型M，视图V和控制器C。
## 二、Django 特点
强大的数据库功能
拥有强大的数据库操作接口（QuerySet API），如需要也能执行原生SQL。

自带强大后台
几行简单的代码就让你的网站拥有一个强大的后台，轻松管理内容！

优雅的网址
用正则匹配网址，传递到对应函数，随意定义，如你所想

模板系统
强大，易扩展的模板系统，设计简易，代码，样式分开设计，更容易管理。

**注：**前后端分离时，也可以用Django开发API，完全不用模板系统。

缓存系统
与Memcached, Redis等缓存系统联用，更出色的表现，更快的加载速度

国际化
完全支持多语言应用，允许你定义翻译的字符，轻松翻译成不同国家的语言。

## 三、Django 模型
Django 对各种数据库提供了很好的支持，包括：PostgreSQL、MySQL、SQLite、Oracle。
Django 为这些数据库提供了统一的调用API。 我们可以根据自己业务需求选择不同的数据库。
MySQL 是 Web 应用中最常用的数据库



## 四、入门

### 4.1、安装django

- 手动安装：
  - 下载 Django 压缩包，解压并和 Python安装目录放在同一个根目录，进入 Django 目录，执行 python setup.py install，然后开始安装，Django 将要被安装到 Python 的 Lib下site-packages。
- 自动安装：

```shell
pip3 install django==3.1.2
```

成功安装Django后，在下图中的路径可找到`django-admin.exe`文件，将它加入操作系统环境变量中。这样以后调用会比较方便。

### 4.2、初始化django项目

```shell
# 初始化一个django项目
django-admin.py startproject projectname

# 添加新的APP
python manage.py startapp appname
```

### 4.3、启动django项目

```shell
python manage.py runserver
```









## 参考文章

-  [Django操作mysql数据库增删改查](https://blog.csdn.net/zhangcongyi420/article/details/102313888)