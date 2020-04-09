

## 一、Selenium介绍
　　Selenium 是什么？一句话，自动化测试工具。它支持各种浏览器，包括 Chrome，Safari，Firefox 等主流界面式浏览器，如果你在这些浏览器里面安装一个 Selenium 的插件，那么便可以方便地实现Web界面的测试。换句话说叫 Selenium 支持这些浏览器驱动。Selenium支持多种语言开发，比如 Java，C，Ruby等等，有 Python 吗？那是必须的！安装只要cmd中 pip install selenium 即可。

## 二、PhantomJS介绍
对于一般网站来说scrapy、requests、beautifulsoup等都可以爬取，但是有些信息需要执行js才能显现，而且你肉眼所能看到的基本都能爬取下来，在学习中遇到了，就记录下来方便以后查看。

其中PhantomJS同时可以换成Chrome、Firefox、Ie等等，但是PhantomJS是一个无头的浏览器，运行是不会跳出相应的浏览器，运行相对效率较高。在调试中可以先换成Chrome，方便调试，最后再换成PhantomJS即可。

PhantomJS是一个基于webkit的JavaScript API。它使用QtWebKit作为它核心浏览器的功能，使用webkit来编译解释执行JavaScript代码。任何你可以在基于webkit浏览器做的事情，它都能做到。它不仅是个隐形的浏览器，提供了诸如CSS选择器、支持Web标准、DOM操作、JSON、HTML5、Canvas、SVG等，同时也提供了处理文件I/O的操作，从而使你可以向操作系统读写文件等。PhantomJS的用处可谓非常广泛，诸如前端无界面自动化测试（需要结合Jasmin）、网络监测、网页截屏等。



## 三、PhantomJS安装
本人windowns7系统，把下载下来的phantomjs.exe移到你所用python文件夹下的Script中就可以使用了。