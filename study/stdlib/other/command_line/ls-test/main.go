package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	//首先生成cmd结构体,该结构体包含了很多信息，如执行命令的参数，命令的标准输入输出等
	command := exec.Command("ls", "-l")
	//给标准输入以及标准错误初始化一个buffer，每条命令的输出位置可能是不一样的，
	//比如有的命令会将输出放到stdout，有的放到stderr
	command.Stdout = &bytes.Buffer{}
	command.Stderr = &bytes.Buffer{}
	//执行命令，直到命令结束
	err := command.Run()

	if err != nil {
		fmt.Println(err)
		fmt.Println(command.Stderr.(*bytes.Buffer).String())
		return
	}

	// 打印命令行的标准输出
	fmt.Println(command.Stdout.(*bytes.Buffer).String())

}
