package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	count := 10000000
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			fmt.Printf("num = %d\n", i)
		}()
	}
	wg.Wait()
}
