package main

import "fmt"

// x 为反推出的兔子数，和代码中的x不一样
// (x/y+1)*(y+1)

func numRabbits(answers []int) int {
	var total int
	if len(answers) == 0 {
		return 0
	}
	count := make(map[int]int)
	for _,y := range answers {
		count[y]++ //记录相同的y出现的次数
	}
	fmt.Println(count)
	for y,x := range count {
		total += (x+y)/(y+1)*(y+1)
	}
	return total
}


func main()  {
	ans := []int{1,1,2}
	numRabbits(ans)
}