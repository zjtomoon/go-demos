package dao

import (
	"document_server/model"
	"document_server/utils"
	"log"
	"strconv"
)

// 使用gorm创建user_table，仅执行一次
//todo: create table if not exists
func CreateUserTable() {
	err := db.Table("user_table").AutoMigrate(&model.User{})
	if err != nil {
		log.Println("failed to create user table")
		return
	}
}

//todo: 测试 根据id查找username
func FindUserName(id int64) (name string) {
	var user model.User
	err := db.Table("user_table").Select("user_name").Where(model.User{Id: id}).Find(&user).Error
	if err != nil {
		log.Println("Failed to find username,err:", err)
	}
	return user.UserName
}

// 根据id_username的格式创建表名，按需求执行多次
// 结束后，返回表名
func CreateFileTableX(username string) string {
	var newName string
	id := FindIDByName(username)
	newName = strconv.FormatInt(id, 10) + "_" + username
	err := db.Table(newName).AutoMigrate(&model.FileTable{})
	if err != nil {
		log.Println("failed to create file table")
	}
	return newName
}

func CreateUserInfo(username string, password string) {
	id := utils.GenerateIDS()
	user := model.NewUser(id, username, password)
	err := db.Table("user_table").Create(user).Error
	if err != nil {
		log.Println("Failed to Insert userinfo", err)
	}
}
