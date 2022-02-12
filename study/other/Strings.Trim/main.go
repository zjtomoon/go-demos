package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// golang的strings包中的Trim的用法介绍
func main() {
	input := bufio.NewScanner(os.Stdin)

	fmt.Println("Please input your words:")

	for input.Scan() {
		result := input.Text()
		arr1 := strings.Trim(result, " ")
		arr2 := strings.Split(result, " ")
		fmt.Println("arr1 = ", arr1)
		fmt.Println("arr2 = ", arr2)
		fmt.Println(reflect.TypeOf(arr2))
		var arr3 []int

		//输入的字符串数字转为整形数组
		for i := 0; i < len(arr2); i++ {
			tmp, _ := strconv.Atoi(arr2[i])
			arr3 = append(arr3, tmp)
		}
		fmt.Println(arr3)
		fmt.Println("typeof arr3:", reflect.TypeOf(arr3))
	}
}
