# python快速开始

- [python快速开始](#python快速开始)
  - [一、快速开始](#一快速开始)
  - [二、python类库](#二python类库)
    - [1、软件安装源替换](#1软件安装源替换)
    - [2、常用库安装](#2常用库安装)
  - [三、虚拟环境](#三虚拟环境)
  - [四、打包部署](#四打包部署)
  - [五、编程规约](#五编程规约)
    - [1、命名规范](#1命名规范)
    - [2、建议](#2建议)

## 一、快速开始


## 二、python类库
### 1、软件安装源替换

- __pip国内的一些镜像__
  - 阿里云 http://mirrors.aliyun.com/pypi/simple/ 
  - 中国科技大学 https://pypi.mirrors.ustc.edu.cn/simple/ 
  - 豆瓣(douban) http://pypi.douban.com/simple/ 
  - 清华大学 https://pypi.tuna.tsinghua.edu.cn/simple/ 
  - 中国科学技术大学 http://pypi.mirrors.ustc.edu.cn/simple/

- __临时替换方法__
  可以在使用pip的时候在后面加上-i参数，指定pip源，并信任该站点

  ```shell
  pip install scrapy -i http://mirrors.aliyun.com/pypi/simple/  --trusted-host  mirrors.aliyun.com
  ```

- __永久修改__  
  - linux: 修改 ~/.pip/pip.conf (没有就创建一个)， 内容如下：
  ```shell
  index-url = http://mirrors.aliyun.com/pypi/simple/
  ```
  - windows: 直接在user目录中创建一个pip目录，如：C:\Users\xx\pip，新建文件pip.ini，内容如下
  ```shell
  index-url = http://mirrors.aliyun.com/pypi/simple/
  ```

### 2、常用库安装

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





## 三、虚拟环境

```python
# 创建虚拟环境
python -m venv name

# 

```



## 四、打包部署

```shell
pip3 freeze >requirements.txt
pip3 install -r requirements.txt

# 导出依赖包到当前目录
pip download -r requirements.txt
# 导入当前目前的依赖包到项目里
pip install   --no-index   --find-links=. -r requirements.txt
```


## 五、编程规约
### 1、命名规范
- 模块尽量使用小写命名，首字母保持小写，尽量不要用下划线(除非多个单词，且数量不多的情况)

    ```python
    # 正确的模块名
    import decoder
    import html_parser
    
    # 不推荐的模块名
    import Decoder
    ```
- 类名使用驼峰(CamelCase)命名风格，首字母大写，私有类可用一个下划线开头

    ```python
    class Farm():
        pass
    
    class AnimalFarm(Farm):
        pass
    
    class _PrivateFarm(Farm):
        pass
    ```
- 函数名一律小写，如有多个单词，用下划线隔开,私有函数在函数前加一个下划线_
    ```python
    def run():
    pass
 
    def run_with_env():
        pass
    
    def _private_func():
        pass
    ```
- 变量名尽量小写, 如有多个单词，用下划线隔开。常量采用全大写，如有多个单词，使用下划线隔开

    ```python
    if __name__ == '__main__':
        count = 0
        school_name = ''
    
    MAX_CLIENT = 100
    MAX_CONNECTION = 1000
    CONNECTION_TIMEOUT = 600
    ```

### 2、建议

- 每个模板的第一行声明编码字符集

    ```python
    #!/usr/bin/env python
    # encoding=utf-8
    ```
