package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var counter int32
var mtx sync.Mutex

// r如果buffer为0，无法拿取最后一次循环的值
// 这里改为 bufferd 原因是最后一个循环把数送进ch后没有再一个循环从ch拿值
//加上bufferd 最后ch存有一个值， 这样可以结束goroutine 外面从里面拿出那个结果
var ch = make(chan int, 1)

func UnsafeIncCounter() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		//counter++
		// 实际上是按照以下逻辑执行的
		temp := counter
		temp = temp + 1
		counter = temp
	}
}

func raceConditionDemo() {
	wg.Add(2)

	go UnsafeIncCounter()
	go UnsafeIncCounter()

	wg.Wait()

	fmt.Println(counter)
}

// 互斥锁
func MutexIncCounter() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		mtx.Lock()
		counter++
		mtx.Unlock()
	}
}

func mutexDemo() {
	wg.Add(2)

	go MutexIncCounter()
	go MutexIncCounter()

	wg.Wait()

	fmt.Println(counter)
}

// 原子操作

func AtomicIncCounter() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		atomic.AddInt32(&counter, 1)
	}
}

func atomicDemo() {
	wg.Add(2)

	go AtomicIncCounter()
	go AtomicIncCounter()

	wg.Wait()
	fmt.Println(counter)
}

func ChannelIncCounter() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		count := <-ch
		count++
		//fmt.Println(count)
		ch <- count
	}
}

func channelDemo() {
	wg.Add(2)

	go ChannelIncCounter()
	go ChannelIncCounter()

	ch <- 0

	wg.Wait()

	fmt.Println(<-ch)
}

func main() {
	//raceConditionDemo() // 结果一直变

	//mutexDemo() // 20000

	//atomicDemo() // 20000

	channelDemo()
}
