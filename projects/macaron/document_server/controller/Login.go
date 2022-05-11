package controller

import (
	"document_server/dao"
	"github.com/wonderivan/logger"
)

// todo: 接收Login信息

func Login(username string, password string) bool {
	auth := dao.CheckAuth(username, password)
	if auth {
		return true
	}
	return false
}

// 接收到登录请求后执行登录
func DoLogin(loginRequest string, username string, password string) {
	if loginRequest == "Login" || loginRequest == "login" {
		isLogin := Login(username, password)
		if isLogin {
			logger.Info(username + "succeed to login")
		} else {
			logger.Info(username + "failed to login")
		}
	}
}
