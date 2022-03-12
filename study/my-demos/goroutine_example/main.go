package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(id string) {
	time.Sleep(time.Second)
	fmt.Println("I am done ! id : " + id)
	wg.Done()
}

func main() {
	wg.Add(2) // 总共有两个任务

	go func(id string) {
		fmt.Println(id)
		wg.Done()
	}("hello")

	go say("world")

	wg.Wait() // 等待所有任务完成 卡住 如果 wg不是0
	fmt.Print("exit")
}
