package main

import (
	"fmt"

	"github.com/alitto/pond"
)

func main() {
	pool := pond.New(100, 1000)

	for i := 0; i < 1000; i++ {
		n := i
		pool.Submit(func() {
			fmt.Printf("Running task #%d\n", n)
		})
	}

	pool.StopAndWait()
}
