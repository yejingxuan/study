# springcloud入门

## 一. springcloud服务注册中心搭建


## 二. springcloud服务提供者构建



## 三. springcloud服务消费者构建

### 3.1 springcloud提供两种服务调用方法ribbon和feign

    Ribbon和Feign都是用于调用其他服务的，不过方式不同。

    1.启动类使用的注解不同，Ribbon用的是@RibbonClient，Feign用的是@EnableFeignClients。

    2.服务的指定位置不同，Ribbon是在@RibbonClient注解上声明，Feign则是在定义抽象方法的接口中使用@FeignClient声明。

    3.调用方式不同，Ribbon需要自己构建http请求，模拟http请求然后使用RestTemplate发送给其他服务，步骤相当繁琐。
    Feign则是在Ribbon的基础上进行了一次改进，采用接口的方式，将需要调用的其他服务的方法定义成抽象方法即可，不需要自己构建http请求。不过要注意的是抽象方法的注解、方法签名要和提供服务的方法完全一致。

从实践上看，采用feign的方式更优雅（feign内部也使用了ribbon做负载均衡）。