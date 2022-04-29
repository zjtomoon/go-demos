package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

func main() {
	server := http.Server{Addr: "localhost:8899"}

	http.HandleFunc("/showdownload", showDownloadPage)
	http.HandleFunc("/download", download)
	server.ListenAndServe()
}

func showDownloadPage(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(rw, nil)
}

func download(rw http.ResponseWriter, r *http.Request) {
	// 获取请求参数
	fn := r.FormValue("filename")
	// 设置响应头
	header := rw.Header()
	header.Add("Content-Type", "application/octet-stream")
	header.Add("Content-Disposition", "attachment;filename="+fn)
	// 使用ioutil包读取文件
	b, _ := ioutil.ReadFile("D:/" + fn)
	// 写入到响应流中
	rw.Write(b)
}
