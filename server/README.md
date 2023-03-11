# fdd-crud-fullStack
# author：fangdongdong
# date：2023年3月11日

# 项目技术栈：

### 前端：
Vue3.2 + Vite + ElementPlus

### 后端：
go+gin+gorm+mysql

# 实现功能：实现最简单的crud

# 适用人群：
### 1、学习前端最新技术
### 2、零基础入门go语言
### 3、进阶成为全栈工程师


===================================================== 

# 后端项目  server

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
3、引入mysql到项目中，下载sqlite就有mysql,可以直接引入
import (
  "gorm.io/driver/sql"
)

# 二、创建数据库、连接数据库
### 1.navacat可视化数据库
新建数据库、设置密码账号

### 2.定义结构体，写入数据库

### 3.连接数据库
### 4.解决主键缺失和表名重复问题
### 5.优化定义结构体

# 三、实现最简单的crud
### 1.post新增接口
### 2.put修改接口
### 3.delete删除接口
### 4.get条件查询接口
### 5.get请求全部数据和分页数据




