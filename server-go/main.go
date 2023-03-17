package main



import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/schema"
	"strconv"


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


// 结构体  数据库大写，json小写
type List struct {
	gorm.Model    //解决主键缺失
	ID           uint8   `json:"id"`
	Name         string   `json:"name"`
	Age          uint8     `json:"age"`
	School       string     `json:"school"`
	Phone        string      `json:"phone"`
	Address       string       `json:"address"`

	
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
				"code":200,

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

	// 查询
	r.GET("/list/query",func(c*gin.Context){
		// 获取路径参数
		var  dataList []List

		    //获取分页参数
    pageSize, _ := strconv.Atoi(c.Query("pageSize"))
    page, _ := strconv.Atoi(c.Query("page"))

     // 计算偏移量和限制数量，支持更灵活的分页参数设置
	 offset := (page - 1) * pageSize
	 limit := pageSize

    // 返回一个总数
    var total int64
    name := c.Query("name")

	fmt.Printf("查询条件：%s\n", name)
	fmt.Printf("offset: %d\n", offset)
    fmt.Printf("limit: %d\n", limit)

    db.Model(&List{}).Where("name LIKE ?", fmt.Sprintf("%%%s%%", name)).Count(&total)

    // 条件查询并进行分页
    db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name)).Offset(offset).Limit(limit).Find(&dataList)
	   


		//判断是否查询到数据
		if len(dataList)==0{
			c.JSON(200,gin.H{
				"msg":"没有查询到数据",
				"code":200,
				"data":gin.H{},

			})

		}else{
			c.JSON(200,gin.H{
				"msg":"查询成功",
				"code":200,
				"data":dataList,
				"total":total,

			})
		}
	})




	// 端口
	PORT:="3000"
	r.Run(":"+PORT)





	
}

