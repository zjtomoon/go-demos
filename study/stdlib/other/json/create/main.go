package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// 结构体内元素注意大写，否则写入文件为{}
type ConfigFile struct {
	Path string
	Name string
}

func NewConfigFile(path string, name string) *ConfigFile {
	return &ConfigFile{
		Path: path,
		Name: name,
	}
}

func main() {
	CreateJson("path", "filename")
}

func CreateJson(path string, name string) {
	file := NewConfigFile(path, name)
	//file, err := os.Create("FileConfig.json")
	//if err != nil {
	//	fmt.Println("Create file failed,err:", err)
	//	return
	//}
	//
	//defer file.Close()
	//data, err := json.MarshalIndent(configfile, "", " ")
	//if err != nil {
	//	fmt.Println("Encoder failed,err:", err)
	//	return
	//} else {
	//	fmt.Println("Encoder success")
	//}
	//file.Write(data)

	//configfile := ConfigFile{
	//	path:path,
	//	name:name,
	//}

	b, err := json.Marshal(file)
	if err != nil {
		fmt.Println("error:", err)
	}
	//生成json文件
	err = ioutil.WriteFile("test.json", b, os.ModeAppend)
	if err != nil {
		return
	}

	var data interface{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("data: ", data)

}
