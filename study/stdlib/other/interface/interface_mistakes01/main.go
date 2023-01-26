package main

import (
	"fmt"
)

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
		age := v.(*Person).GetNum()
		total += age
	}
	return total
}

func Add2(s ...Person) int {
	total := 0
	for _, v := range s {
		total += v.GetNum()
	}
	return total
}

func mistake2() {
	persons := []Person{
		{18}, {28},
	}
	fmt.Println(persons)
	// fmt.Println(Add(persons...))
	fmt.Println(Add2(persons...))

	var personArr []interface{}
	personArr = append(personArr, persons)
	fmt.Println(personArr...)
	fmt.Println(personArr[0])
	//fmt.Println(Add(personArr...)) // interface {} is []main.Person, not *main.Person

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
