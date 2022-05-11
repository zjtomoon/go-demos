package controller

import "document_server/utils"

// todo: 处理请求体
// 约定传入的参数为`操作:username+password`,例如：login:szf+123
func HandleReq(req string) (reqtype string, username string, password string) {
	indexofConn := utils.FindindexofSymbol(req, ':')
	reqtype = utils.SubString(req, 0, indexofConn)
	indexofPlus := utils.FindindexofSymbol(req, '+')
	username = utils.SubString(req, indexofConn+1, indexofPlus) // 切片左闭右开，所以无须减1
	password = utils.SubString(req, indexofPlus+1, len(req))
	return reqtype, username, password
}
