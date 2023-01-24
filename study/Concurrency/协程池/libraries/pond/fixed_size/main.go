package main

import (
	"fmt"

	"github.com/alitto/pond"
)

func main() {
	pool := pond.New(10, 0, pond.MinWorkers(10))
	for i := 0; i < 1000; i++ {
		n := i
		pool.Submit(func() {
			fmt.Printf("Running task #%d\n", n)
		})
	}
	pool.StopAndWait()
}
