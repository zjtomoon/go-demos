package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name string
}

// reflection objects to interface values
func main() {

	// first: show convert reflect.value to float
	floatVar := 3.14
	v := reflect.ValueOf(floatVar)
	newFloat := v.Interface().(float64)
	fmt.Println(newFloat + 1.2)

	// second: show convert reflect.value to slice
	sliceVar := make([]int, 5)
	v = reflect.ValueOf(sliceVar)

	// method 1
	v = reflect.Append(v, reflect.ValueOf(2))

	// method 2
	newSlice := v.Interface().([]int)
	newSlice = append(newSlice, 4)
	fmt.Println(newSlice)

	//third: show convert reflect.value to student
	stuPtr := reflect.New(reflect.TypeOf(student{}))
	stu := stuPtr.Elem()
	nameField := stu.FieldByName("Name")
	if nameField.IsValid() {
		if nameField.CanSet() {
			nameField.SetString("alex")
		}
		realStudent := stu.Interface().(student)
		fmt.Println(realStudent)
	}
}
