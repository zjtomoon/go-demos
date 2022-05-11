package dao

import (
	"document_server/model"
	"log"
)

// 判断用户是否重复
func CheckRepeated(name string) bool {
	var user model.User
	err := db.Table("user_table").Select("UserName").Where("UserName = ?", name).Find(&user).Error
	if err != nil {
		log.Println("Failed to find ,err: ", err)
	}
	if user.UserName == name {
		return true
	}
	return false
}
