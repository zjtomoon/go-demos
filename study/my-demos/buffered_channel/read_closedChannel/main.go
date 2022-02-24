package main

import "fmt"

func main() {
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5

	//ch <- 6 // deadlock

	//defer close(ch) //deadlock
	close(ch)

	for val := range ch {
		fmt.Println(val)
	}
	//ch <- 6 //panic : send on closed channel
}
