

- [一、准备工作](#%e4%b8%80%e5%87%86%e5%a4%87%e5%b7%a5%e4%bd%9c)
- [二、项目初始化](#%e4%ba%8c%e9%a1%b9%e7%9b%ae%e5%88%9d%e5%a7%8b%e5%8c%96)
- [三、组件引用](#%e4%b8%89%e7%bb%84%e4%bb%b6%e5%bc%95%e7%94%a8)

## 一、准备工作

1. 安装淘宝镜像
    ```
    npm  install  -g  cnpm  --registry=https://registry.npm.taobao.org
    ```

## 二、项目初始化
1. 安装vue-cli
    ```shell
    cnpm install vue-cli -g
    ```

2. 检测vue-cli是否安装成功
    ```shell
    vue list
    ```

3. 初始化项目
    ```shell
    vue init webpack  ”项目名称“
    ```
    ![](https://gitee.com/jingxuanye/yjx-pictures/raw/master/pic/20200305201713.png)

4. 安装依赖包
    ```shell
    cnpm install
    ```

5. 启动项目
   ```shell
   cnpm run dev
   ```


## 三、组件引用

1. element-ui引用
    ```shell
    cnpm i element-ui -S
    ```

    main.js引入element-ui
    ```js
    import ElementUI from 'element-ui'
    import 'element-ui/lib/theme-chalk/index.css'

    Vue.use(ElementUI)
    ```

    app.vue添加如下代码验证
    ```html
    <template>
        <div id="app">
            <img src="./assets/logo.png">
            <router-view/>
            <el-row>
                <el-button>默认按钮</el-button>
                <el-button type="primary">主要按钮</el-button>
                <el-button type="success">成功按钮</el-button>
                <el-button type="info">信息按钮</el-button>
                <el-button type="warning">警告按钮</el-button>
                <el-button type="danger">危险按钮</el-button>
            </el-row>
        </div>
    </template>

2.  vue-resource引用

    ```shell
    cnpm install vue-resource --save
    ```
    【main.js】引入
    ```
    // 引入vue-resource
    import VueResource from 'vue-resource'
    // 使用vue-resource
    Vue.use(VueResource)
    ```