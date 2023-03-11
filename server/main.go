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
	Email        *string
	Age          uint8
	School       string
	
}

//自动迁移
db.AutoMigrate(&List{})


	// 接口
	r:=gin.Default()


	// 端口
	PORT:="3000"
	r.Run(":"+PORT)





	
}

