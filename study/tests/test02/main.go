package main

import (
	"fmt"
	"strconv"
)

func main() {
	id := 1
	name := "user1"
	var newName string
	newName = strconv.Itoa(id) + name
	fmt.Println(newName)
}
