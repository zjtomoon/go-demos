package main

import (
	"fmt"
	"strings"
)

func main() {
	dataA := "JavaScript"
	dataB := "Golang"

	source := "this is a Golang"
	if strings.Contains(source,dataA){
		fmt.Println("JavaScript")
	} 
	if strings.Contains(source,dataB) {
		fmt.Println("Golang")
	}
	nums := [...]int{1,2,3}
	contains := Contains_Int(nums[:],2)
	fmt.Println(contains)
}

// 判断数组中是否包含目标元素
func Contains_Int(nums []int,target int) bool  {
	contains := make(map[int]bool)
	for i := 0 ; i < len(nums) ; i++ {
		contains[nums[i]] = true
		if _,ok := contains[target];ok {
			return true
		}
		// contains[target] = true
	}
	return false
}