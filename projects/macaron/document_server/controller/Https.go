package controller

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var Des string

type UrlStruct struct {
	UrlName string `json:"url_name"`
}

//type HttpServer struct {}

// 参考：https://studygolang.com/articles/5665
// todo: 转httpHandleFunc为httpHandle

func GetResponse(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln("failed to  get response ,err ", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if resp.StatusCode == 200 {
		fmt.Println("ok")
	}
}

// todo:golang获取http body内容
// 参考：https://www.jianshu.com/p/1b0f647e8519
func PrintRequestInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("requst url : %v\n", r.URL)
	fmt.Printf("request Header: %v\n", r.Header)
	fmt.Printf("request body : %v\n", r.Body)
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Printf("request body : %v\n", string(body))

	if string(body) == "" {
		w.Write([]byte("failed to read body"))
	} else {
		w.Write([]byte("succeed to read body"))
	}
}

// todo:下载文件服务
// 参考：https://blog.csdn.net/xmcy001122/article/details/104654096
// 参考: https://blog.csdn.net/lightjia/article/details/82700419

func SendJson(w http.ResponseWriter, req *http.Request) {

	//w.Write([]byte("success download\n"))
	fileData, err := ioutil.ReadFile(Des)
	if err != nil {
		log.Println("Read File err:", err.Error())
	} else {
		log.Println("Send File:", Des)
		w.Write(fileData)
	}
}

//var bodystr string
//// 接收请求体信息
//func ReceiveRequest(w http.ResponseWriter, r *http.Request)  {
//	body, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		logger.Error("Failed to read request body")
//	}
//	defer r.Body.Close()
//	bodystr = string(body)
//}

// 登录响应
// 登录成功，返回 succeed to login
// 登录失败，返回 fail to login
//func LoginResponse(w http.ResponseWriter, r *http.Request) {
//
//}
//
//func RegisterResponse(w http.ResponseWriter, r *http.Request) {
//
//}
