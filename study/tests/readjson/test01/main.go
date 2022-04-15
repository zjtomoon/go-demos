package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//type SearchReq struct {
//	Path     string `json:"path"`
//	Filename string `json:"filename"`
//}

type UrlStruct struct {
	UrlName string `json:"url_name"`
}

type JsonStruct struct {
}


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

func main() {
	jsonParse := NewJsonStruct()
	v := UrlStruct{}
	jsonParse.Load("./config.json",&v)
	fmt.Println(v.UrlName)
}