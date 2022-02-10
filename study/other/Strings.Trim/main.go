package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// golang的strings包中的Trim的用法介绍
func main()  {
  input := bufio.NewScanner(os.Stdin)

  fmt.Println("Please input your words:")

  for input.Scan() {
    result := input.Text()
    arr1 := strings.Trim(result,";")
    arr2 := strings.Split(result,";")
    fmt.Println("arr1 = ",arr1)
    fmt.Println("arr2 = ",arr2)
  }
}
