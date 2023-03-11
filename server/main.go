package main



import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/schema"
  )


func main() {
	// 小试牛刀
	fmt.Println("hello，全栈工程师～")

	// 连接数据库
	dsn := "root:12345678@tcp(127.0.0.1:3306)/crud-list?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{ 
		//解决表名复数问题
		NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
		},

	})
	fmt.Println("db:",db)
	fmt.Println("err:",err)


	// 连接池
// GORM 使用database/sql维护连接池
	sqlDB, err := db.DB()
// SetMaxIdleConns 设置空闲连接池中连接的最大数量
sqlDB.SetMaxIdleConns(10)

// SetMaxOpenConns 设置打开数据库连接的最大数量。
sqlDB.SetMaxOpenConns(100)

// SetConnMaxLifetime 设置了连接可复用的最大时间。
sqlDB.SetConnMaxLifetime(10 * 1000)  //10秒钟


// 结构体
type List struct {
	gorm.Model    //解决主键缺失
	Name         string
	Age          uint8
	School       string
	Phone        string
	Address       string

	
}

//自动迁移
db.AutoMigrate(&List{})


	// 接口
	r:=gin.Default()

	//错误优先原则、业务码约定：
	// 成功：200
	// 失败：400


	// 增加：
	r.POST("/list/add",func(c*gin.Context){
		var data  List
		err :=c.ShouldBindJSON(&data)	
		if err != nil {
			c.JSON(200,gin.H{
				"msg":"添加失败",
				"data":gin.H{},
				"code":400,
			})
		}else{
			// 操作数据库
			db.Create(&data)  

			c.JSON(200,gin.H{
				"msg":"添加成功",
				"data":gin.H{},
				"code":400,

			})
		}


	})

	// 删除
	r.DELETE("/list/delete/:id",func(c*gin.Context){
		var data  []List
		// 接收前端传过来的ID
		id :=c.Param("id")
		// 判断ID是否存在
		db.Where("id =?",id).Find(&data)

		// ID存在则进行删除、不存在则进行报错
		if len(data)==0{
			c.JSON(200,gin.H{
				"msg":"id没有找到,删除失败",
				"code":400,

			})
		}else{
			// 操作数据库
			db.Where("id=?",id).Delete(&data)
			c.JSON(200,gin.H{
				"msg":"删除成功",
				"code":200,
			})
		}

	})


	// 修改
	r.PUT("/list/update/:id",func(c*gin.Context){
		var data  List
		// 接收前端传过来的ID
		id :=c.Param("id")
		// 判断ID是否存在
		db.Select("id").Where("id =?",id).Find(&data)

		// ID存在则进行修改、不存在则进行报错
		if data.ID==0{
			c.JSON(200,gin.H{
				"msg":"id没有找到,修改失败",
				"code":400,

			})
		}else{
			err :=c.ShouldBindJSON(&data)	
			if err !=nil{
				c.JSON(200,gin.H{
					"msg":"修改失败",
					"code":400,
	
				})

			}else{
				// 操作数据库
				db.Where("id=?",id).Updates(&data)
				c.JSON(200,gin.H{
					"msg":"修改成功",
					"code":200,
	
				})
			}
		}

	})




	// 端口
	PORT:="3000"
	r.Run(":"+PORT)





	
}

