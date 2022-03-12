package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func player(name string, ch chan int) {
	defer wg.Done()

	for {
		ball, ok := <-ch
		if !ok {
			fmt.Printf("channel is closed %s wins!\n", name)
			return
		}

		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100)

		if n%10 == 0 { // 模拟十分之一概率
			// 把球打飞，用关闭通道模拟
			close(ch)
			fmt.Printf("%s missed the ball! %s losses\n", name, name)
			return
		}
		ball++
		fmt.Printf("%s receives ball %d \n ", name, ball)
		ch <- ball
	}

}

func main() {
	wg.Add(2)

	ch := make(chan int, 0)

	go player("alex", ch)
	go player("william", ch)

	ch <- 0

	wg.Wait()
}
