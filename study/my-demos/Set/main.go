package main

import "sync"

type Set struct {
	m map[int]bool
	sync.Mutex
}

func New() *Set {
	return &Set{
		m: map[int]bool{},
	}
}

func (s *Set) Add(item int) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

// 模拟实现set
func main() {

}
