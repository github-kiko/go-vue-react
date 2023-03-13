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


# 使用方法
请查看各项目下的readme文件
go-vue-react/README.md :项目总体概述
go-vue-react/server-go/README.md:后端go项目概述
go-vue-react/web-vue/README.md：前端vue项目概述
go-vue-react/web-react/README.md：前端react项目概述



# 实现功能：
1. 实现最简单的crud


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




