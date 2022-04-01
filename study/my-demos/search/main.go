package main

import (
	"fmt"
	"search/cmd"
	"time"
)

func main() {

	// 参考: https://blog.csdn.net/benben_2015/article/details/79134734
	for {
		fmt.Println("current time :", time.Now())
		cmd.SearchOperator()
		time.Sleep(10 * time.Second)
	}
	// todo 守护进程
	// todo 配置文件法/args 选择搜索特定的目录和文件
	// todo 搜索到该文件 打印出该文件的目录地址
	// todo 搜索到文件，返回字符串
}
