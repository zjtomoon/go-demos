package main

import "fmt"

// Mutex 使用channel 实现互斥锁
type Mutex struct {
	ch chan struct{}
}

// NewMutex 初始化锁
func NewMutex() *Mutex {
	mu := &Mutex{make(chan struct{},1)}
	return mu
}

// Lock 获取锁
func (m *Mutex) Lock() bool {
	select {
	case m.ch <- struct{}{}:
		return true
	default:
		return false
	}
}

// Unlock 释放锁
func (m *Mutex) Unlock() bool {
	select {
	case <- m.ch:
		return true
	default:
		panic("unlock of unlocked mutex")
	}
}

// IsLocked 查询锁状态
func (m *Mutex) IsLocked() bool {
	return len(m.ch) == 1
}

func main() {
	m := NewMutex()
	ok := m.Lock()
	fmt.Printf("locked is %v\n",ok)

	ok = m.IsLocked()
	fmt.Printf("locked is %v\n",ok)

	m.Unlock()
	ok = m.IsLocked()
	fmt.Printf("locked v %v\n",ok)
}
