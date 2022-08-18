package main

import (
	"fmt"
	"time"
)

func main() {
	numCh := make(chan int)
	timeoutCh := make(chan struct{})
	exceptionCh := make(chan string)
	go ExceptionSig(exceptionCh)
	go GenerateNum(numCh)
	for {
		select {
		case num := <-numCh:
			fmt.Println(num)
		case <-timeoutCh:
			fmt.Println("timeout")
		case <-exceptionCh:
			go func() {
				time.Sleep(10 * time.Second)
				timeoutCh <- struct{}{}
			}()
		}
	}
}

func GenerateNum(ch chan int) {
	for i := 0; i < 1000; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
}

func ExceptionSig(exceptionsigCh chan string) {
	fmt.Println("generate exception sig")
	time.Sleep(2 * time.Second)
	exceptionsigCh <- "exception"
}
