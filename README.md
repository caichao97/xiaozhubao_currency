# xiaozhubao_currency
框架介绍：

beego 是一个快速开发 Go 应用的 HTTP 框架，也是一个 RESTful模式的框架。
REST：指的就是一组约束条件和原则。满足这些约束条件和原则的应用程序或设计就是RESTful。Web应用程序最重要的原则是客户端和服务器之间的交互，从客户端到服务器的每个请求都必须包含理解请求所必须的信息。如果服务器在请求之间的任何时候重启或宕机，客户端不会得到通知。客户端可以通过缓存数据的形式来提高性能。
架构目录：

├── conf           //配置文件
│   └── app.conf
├── controllers    //控制器
│   ├── admin
│   └── default.go
├── main.go        //项目入口
├── models         //模型   
│   └── models.go
├── routers       //路由
│   └──router.go
├── static         //静态文件
│   ├── css
│   ├── ico
│   ├── img
│   └── js
└── views          //界面模板
    ├── admin
    └── index.tpl


Linux 常见报错及解决方案： windows 待补充

1、beego 安装无反应：

1.1 执行 go get github.com/astaxie/beego后无反应，将以下添加到hosts文件后再次安装
192.30.253.112 github.com
151.101.185.194 github.global.ssl.fastly.net

1.2 执行 go get github.com/beego/bee后无反应，执行以下语句
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io,direct
2、后台运行项目：

nohup *** & >> PATH/***.log 
nohup                    如果你正在运行一个进程，而且你觉得在退出帐户时该进程还不会结束，那么可以使用nohup命令。该命令可以在你退出帐户/关闭终端之后继续运行相应的进程。
***                        你的项目中生成的可执行二进制文件。注意：需有可执行权限
&                          后台执行
>>                         输出到
PATH/***.log         定义的日志文件地址。注意：PATH为你定义的实际绝对路径
例如：nohup xiaozhubao_currency & >> /tmp/xiaozhubao_currency

中文文档：http://www.topgoer.com/beego框架/
