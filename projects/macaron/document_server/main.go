package main

import (
	"document_server/controller"
	"document_server/dao"
	"document_server/logs"
	"document_server/utils"
	"gopkg.in/macaron.v1"
	"log"
	"net/http"
)

var m = macaron.Classic()

func main() {

	PrintReq()
	SendJsonData()

	// 开启临时https服务
	log.Fatalln(http.ListenAndServeTLS(":3000", "tls/server.crt", "tls/server.key", m))

	// todo: 推荐使用http + nginx方式部署https服务，性能更高
	//log.Fatalln(http.ListenAndServe(":3000",nil))
	//m.Run(3000)
}

func PrintReq() {
	//http.HandleFunc("/", controller.PrintRequestInfo)
	m.Post("/", controller.PrintRequestInfo)
}

func SendJsonData() {

	jsondir, jsonpath := dao.SetJsonPath("./config", "user1")

	utils.CreateDirs(jsondir)
	//utils.WriteJson(jsonpath, 1, "~", "user2", 0, "123")
	arr1 := utils.CreateFileElements(1, "~", "user1.txt", 0, "123")
	arr2 := utils.CreateFileElements(2, "~", "user2.txt", 0, "12345")
	arr3 := utils.CreateFileElements(3, "~", "user3.txt", 0, "123456")

	file := utils.InsertConfigFile2Arr(*arr1, *arr2, *arr3)
	utils.WriteJsons(jsonpath, *file)

	controller.Des = jsonpath
	//http.HandleFunc("/download", controller.SendJson)
	m.Get("/download", controller.SendJson)
}

func SetLogConfig() {
	//自定义日志收集方式
	logs.SetLogger("config/log.json")
}
