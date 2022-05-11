package controller

import (
	"document_server/dao"
	"document_server/model"
	"github.com/wonderivan/logger"
)

var User *model.User

// todo: 接收注册信息

// 注册
// 返回true 表示注册成功
// 返回false 表示注册失败
func Register(username string, password string) bool {
	// todo: 校验是否在数据库中存在，不存在则注册
	// todo: 对password进行加密后存储到数据库
	isrepeated := dao.CheckRepeated(username)
	if !isrepeated {
		dao.CreateUserInfo(username, password)
		return true
	}
	return false
}

// 接收到注册请求后执行注册
func DoRegister(registerRequest string, username string, password string) {
	if registerRequest == "Register" || registerRequest == "register" {
		isRegister := Register(username, password)
		if isRegister {
			logger.Info(username + "succeed to register ")
		} else {
			logger.Info(username + "failed to register")
		}
	}
}
