package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

type User struct {
	Id   int
	Age  int
	Name string `gorm:"size:255"`
	//Admin interface{}

	//TabName string `gorm:"_"` //忽略这个字段，给表设置名字的
}

func NewUser(id int, age int, name string) *User {
	return &User{
		Id:   id,
		Age:  age,
		Name: name,
	}
}

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

func init() {
	// gorm正确的初始化方式
	var err error
	db, err = gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/document_server?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("连接数据库成功!")
}

func main() {
	Search()
}

func Create() {
	// create
	//
	db.Table("user2").AutoMigrate(&User{})
	db.Table("user3").AutoMigrate(&User{})
	//user1 := User{
	//	Id:   1,
	//	Age:  19,
	//	Name: "user1",
	//}

	user1 := NewUser(1, 19, "user1")
	user2 := NewUser(2, 26, "user2")
	user3 := NewUser(3, 23, "user3")
	//
	err := db.Table("user2").Create(&user1).Error
	if err != nil {
		log.Println("err:", err)
	}
	err = db.Table("user2").Create(&user2).Error
	if err != nil {
		log.Println("err:", err)
	}
	err = db.Table("user2").Create(&user3).Error
	if err != nil {
		log.Println("err:", err)
	}

	//db.AutoMigrate(&User{})
}

func Search() {
	// search
	user := User{}
	err := db.Table("user2").Select("name").Where("id = ?", 2).Find(&user).Error
	if err != nil {
		log.Println("err:", err)
	}
	fmt.Println(user.Id, user.Name, user.Age)
}

func UpdateTable() {

	// update
	err := db.Table("user2").Where("name = ?", "user1").Update("age", 20).Error
	if err != nil {
		log.Println("err:", err)
	}
	user := User{}
	err = db.Table("user2").Select("name", "age").Where("id = ?", 1).Find(&user).Error
	if err != nil {
		log.Println("err:", err)
	}
	fmt.Println("user name :", user.Name)
	fmt.Println("user age :", user.Age)
}

func DeleteData() {
	//delete
	user := User{}
	db.Table("user2").Where("id = ?", 1).Delete(&user)
	fmt.Println(user.Name, user.Age, user.Id)
}
