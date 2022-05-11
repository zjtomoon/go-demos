package dao

import (
	"document_server/model"
	"log"
)

func DropUserbyID(id int) {
	user := model.User{}
	err := db.Table("user_table").Where("Id = ?", id).Delete(&user).Error
	if err != nil {
		log.Println("Failed to drop data,err:", err)
	}
}

func DropUserByName(name string) {
	user := model.User{}
	err := db.Table("user_table").Where("UserName = ?", name).Delete(&user).Error
	if err != nil {
		log.Println("Failed to drop data,err:", err)
	}
}
