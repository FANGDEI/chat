# chat
使用Go基于WebSocket的IM系统
## 技术框架
* web框架iris
* ORM框架gorm
* RPC框架gRPC
* 数据库MySQL
* 消息队列Redis
* 对象存储Minio
* 长连接websocket
* 日志框架Uber的zap
* 服务注册与配置中心consul
* 部署使用Makefile以及Dockerfile
## 目录结构
```FILE
.
├── api ------------> 模块PROTO
├── cmd ------------> 模块及网关main文件
├── config ---------> 配置文件
├── database -------> 数据库文件存放目录
├── internal -------> 业务逻辑存放目录
│   ├── app --------> GRPC模块业务逻辑
│   └── pkg --------> 底层业务逻辑代码包
│   └── web --------> 网关业务逻辑存放目录
└── package --------> 模块及网关打包目录
```