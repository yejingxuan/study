# 常用类库
- [常用类库](#常用类库)
  - [一、xml操作](#一xml操作)
  - [二、深拷贝、浅拷贝](#二深拷贝浅拷贝)


## 一、xml操作

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

## 二、深拷贝、浅拷贝

```python
dict1 = {'user': 'test', 'num': [1, 2, 3]}      # 原字典
dict2 = dict1                                   # 直接赋值
dict3 = dict1.copy()                            # 浅拷贝，只深拷贝父级目录
dict4 = copy.deepcopy(dict1)                    # 深拷贝拷贝，父级目录，子级目录全部拷贝（需导入copy模块）
```
