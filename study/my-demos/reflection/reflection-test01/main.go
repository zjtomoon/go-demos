package main

import (
	"fmt"
	"reflect"
)

// interface values to reflection objects
func printMeta(obj interface{}) {
	t := reflect.TypeOf(obj)
	n := t.Name()
	k := t.Kind()
	v := reflect.ValueOf(obj)
	fmt.Printf("Type: %s Type.Name: %s Kind: %s Value: %s Value: %v\n", t, n, k, v)
}

type Handler func(int, int) int

func main() {
	var intVar int64 = 10
	stringVar := "hello"
	type book struct {
		name  string
		pages int
	}
	sum := func(a, b int) int {
		return a + b
	}

	var sub Handler = func(a int, b int) int {
		return a - b
	}

	sli := make([]int, 5)

	printMeta(intVar)
	printMeta(stringVar)
	printMeta(book{
		name:  "harry potter",
		pages: 600,
	})
	printMeta(sum)
	printMeta(sub)
	printMeta(sli)
}
