# mall
### 介绍
mall 是一套商城系统，基于 golang、 gin、 gorm、 vue3、element plus、 vant weapp 开发，包括 golang服务端、 Vue后台管理员前端、微信小程序用户前端，主要功能有商品管理、订单管理、用户管理、商品浏览、收藏加购、地址管理、订单查询等。


## 本仓库停止更新！！！

## 最新版项目地址：https://github.com/zchengo/imall

### 技术选型

前端技术：

| 技术 | 说明 | 相关文档 |
|---|---|---|
| vue3 | 前端框架 | https://v3.cn.vuejs.org |
| vue-router | 页面路由 | https://next.router.vuejs.org |
| axios | 网络请求库 | https://axios-http.com |
| vuex | 状态管理 | https://next.vuex.vuejs.org |
| element plus | 前端UI组件库 | https://element-plus.org |

后端技术：

| 技术 | 说明 | 相关文档 |
|---|---|---|
| gin | Web框架 | https://gin-gonic.com |
| gorm | ORM框架 | https://gorm.io |
| jwt | 用户认证 | https://github.com/golang-jwt/jwt |
| captcha | 验证码生成器 | https://github.com/mojocn/base64Captcha |
| viper | 配置管理 | https://github.com/spf13/viper |
| redis | 缓存 | https://github.com/go-redis/redis |
| elasticsearch | 搜索引擎 | https://github.com/olivere/elastic |

微信小程序技术：

| 技术 | 说明 | 相关文档 |
|---|---|---|
| vant weapp | UI组件库 | https://vant-contrib.gitee.io/vant-weapp |
| 微信小程序开发文档 | 官方文档 | https://developers.weixin.qq.com/miniprogram/dev/framework |


### 运行环境

| 环境 | 版本 |
|---|---|
| go | 1.17.x |
| mysql | 8.0.x |
| redis | 6.0.x |
| elasticsearch | 7.14.x |
| node.js | 14.13.x |
| npm | 6.14.x |

### 安装与启动

#### 项目下载

使用git下载项目到本地：
```
git clone https://github.com/zchengo/mall.git
```

#### 相关资源文件

数据库文件在 mall/demo/sql 目录中，推荐使用 Navicat 来运行SQL文件。图片文件在 mall/demo/image 目录中，请在自己的电脑中新建一个目录，用于存放这些图片。


#### 安装一：Go服务端（server）

修改配置文件：

配置文件位于 /server/config.yaml，请按实际情况进行修改

初始化并运行：

推荐使用 Goland 打开 server 目录，找到 Terminal(终端)，执行如下命令。
```
$ cd server
$ go generate
$ go build -o server main.go (windows编译命令为 go build -o server.exe main.go )

# 运行二进制
$ ./server (windows运行命令为 server.exe)
```


#### 安装二：Web后台管理前端（web）

推荐使用 Goland 或 WebStorm 打开 web 目录，找到 Terminal(终端)，执行如下命令。
```
$ cd web
$ npm install
$ npm run serve
```

成功启动后，即可通过浏览器访问：http://localhost:8080/#/login  

用户名: admin 密码: 123


#### 安装三：微信小程序用户前端（app）

需要使用微信开发者工具打开 app 目录， 找到 Terminal(终端)，执行如下命令。
```
$ cd app 
$ npm install
```

在编译运行微信小程序之前，你需要进行以下设置：

1、在微信开发者工具右上角->【详情】->【本地设置】-> 选择【使用npm模块】和【不校验合法域名，web-view（业务域名）、TLS版本...】。

2、在微信开发者工具的工具栏->【工具】->【构建npm】。

3、如果你没有自己的 appid 和 appSecret， 也就是说你没有修改 mall/server 目录下的配置文件里的 appid 和 appSecret，那么你是无法登录小程序商城的，你可以通过修改 mall/app/app.json 文件：将  "pages/login/login" 与 "pages/main/home/home" 交换位置，让他优先加载首页，从而跳过登录页。

4、如果你想要使用微信扫码的方式预览小程序商城，你需要修改 mall/app/utils/request.js 中的 'http://localhost:8000/app', 将 localhost 改成本地网络的ip地址，同时还要保证你扫码预览的手机所连接的网络必须是同一个网络（就是手机连接的网络和运行mall/app的电脑所连接的网络要一样）。

也为了能够正常访问图片，你需要修改 mall/demo/sql/product.sql 执行所生成的 product表 的image_url字段中的路径，同样也是把localhost 改成本地网络的ip地址。

#### 运行截图

web后台管理员前端：

![运行结果1](https://github.com/zchengo/mall/blob/main/demo/result/web_1.png?raw=true)
![运行结果2](https://github.com/zchengo/mall/blob/main/demo/result/web_2.png?raw=true)
![运行结果3](https://github.com/zchengo/mall/blob/main/demo/result/web_3.png?raw=true)

微信小程序用户前端：

![运行结果1](https://github.com/zchengo/mall/blob/main/demo/result/app_1.png?raw=true)
![运行结果2](https://github.com/zchengo/mall/blob/main/demo/result/app_2.png?raw=true)
![运行结果3](https://github.com/zchengo/mall/blob/main/demo/result/app_3.png?raw=true)

### 使用说明

1、本商城系统使用MIT开源许可证，完全免费，请放心使用。

2、部分图标来源www.iconfont.cn，图标仅供学习使用。

### 问题反馈

本项目持续更新、优化，在使用过程中遇到问题，可以提交 Issues，也可以通过以下方式反馈：

作者的知乎：https://www.zhihu.com/people/ku-ku-de-yuan-4

作者的CSDN：https://blog.csdn.net/m0_47890251?spm=1000.2115.3001.5343



