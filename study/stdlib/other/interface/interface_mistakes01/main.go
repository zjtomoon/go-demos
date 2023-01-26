package main

import "fmt"

func main() {
	// mistake1()
	mistake2()
	//TestInterface()
}

type ErrorImpl struct{}

func (e *ErrorImpl) Error() string {
	return ""
}

var ei *ErrorImpl

func ErrorImplFun() error {
	return ei
}

func mistake1() {
	f := ErrorImplFun()
	fmt.Println(f == nil) // false
}

type Person struct {
	Age int
}

func (p *Person) GetNum() int {
	return p.Age
}

func Add(s ...interface{}) int {
	total := 0
	for _, v := range s {
		total += v.(*Person).GetNum()
	}
	return total
}

func mistake2() {
	persons := []Person{
		{18}, {28},
	}
	var personArr []interface{}
	personArr = append(personArr, persons)
	fmt.Println(personArr...)
}

func TestInterface() {
	is := []int{7, 8, 9, 10}
	var mis []interface{}
	mis = append(mis, is)
	test(mis)
}

func test(i interface{}) {
	fmt.Println(i)
}
