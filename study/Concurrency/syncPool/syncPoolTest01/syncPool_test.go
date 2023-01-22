package main

import (
	"sync"
	"testing"
)

// type Person struct {
// 	Age int
// }

// Benchmark 命令： go test -bench=. -run=none

var (
	personPool = sync.Pool{
		New: func() interface{} {
			return &Person{}
		},
	}
)

func BenchmarkWithoutPool(b *testing.B) {
	var p *Person
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			p = new(Person)
			p.Age = 23
		}
	}
}

func BenchmarkWithPool(b *testing.B) {
	var p *Person
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0 ; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			p = personPool.Get().(*Person)
			p.Age = 23
			personPool.Put(p)
		}
	}
}