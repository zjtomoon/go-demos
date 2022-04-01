package main

import (
	"fmt"
	"search/cmd"
	"search/config"
	"time"
)

func main() {

	// 参考: https://blog.csdn.net/benben_2015/article/details/79134734
	jsonParse := config.NewJsonStruct()
	v := config.SearchReq{}
	jsonParse.Load("./config.json", &v)
	fmt.Println(v.Filename)
	fmt.Println(v.Path)

	// todo 守护进程
	for {
		fmt.Println("current time :", time.Now())
		cmd.SearchOperator()
		time.Sleep(10 * time.Second)
	}
}
