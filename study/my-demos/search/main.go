package main

import (
	"fmt"
	"search/cmd"
	"search/config"
	"time"
)

func main() {

	// todo 读取配置文件 选择搜索特定的目录和文件
	jsonParse := config.NewJsonStruct()
	v := config.SearchReq{}
	jsonParse.Load("./config.json", &v)
	cmd.Query = v.Filename
	// todo 守护进程
	for {
		fmt.Println("current time :", time.Now())
		cmd.SearchOperator(v.Path)
		time.Sleep(10 * time.Second)
	}
}
