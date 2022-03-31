package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"
	"time"
)

type searchReq struct {
	Path     string `json:"path"`
	Filename string `json:"filename"`
}

type searchRes struct {
	FilePath string
	FileSize int64
}

var matches int
var query = ".go"

var filepath string
var workerCount = 0
var maxWorkerCount = 32
var searchRequest = make(chan string)
var workerDone = make(chan bool)
var foundMatch = make(chan bool)

func main() {

	// 参考: https://blog.csdn.net/benben_2015/article/details/79134734
	for {
		fmt.Println("current time :", time.Now())
		searchOperator()
		time.Sleep(10 * time.Second)
	}
	// todo 守护进程
	// todo 配置文件法/args 选择搜索特定的目录和文件
	// todo 搜索到该文件 打印出该文件的目录地址
	// todo 搜索到文件，返回字符串
}

func searchOperator() {
	start := time.Now()
	workerCount = 1
	go search("D:\\workspace\\", true)
	waitForWorkers()
	fmt.Println(matches, "matches")
	fmt.Println("spent time :", time.Since(start))
	fmt.Println()
	matches = 0
}

func waitForWorkers() {
	for {
		select {
		case path := <-searchRequest:
			workerCount++
			go search(path, true)
		case <-workerDone:
			workerCount--
			if workerCount == 0 {
				return
			}
		case <-foundMatch:
			matches++
		}
	}
}

func search(path string, master bool) {
	sysType := runtime.GOOS
	files, err := ioutil.ReadDir(path)
	if err == nil {
		for _, file := range files {
			name := file.Name()
			if strings.Contains(name, query) {
				matches++
			}
			if file.IsDir() {
				if workerCount < maxWorkerCount {
					if sysType == "linux" {
						searchRequest <- path + name + "/"
					}
					if sysType == "windows" {
						searchRequest <- path + name + "\\"
					}
				} else {
					if sysType == "linux" {
						search(path+name+"/", false)
					}
					if sysType == "windows" {
						search(path+name+"\\", false)
					}
				}
			}
		}
		if master {
			workerDone <- true
		}
	}
}
