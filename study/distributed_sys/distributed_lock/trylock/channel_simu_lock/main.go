package main

import (
	"fmt"
	"sync"
)

/*
  trylock 尝试加锁，如果加锁失败的话也不会阻塞，而会直接返回锁的结果
  在go语言中我们可以用大小为1的channel来模拟trylock
*/

type Lock struct {
  c chan struct{}
}

// NewLock generate a try lock
func NewLock() Lock {
  var l Lock
  l.c = make(chan struct{},1)
  l.c <- struct{}{}
  return l
}

// Lock try lock,return lock result
func (l Lock) Lock() bool {
  lockResult := false
  select {
  case <- l.c:
    lockResult = true
  default:
  }
  return lockResult
}

// Unlock the try lock
func (l Lock) Unlock() {
  l.c <- struct{}{}
}

var counter int

func main()  {
  var l = NewLock()
  var wg sync.WaitGroup
  for i := 0; i < 10; i++ {
    wg.Add(1)
    go func()  {
      defer wg.Done()
      if !l.Lock() {
        println("lock failed")
        return
      }
      counter++
      println("current counter",counter)
      l.Unlock()
    }()
  }
  wg.Wait()
  fmt.Println("try lock test finished")
}
