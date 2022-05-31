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
