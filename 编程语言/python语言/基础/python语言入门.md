# python语言入门

## 一、编程规约

### 1.1、命名规范

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

### 1.2、建议

- 每个模板的第一行声明编码字符集

    ```python
    #!/usr/bin/env python
    # encoding=utf-8
    ```



## 二、常用类库

### xml操作

```python
def createPositiveRepXml(jsonResult):
    # 在内存中创建一个空的文档
    doc = xml.dom.minidom.Document()
    # 创建一个根节点Managers对象
    root = doc.createElement('GeocodeResponse')

    nodeManagerStatus = doc.createElement('status')
    nodeManagerStatus.appendChild(doc.createTextNode("true"))

    nodeManagerCount = doc.createElement('count')
    nodeManagerCount.appendChild(doc.createTextNode("10"))

    root.appendChild(nodeManagerStatus)
    root.appendChild(nodeManagerCount)

    # 将根节点添加到文档对象中
    doc.appendChild(root)

    return doc

# 实现一个重定向的类，该类有write方法
class XmlStdin():
    def __init__(self):
        self.str = ""

    def write(self, value):
        self.str += value

    def toString(self):
        return self.str
    
# 写入本地xml文档
fp = open('Manager.xml', 'w')
doc.writexml(fp, indent='\t', addindent='\t', newl='\n', encoding="utf-8")
print(doc)

# 输出xml字符串
def DoXmlDomToStr(xmlDom):
    # 修改标准输出流
    xmlStdin = XmlStdin()
    sys.stdin = xmlStdin
    xmlDom.writexml(sys.stdin, addindent='\t', newl='\n', encoding='utf-8')
    return sys.stdin.toString()
```

### 深拷贝、浅拷贝

```python
dict1 = {'user': 'test', 'num': [1, 2, 3]}                  #  原字典
dict2 = dict1                                                       # 直接赋值
dict3 = dict1.copy()                                            # 浅拷贝，只深拷贝父级目录
dict4 = copy.deepcopy(dict1)                             # 深拷贝拷贝，父级目录，子级目录全部拷贝（需导入copy模块）
```



### 三、打包部署

```shell
pip3 freeze >requirements.txt
pip3 install -r requirements.txt

# 导出依赖包到当前目录
pip download -r requirements.txt
# 导入当前目前的依赖包到项目里
pip install   --no-index   --find-links=.   -r   requirements.txt
```

