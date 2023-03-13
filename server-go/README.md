# go-gin-gorm+mysql


# 使用说明：
1. 新建数据库
### 数据库名：crud-list
### 账户名：root
### 密码：12345678
### 端口：3306

2. 切换到项目目录
### cd server-go

3. 打包：
### go build  
go build ：如果当前项目使用 Go Modules 管理依赖，则编译或运行时会自动下载所需的依赖。

4. 运行打包文件：
### ./server 


# 开发笔记

# 一、安装项目相关环境和依赖
### 1、安装go环境和安装mysql数据库


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
数据库名：crud-list
账户名：root
密码：12345678
端口：3306
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








