package cmd

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"
	"sync"
	"time"
)

var matches int
var Query string

var workerCount = 0
var maxWorkerCount = 32
var searchRequest = make(chan string)
var workerDone = make(chan bool)
var foundMatch = make(chan bool)

var mtx sync.Mutex

func SearchOperator(path string) {
	start := time.Now()
	workerCount = 1
	go search(path, true)
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
			if strings.Contains(name, Query) {
				mtx.Lock()
				matches++
				mtx.Unlock()
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

// todo 搜索到该文件 打印出该文件的目录地址
// todo 搜索到文件，返回字符串
