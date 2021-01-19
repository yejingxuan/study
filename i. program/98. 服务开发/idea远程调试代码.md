# idea远程调试java代码

## 1. 服务器端项目启动时加上参数

```
-jar -Xdebug -Xrunjdwp:server=y,transport=dt_socket,address=指定端口,suspend=n 
```


## 2. 在idea端进行远程配置

- edit configurations
    - 设置host和port，host为服务器ip , port为服务器端配置的指定端口，
    - 执行debug