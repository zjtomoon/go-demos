package main

/*
#cgo CFLAGS: -I./student
#cgo LDFLAGS: -L./student -lstu
#include <stdlib.h>
#include <stdio.h>
#include "interface.h"

 */
import "C"

import (
	"unsafe"
)

func main() {
	name := "test!"
	cStr := C.CString(name)
	defer C.free(unsafe.Pointer(cStr))
	obj := C.stuCreate()
	C.initName(obj,cStr)
	C.getStuName(obj)
}
