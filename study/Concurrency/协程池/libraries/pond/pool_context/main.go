package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/alitto/pond"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	pool := pond.New(1, 100, pond.Context(ctx))
	defer pool.StopAndWait()

	var count int = 100
	for i := 0; i < count; i++ {
		n := i
		pool.Submit(func() {
			fmt.Printf("Task #%d started\n", n)
			time.Sleep(1 * time.Second)
			fmt.Printf("Task #%d finished\n", n)
		})
	}
}
