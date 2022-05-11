package utils

import (
	"document_server/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// todo: 生成json配置文件
// todo:添加jsonpath/filename
// 参考: https://www.cnblogs.com/baylorqu/p/9663016.html
// 将单个配置文件写入json
func WriteJson(jsonpath string, no int, path string, name string, update int, content string) {
	//configfile := model.NewConfigFile(path, name)
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

	configfile := model.NewConfigFile(no, path, name, update, content)
	jsonfile := model.JsonFile{[]model.ConfigFile{*configfile}}
	b, err := json.MarshalIndent(jsonfile, "", "")
	if err != nil {
		fmt.Println("error:", err)
	}
	// 生成json文件
	err = ioutil.WriteFile(jsonpath, b, os.ModeAppend)
	if err != nil {
		return
	}

	var data interface{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("data:", data)
}

func CreateFileElements(no int, path string, name string, update int, content string) *model.ConfigFile {
	configfile := model.NewConfigFile(no, path, name, update, content)
	return configfile
}

// 写入多行json数据
func WriteJsons(jsonpath string, configArr []model.ConfigFile) {

	jsonfile := model.JsonFile{configArr}
	b, err := json.MarshalIndent(jsonfile, "", "")
	if err != nil {
		fmt.Println("error:", err)
	}
	// 生成json文件
	err = ioutil.WriteFile(jsonpath, b, os.ModeAppend)
	if err != nil {
		return
	}

	var data interface{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("data:", data)
}

// todo: 在指定目录(例如./config目录)下创建形如1_user1/的文件夹
func CreateDirs(jsonPath string) {
	if _, err := os.Stat(jsonPath); os.IsNotExist(err) {
		os.Mkdir(jsonPath, 0777)
		os.Chmod(jsonPath, 0777)
	}
}

// todo:生成不定长度的结构体数组
func InsertConfigFile2Arr(params ...model.ConfigFile) *[]model.ConfigFile {
	var file []model.ConfigFile = make([]model.ConfigFile, 0)
	file = append(file, params...)
	return &file
}
