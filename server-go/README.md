# go-vue-react
# author：fangdongdong
# date：2023年3月11日

# 项目技术栈：
该项目是一个全栈项目，前后端分离开发，其中包括一个后端，两个前端。
大家根据自己的需求，自行选择学习其中的一个或者多个项目。

### 后端：
go+gin+gorm+mysql

### 前端vue：
Vue3.2 + Vite + ElementPlus+ts

###  前端react:


# 适用人群：
### 1、学习前端最新技术
### 2、零基础入门go语言
### 3、进阶成为全栈工程师


# 实现功能：
1. 实现最简单的crud


# 使用方法
请查看各项目下的readme文件
go-vue-react/README.md :项目总体概述
go-vue-react/server-go/README.md:后端go项目概述
go-vue-react/web-vue/README.md：前端vue项目概述
go-vue-react/web-react/README.md：前端react项目概述


# 以下是当前后端go开发中使用的一些最新技术栈及其相关资源：

Go 1.16：Go是一个快速、高效且易于编写的编程语言，用于构建高性能的网络服务器和分布式系统。Go 1.16是最新版本，提供了更好的性能和可选功能。官方网站：https://golang.org/

Gin：Gin是一个快速、高效、轻量级的HTTP Web框架，它采用Go的优秀特性，如多路复用和协程，使得处理高并发请求变得轻松愉悦。官方网站：https://gin-gonic.com/

GORM：GORM是一个适用于Go的简单易用的ORM库，它支持MySQL、PostgreSQL和SQLite等多种数据库，并提供了大量的功能，如关联查询、事务处理等。官方网站：http://gorm.io/

Redis：Redis是一个开源的高性能键值存储数据库，它支持各种数据类型，如字符串、哈希表、列表等，并提供了丰富的命令集合，使得缓存和队列变得容易。官方网站：https://redis.io/

Kafka：Kafka是一个高性能、可伸缩的消息中间件平台，它可以处理大规模的实时消息流，并提供了持久化、副本、批处理等功能。官方网站：https://kafka.apache.org/

Prometheus：Prometheus是一个开源的监控系统和时间序列数据库，它专门针对分布式应用程序设计，能够收集和记录度量数据，并提供强大的查询语言和可视化工具。官方网站：https://prometheus.io/

Docker：Docker是一个开源的容器化平台，它使得应用程序的部署、管理和扩展变得更加容易和可靠。官方网站：https://www.docker.com/

Kubernetes：Kubernetes是一个开源的容器编排平台，它可以自动化容器的部署、扩展、管理和运行，从而简化了应用程序的部署和管理。官方网站：https://kubernetes.io/


===================================================== 

# 后端项目  server

# 一、安装项目相关环境和依赖
### 1、安装go环境和安装mysql数据库
新建数据库crud-list
账户名：root
密码：12345678
端口：3306

### 2、新建项目名、初始化项目依赖
1.新建server项目名，在server下新建main.go入口文件
2.在server下执行go mod init server  初始化项目依赖，生成go.mod文件

### 3、引入gin到项目中
1.下载并安装 gin：
go get -u github.com/gin-gonic/gin
2.将 gin 引入到代码中：
import "github.com/gin-gonic/gin"

### 4、引入gorm到项目中
1.下载安装
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
2.引入gorm
import (
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
)
3、引入mysql到项目中
go get gorm.io/driver/mysql
import (
 "gorm.io/driver/mysql"
)

# 二、创建数据库、连接数据库
### 1.navacat可视化数据库
### 2.连接数据库
### 3.配置连接池、定义结构体、自动迁移
### 4.解决主键缺失和表名重复问题
### 5.优化定义结构体

# 三、实现最简单的crud
### 1.post新增接口
### 2.put修改接口
### 3.delete删除接口
### 4.get条件查询接口
### 5.get请求全部数据和分页数据




