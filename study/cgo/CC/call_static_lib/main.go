package main

/*
#cgo CXXFLAGS: -std=c++11
#cgo CFLAGS: -I./student
#cgo LDFLAGS: -L./student -lstu -lstdc++
#include <stdlib.h>
#include <stdio.h>
#include "interface.h"

 */
import "C"

import (
	"unsafe"
)

//必须添加 -lstdc++，否则编译出错

func main() {
	name := "test!"
	cStr := C.CString(name)
	defer C.free(unsafe.Pointer(cStr))
	obj := C.stuCreate()
	C.initName(obj,cStr)
	C.getStuName(obj)
}
