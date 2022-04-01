package config

import (
	"encoding/json"
	"io/ioutil"
)

type SearchReq struct {
	Path     string `json:"path"`
	Filename string `json:"filename"`
}

type SearchRes struct {
	FilePath string
	FileSize int64
}

// todo 配置文件法/args 选择搜索特定的目录和文件
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
