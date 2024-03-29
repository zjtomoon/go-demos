package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/panjf2000/ants"
)

var sum int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum,n)
	fmt.Printf("run with %d\n",n)
}

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

func main() {
	defer ants.Release()

	runTimes := 1000

	var wg sync.WaitGroup

	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		ants.Submit(syncCalculateSum)
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n",ants.Running())
	fmt.Printf("finish all tasks.\n")

	p,_ := ants.NewPoolWithFunc(10,func(i interface{}){
		myFunc(i)
		wg.Done()
	})

	defer p.Release()
	for i := 0 ; i < runTimes; i++ {
		wg.Add(1)
		p.Invoke(int32(i))
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n",p.Running())
	fmt.Printf("finish all tasks,result is %d\n",sum)
}
