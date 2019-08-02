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
    #-*- coding: utf-8 -*-
    ```