# webmagic爬取天气预报

- [webmagic爬取天气预报](#webmagic%e7%88%ac%e5%8f%96%e5%a4%a9%e6%b0%94%e9%a2%84%e6%8a%a5)
  - [一、WebMagic简介](#%e4%b8%80webmagic%e7%ae%80%e4%bb%8b)
  - [二、基于springboot和WebMagic构建天气预报爬虫](#%e4%ba%8c%e5%9f%ba%e4%ba%8espringboot%e5%92%8cwebmagic%e6%9e%84%e5%bb%ba%e5%a4%a9%e6%b0%94%e9%a2%84%e6%8a%a5%e7%88%ac%e8%99%ab)
    - [2.1、springboot项目搭建，此处不再进行详述](#21springboot%e9%a1%b9%e7%9b%ae%e6%90%ad%e5%bb%ba%e6%ad%a4%e5%a4%84%e4%b8%8d%e5%86%8d%e8%bf%9b%e8%a1%8c%e8%af%a6%e8%bf%b0)
    - [2.2、webmagic的核心jar包引入](#22webmagic%e7%9a%84%e6%a0%b8%e5%bf%83jar%e5%8c%85%e5%bc%95%e5%85%a5)
    - [2.3、自定义PageProcessor来爬取天气预报](#23%e8%87%aa%e5%ae%9a%e4%b9%89pageprocessor%e6%9d%a5%e7%88%ac%e5%8f%96%e5%a4%a9%e6%b0%94%e9%a2%84%e6%8a%a5)
    - [2.4、创建定时任务来周期性爬取天气预报](#24%e5%88%9b%e5%bb%ba%e5%ae%9a%e6%97%b6%e4%bb%bb%e5%8a%a1%e6%9d%a5%e5%91%a8%e6%9c%9f%e6%80%a7%e7%88%ac%e5%8f%96%e5%a4%a9%e6%b0%94%e9%a2%84%e6%8a%a5)


## 一、WebMagic简介

> WebMagic是一款简单灵活的爬虫框架。基于它可以很容易的编写一个爬虫。

---

## 二、基于springboot和WebMagic构建天气预报爬虫

### 2.1、springboot项目搭建，此处不再进行详述


### 2.2、webmagic的核心jar包引入

```xml
<!--webmagic爬虫集成-->
<dependency>
    <groupId>us.codecraft</groupId>
    <artifactId>webmagic-core</artifactId>
    <version>0.7.3</version>
    <!--日志包冲突的话去掉webmagic里的日志包-->
    <exclusions>
        <exclusion>
            <groupId>org.slf4j</groupId>
            <artifactId>slf4j-log4j12</artifactId>
        </exclusion>
        <exclusion>
            <groupId>org.slf4j</groupId>
            <artifactId>slf4j-api</artifactId>
        </exclusion>
    </exclusions>
</dependency>

<dependency>
    <groupId>us.codecraft</groupId>
    <artifactId>webmagic-extension</artifactId>
    <version>0.7.3</version>
</dependency>
```

- webmagic-core是WebMagic核心部分，只包含爬虫基本模块和基本抽取器。WebMagic-core的目标是成为网页爬虫的一个教科书般的实现。

- webmagic-extension是WebMagic的主要扩展模块，提供一些更方便的编写爬虫的工具。包括注解格式定义爬虫、JSON、分布式等支持。




### 2.3、自定义PageProcessor来爬取天气预报

- WebMagic的四个组件

1. Downloader：负责从互联网上下载页面，以便后续处理。WebMagic默认使用了Apache HttpClient作为下载工具。

2. PageProcessor：负责解析页面，抽取有用信息，以及发现新的链接。WebMagic使用Jsoup作为HTML解析工具，并基于其开发了解析XPath的工具Xsoup。

3. Scheduler：负责管理待抓取的URL，以及一些去重的工作。WebMagic默认提供了JDK的内存队列来管理URL，并用集合来进行去重。也支持使用Redis进行分布式管理。除非项目有一些特殊的分布式需求，否则无需自己定制Scheduler。

4. Pipeline：负责抽取结果的处理，包括计算、持久化到文件、数据库等。WebMagic默认提供了“输出到控制台”和“保存到文件”两种结果处理方案。Pipeline定义了结果保存的方式，如果你要保存到指定数据库，则需要编写对应的Pipeline。对于一类需求一般只需编写一个Pipeline。

- 想要快速完成一个爬虫，一般只需要实现PageProcessor组件即可。

```java
@Slf4j
@Service
public class WeatherProcessor implements PageProcessor {

    //爬虫网站的根路径
    public static final String list = "http://www.weather.com.cn";

    private Site site = Site.me()
            .setDomain("http://www.weather.com.cn")
            .setSleepTime(3000)
            .setTimeOut(100000)
            .setUserAgent(AgentUtil.Mac_AGENT);


    @Override
    public void process(Page page) {
        //判断是否爬取的目标网站
        if (page.getUrl().regex(list).match()) {
            //使用xpath对页面元素进行提取
            List<Selectable> weatherList = page.getHtml().xpath("//ul[@class='t clearfix']/li").nodes();
            //解析出当前天气的归属城市
            String cityName = page.getHtml().xpath("//a[@href='"+page.getUrl()+"']/text()").toString();
            if(weatherList.size()>1){
                Selectable today =  weatherList.get(0);
                Selectable tomorrow =  weatherList.get(1);
                this.parseWeather(today, cityName);
                this.parseWeather(tomorrow, cityName);
            }
        }
    }

    @Override
    public Site getSite() {
        return this.site;
    }

    private void parseWeather(Selectable item, String cityName) {
        //天气内容
        String wConent = item.xpath("//[@class='wea']/text()").toString();
        //最高气温
        String wMaxC = item.xpath("//[@class='tem']/span/text()").toString().replaceAll("℃","");
        //最低气温
        String wMinC = item.xpath("//[@class='tem']/i/text()").toString().replaceAll("℃","");

        log.info(cityName+wConent+wMaxC+wMinC, e);
    }

```


### 2.4、创建定时任务来周期性爬取天气预报

通过Spider.create()方法来创建爬虫，通过run()方法来启动爬虫。  
注意在springboot启动类上加上@EnableScheduling注解，开启定时任务。


```java
@Component
@Slf4j
public class SpiderWeatherSchedule {

    @Autowired
    private WeatherProcessor weatherProcessor;

    //每天20点开始进行爬取
    @Scheduled(cron = "0 0 20 * * ?}")
    public void SpiderWeather() {
        log.info("==============" + System.currentTimeMillis() + "开启爬取武汉天气预报");
        Spider.create(weatherProcessor)
                .addUrl("http://www.weather.com.cn/weather/101200101.shtml").thread(1)
                .run();
        log.info("==============" + System.currentTimeMillis() + "开启爬取深圳天气预报");
        Spider.create(weatherProcessor)
                .addUrl("http://www.weather.com.cn/weather/101280601.shtml").thread(1)
                .run();
    }

}
```



> 参考文章：[webmagic中文文档](http://webmagic.io/docs/zh/)