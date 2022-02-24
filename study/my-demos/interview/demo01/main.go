package main

import (
	"fmt"
	"time"
)

// 20220224面试
func main() {
	//c := make(chan int)
	//
	//go func() {
	//	<- c
	//	fmt.Println("A")
	//}()
	//
	//time.Sleep(2 * time.Second)
	//
	//close(c)
	//
	//time.Sleep(3 * time.Second)
	//fmt.Print("B")

	c := make(chan struct{})
	go func() {
		c <- struct{}{}
		fmt.Println("A")
	}()

	time.Sleep(2 * time.Second)

	<-c
	close(c)

	time.Sleep(3 * time.Second)
	fmt.Println("B")
}
