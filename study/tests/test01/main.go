package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id   int
	Age  int
	Name string `gorm:"size:255"`
	//Admin interface{}

	//TabName string `gorm:"_"` //忽略这个字段，给表设置名字的
}

//func NewUser(tabName string) *User {
//	return &User{
//		TabName:tabName,
//	}
//}

//func (User) TableName() string {
//	return "user_table"
//}

// TableName 不支持动态变化，它会被缓存下来以便后续使用。想要使用动态表名，你可以使用下面的代码：

//func UserTable(user User) func (db *gorm.DB) *gorm.DB {
//	return func (db *gorm.DB) *gorm.DB {
//		if user.Admin {
//			return db.Table("admin_users")
//		}
//		return db.Table("users")
//	}
//}
//DB.Scopes(UserTable(user)).Create(&user)

//func (u User) TableName() string {
//	return u.TabName
//}

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/document_server?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("连接数据库成功!")
	db.Table("user2").AutoMigrate(&User{})
	db.Table("user3").AutoMigrate(&User{})
	user1 := User{
		Id:   1,
		Age:  19,
		Name: "user1",
	}

	db.Table("user1").Create(&user1)
	//db.AutoMigrate(&User{})
}
