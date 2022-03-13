package main

import (
	"fmt"
	"reflect"
)

// setability modify reflection objects
func main() {
	fVar := 3.14
	v := reflect.ValueOf(fVar)
	fmt.Printf("is float canSet: %v canAddr: %v\n", v.CanSet(), v.CanAddr())
	//v.SetFloat(221232.12)
	vp := reflect.ValueOf(&fVar)
	fmt.Printf("is float canSet: %v canAddr %v\n", vp.Elem().CanSet(), vp.Elem().CanAddr())
	vp.Elem().SetFloat(2.3333)

	fmt.Println(fVar)
}
