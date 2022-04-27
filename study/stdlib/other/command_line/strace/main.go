package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func main() {
	//首先生成cmd结构体
	cmd := exec.Command("strace", "-p", "15284")
	//调用stderrPipe生成一个管道，该管道连接到命令行进程的标准错误， 该方法返回一个
	//ReadCloser， 我们可以通过读取返回的ReadCloser来实时读取命令行进程的标准错误
	piper, _ := cmd.StderrPipe()
	//开始执行命令
	cmd.Start()
	//使用bufio包封装的方法实现对reader的读取
	reader := bufio.NewReader(piper)
	for {
		//换行分隔
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		//打印内容
		fmt.Println(line)
	}
	//等待命令结束并回收子进程资源等
	cmd.Wait()
}
