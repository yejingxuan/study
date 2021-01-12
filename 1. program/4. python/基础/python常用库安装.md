# python常用库安装

- [python常用库安装](#python%e5%b8%b8%e7%94%a8%e5%ba%93%e5%ae%89%e8%a3%85)
  - [一、软件安装源替换](#%e4%b8%80%e8%bd%af%e4%bb%b6%e5%ae%89%e8%a3%85%e6%ba%90%e6%9b%bf%e6%8d%a2)
    - [pip国内的一些镜像](#pip%e5%9b%bd%e5%86%85%e7%9a%84%e4%b8%80%e4%ba%9b%e9%95%9c%e5%83%8f)
    - [临时替换方法](#%e4%b8%b4%e6%97%b6%e6%9b%bf%e6%8d%a2%e6%96%b9%e6%b3%95)
    - [永久修改](#%e6%b0%b8%e4%b9%85%e4%bf%ae%e6%94%b9)
  - [二、常用库安装](#%e4%ba%8c%e5%b8%b8%e7%94%a8%e5%ba%93%e5%ae%89%e8%a3%85)

## 一、软件安装源替换

### pip国内的一些镜像

- 阿里云 http://mirrors.aliyun.com/pypi/simple/ 
- 中国科技大学 https://pypi.mirrors.ustc.edu.cn/simple/ 
- 豆瓣(douban) http://pypi.douban.com/simple/ 
- 清华大学 https://pypi.tuna.tsinghua.edu.cn/simple/ 
- 中国科学技术大学 http://pypi.mirrors.ustc.edu.cn/simple/

### 临时替换方法

可以在使用pip的时候在后面加上-i参数，指定pip源，并信任该站点

```shell
pip install scrapy -i http://mirrors.aliyun.com/pypi/simple/  --trusted-host  mirrors.aliyun.com
```

### 永久修改
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


- pymysql ：mysql数据库操作

    ```shell
    pip install pymysql
    ```


- DBUtils ：数据库连接池

    ```shell
    pip install DBUtils
    ```


- beautifulsoup4 ：常用爬虫网页解析类库

    ```shell
    pip install beautifulsoup4
    ```

- requests ：http请求库

    ```shell
    pip install requests
    ```

- openpyxl：常用excel处理类库

    ```shell
    pip install openpyxl
    ```


- pyautogui：常用gui自动化处理工具

    ```shell
    pip install pyautogui
    ```

- selenium：自动化测试工具

    ```shell
    pip install selenium
    ```

- flask：轻量级web框架
    ```shell
    pip install flask
    ```

> ps: 持续更新