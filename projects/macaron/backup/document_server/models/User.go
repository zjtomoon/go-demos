package models

import "gitee.com/chunanyong/zorm"

const TableName = "User"

type User struct {
	zorm.EntityStruct
	Id       string `column:"id"`
	UserName string `column:"username"`
	Password string `column:"password"`
}

func NewUser(username string, password string) *User {
	return &User{
		Id:       zorm.FuncGenerateStringID(),
		UserName: username,
		Password: password,
	}
}

// 获得表名
func (entity User) GetTableName() string {
	return TableName
}

// 获得主键
func (entity User) GetPKColumnName() string {
	return "id"
}
