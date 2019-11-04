# python常用库安装

[TOC]

## 一、软件安装源替换

pip国内的一些镜像

- 阿里云 http://mirrors.aliyun.com/pypi/simple/ 
- 中国科技大学 https://pypi.mirrors.ustc.edu.cn/simple/ 
- 豆瓣(douban) http://pypi.douban.com/simple/ 
- 清华大学 https://pypi.tuna.tsinghua.edu.cn/simple/ 
- 中国科学技术大学 http://pypi.mirrors.ustc.edu.cn/simple/

临时替换方法：可以在使用pip的时候在后面加上-i参数，指定pip源，并信任该站点

```shell
pip install scrapy -i http://mirrors.aliyun.com/pypi/simple/  --trusted-host  mirrors.aliyun.com
```

永久修改： 
linux: 修改 ~/.pip/pip.conf (没有就创建一个)， 内容如下：
```shell
index-url = http://mirrors.aliyun.com/pypi/simple/
```
windows: 直接在user目录中创建一个pip目录，如：C:\Users\xx\pip，新建文件pip.ini，内容如下
```shell
index-url = http://mirrors.aliyun.com/pypi/simple/
```

## 二、常用库安装

- pip更新

    ```shell
    python -m pip install --upgrade pip
    ```


- pymysql ：mysql数据库

    ```shell
    pip install pymysql
    ```


- DBUtils ：数据库连接池

    ```shell
    pip install DBUtils
    ```


- beautifulsoup4 ：常用爬虫类库

    ```shell
    pip install beautifulsoup4
    ```


- openpyxl：常用excel处理类库

    ```shell
    pip install openpyxl
    ```


- pyautogui：常用gui自动化处理工具

    ```shell
    pip install pyautogui
    ```