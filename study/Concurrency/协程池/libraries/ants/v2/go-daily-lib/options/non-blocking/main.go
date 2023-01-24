package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"time"
)

func wrapper(i int, wg *sync.WaitGroup) func() {
	return func() {
		fmt.Printf("hello from task:%d\n", i)
		time.Sleep(1 * time.Second)
		wg.Done()
	}
}

func main() {
	p, _ := ants.NewPool(2, ants.WithNonblocking(true))
	defer p.Release()
	var wg sync.WaitGroup
	wg.Add(8)
	for i := 0; i < 8; i++ {
		err := p.Submit(wrapper(i, &wg))
		if err != nil {
			fmt.Printf("task:%d err:%v\n", i, err)
			wg.Done()
		}
	}
	wg.Wait()
}
