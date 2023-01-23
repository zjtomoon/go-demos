package main

import "sync"

type Person struct {
	Name string
	Age  int
}

var personalPool = sync.Pool {
	New: func() interface{} {
		return &Person{}
	},
}

func main()  {
	newPerson := personalPool.Get().(*Person)
	defer personalPool.Put(newPerson)
	newPerson.Name = "Jack"
}
