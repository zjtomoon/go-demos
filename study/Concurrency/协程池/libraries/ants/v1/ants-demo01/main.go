package main

import (
	"fmt"
	"github.com/panjf2000/ants"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Task() {
	fmt.Println("Hello World")
	time.Sleep(2 * time.Second)
}

func main() {
	defer ants.Release()
	pool, _ := ants.NewPool(5)
	task := func() {
		Task()
		wg.Done()
	}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		_ = pool.Submit(task)
	}
	wg.Wait()
}
