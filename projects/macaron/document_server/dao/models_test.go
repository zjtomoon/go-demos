package dao

import "testing"

func TestCreateUserTable(t *testing.T) {
	CreateUserTable()
}

func TestCreateUserInfo(t *testing.T) {
	CreateUserInfo("user1", "123456")
	CreateUserInfo("user2", "123456")
}

// todo: 测试程序运行失败，但是能在主程序执行成功
func TestUpdatePassword(t *testing.T) {
	UpdatePassword("user1", "123")
}

// todo: 测试程序运行失败，但是能在主程序执行成功
func TestUpdateFileTableName(t *testing.T) {
	UpdateFileTableName("user1")
}

func TestCreateFileTableX(t *testing.T) {
	CreateFileTableX("user2")
}
