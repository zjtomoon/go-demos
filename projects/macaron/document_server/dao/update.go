package dao

import (
	"log"
	"strconv"
)

// todo: 测试出错
// 更新数据库密码
func UpdatePassword(name string, newpasswd string) {
	err := db.Table("user_table").Where("user_name = ?", name).Update("password", newpasswd).Error
	if err != nil {
		log.Println("update failed,err:", err)
	}
}

// 更新user_table的文件表表名
// user_table中file_table初始化为空，通过该方法更新,
// 推荐随后就调用CreateFileTableX方法，创建该用户对应的文件表
func UpdateFileTableName(username string) {
	var newname string
	id := FindIDByName(username)
	newname = strconv.FormatInt(id, 10) + "_" + username
	err := db.Table("user_table").Where("user_name = ?", username).Update("file_table", newname).Error
	if err != nil {
		log.Println("Failed to update file_table name  in user")
		return
	}
}
