package dao

import (
	"document_server/model"
	"log"
)

// todo: password直接保存为string是否安全？
func CheckAuth(name string, password string) bool {
	var user model.User
	err := db.Table("user_table").Select("password").Where("user_name = ?", name).Find(&user).Error
	if err != nil {
		log.Println("Failed to find,err:", err)
	}
	if user.Password == password {
		return true
	}
	return false
}
