package dao

import (
	"document_server/model"
	"log"
	"strconv"
)

// todo: 生成每个用户自定义的配置文件，自定义文件目录
func SetJsonPath(basedir string, username string) (string, string) {
	id := FindIDByName(username)
	jsondir := basedir + "/" + strconv.FormatInt(id, 10) + "_" + username + "/"
	jsonpath := basedir + "/" + strconv.FormatInt(id, 10) + "_" + username + "/" + username + ".json"
	return jsondir, jsonpath
}

func FindIDByName(username string) int64 {
	var user model.User
	err := db.Table("user_table").Select("id").Where("user_name = ?", username).Find(&user).Error
	if err != nil {
		log.Println("failed to find id of this username")
	}
	id := user.Id
	return id
}
