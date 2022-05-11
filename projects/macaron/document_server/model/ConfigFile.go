package model

// todo: 添加文件序列number、是否更新update、文件内容content
// todo: 结构体嵌套结构体
// todo: 使用结构体数组，显示多行
type ConfigFile struct {
	No       int    `json:"no"`
	FilePath string `json:"file_path"`
	FileName string `json:"file_name"`
	Update   int    `json:"update"`
	Content  string `json:"content"`
}

type JsonFile struct {
	Configfile []ConfigFile `json:"configfile"`
}

func NewConfigFile(no int, path string, filename string, update int, content string) *ConfigFile {
	return &ConfigFile{
		No:       no,
		FilePath: path,
		FileName: filename,
		Update:   update,
		Content:  content,
	}
}

//func NewJsonFile(configFile *ConfigFile) *JsonFile {
//	return &JsonFile{
//		Configfile: *configFile,
//	}
//}
