package main

import (
	"fmt"

	"github.com/alitto/pond"
)

func main() {
	pool := pond.New(10, 1000)
	defer pool.StopAndWait()

	group := pool.Group()

	for i := 0; i < 20; i++ {
		n := i
		group.Submit(func() {
			fmt.Printf("Running group task #%d\n", n)
		})
	}
	group.Wait()
}
