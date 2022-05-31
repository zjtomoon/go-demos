package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var wc sync.WaitGroup

//SendData sends data to server.
func SendData(c *http.Client, url string, method string, filePath string) {
	defer wc.Done()

	if c == nil {
		log.Fatalln("client is nil")
	}
	if method == "POST" {
		boundary := "ASSDFWDFBFWEFWWDF" //可以自己设定，需要比较复杂的字符串作
		var data []byte
		if _, err := os.Lstat(filePath); err == nil {
			file, _ := os.Open(filePath)
			defer file.Close()

			data, _ = ioutil.ReadAll(file)
		} else {
			log.Fatal("file not exist")
		}

		picData := "--" + boundary + "\n"
		picData = picData + "Content-Disposition: form-data; name=\"userfile\"; filename=" + filePath + "\n"
		picData = picData + "Content-Type: application/octet-stream\n\n"
		picData = picData + string(data) + "\n"
		picData = picData + "--" + boundary + "\n"
		picData = picData + "Content-Disposition: form-data; name=\"text\";filename=\"1.txt\"\n\n"
		picData = picData + string("data=ali") + "\n"
		picData = picData + "--" + boundary + "--"

		req, err := http.NewRequest(method, url, strings.NewReader(picData))
		req.Header.Set("Content-Type", "multipart/form-data; boundary=" + boundary)
		if err == nil {
			if rep, err := c.Do(req); err == nil {
				content, _ := ioutil.ReadAll(rep.Body)
				log.Println("get response: " + string(content))
				rep.Body.Close()
			}
		}
	} else if method == "GET" {
		//TODO get data from server
	}
}

func main() {
	client := &http.Client{
		Timeout: time.Second * 3,
	}
	postImgPath := "1.png"
	method := "POST"
	url := "http://127.0.0.1:8000/postdata"
	wc.Add(1)

	go SendData(client, url, method, postImgPath)

	wc.Wait()
}


/**
在客户端代码中，有几点需要说明一下：

multipart/form-data：在请求头中的Content-Type字段应该设置成multipart/form-data，
利用表单的形式上传文件。同时还应该加上boundary的内容，boundary可以自己指定。

boundary：如何大家和本文一样自己构建请求体，需要利用boundary进行数据分隔，
分隔的格式有严格的要求，下面给出一个范例。其中name字段是表单中该元素的名词，
filename是内容的名称(可以认为是文件名)，二者是有区别的；另外具体的内容，
如1.png的内容与前一行必须有一个空行(回车)；boundary加上字符"--"作为分隔符，而结束的boundary还应该在末尾加上字符"--"，

todo: 参考 https://zhuanlan.zhihu.com/p/136774587

 */