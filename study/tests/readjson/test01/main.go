package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var url string

type UrlStruct struct {
	
	urlname string `json:"urlname"`
}

type JsonStruct struct{}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func (jst *JsonStruct) Load(filename string, v interface{}) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}

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

func main() {
	jsonParse := NewJsonStruct()
	v := UrlStruct{}
	jsonParse.Load("./config.json",&v)
	fmt.Println(v.urlname)
}