# 搭建大型Web项目通用脚手架
## 1. 大型Web项目CLD分层理念
## 1.1 MVC模式

MVC 模式代表 Model-View-Controller（模型-视图-控制器） 模式。这种模式用于应用程序的分层开发。

- **Model（模型）** - 模型代表一个存取数据的对象或 JAVA POJO。它也可以带有逻辑，在数据变化时更新控制器。
- **View（视图）** - 视图代表模型包含的数据的可视化。
- **Controller（控制器）** - 控制器作用于模型和视图上。它控制数据流向模型对象，并在数据变化时更新视图。它使视图与模型分离开。

<img src="https://mary-aliyun-img.oss-cn-beijing.aliyuncs.com/typora/202311061706918.png" alt="img" style="zoom:40%;" />


### 1.2 CLD模式

<img src="https://mary-aliyun-img.oss-cn-beijing.aliyuncs.com/typora/202311061931687.png" alt="image-20231106193126611" style="zoom:50%;" />

- 协议处理层：支持各种协议
- Controller：服务的入口，负责处理路由、参数校验、请求转发
- Logic/Service：逻辑（服务）层，负责处理业务逻辑
- DAO/Repository：负责数据与存储相关功能

### 1.3 CLD模式脚手架的搭建
- 项目结构
```text
├── README.md
├── app
│   ├── logger
│   ├── controller
│   ├── dao
│   │   └── mysql
│   │   └── redis
│   ├── logic
│   └── router
│   │   └── router.go
│   ├── model
│   │   └── model.go
│   ├── settings
│   │   └── settings.go
│   ├── pkg
├── main.go
├── config.yaml
```
- 项目介绍
```text
web项目整体基于gin框架，加入了优雅的关闭模式
MySQL数据库使用sqlx
redis使用go-redis
配置文件使用viper
日志使用zap
```

- 项目依赖
```shell
go get github.com/gin-gonic/gin
go get github.com/go-sql-driver/mysql
go get github.com/go-redis/redis
go get github.com/spf13/viper
go get go.uber.org/zap
go get github.com/jmoiron/sqlx
```


