package main

import "fmt"

// 给你一个整数数组 nums 和一个整数 k ，请你统计并返回该数组中和为 k 的连续子数组的个数。

//pre[i] == pre[i-1] + nums[i]

// [j..i]子数组和为k
// pre[i] - pre[j-1] == k
// pre[j-1] == pre[i] - k
 
func subarraySum(nums []int,k int) int {
	//枚举法
	// var count int
	// for i := 0; i < len(nums); i++{
	// 	var sum int
	// 	for j := i; j >= 0; j-- {
	// 		sum += nums[j]
	// 		if sum == k {
	// 			count++
	// 		}
	// 	}
	// }
	// return count

	// 前缀和 + 哈希优化
	// var pre int
	// var count int
	// m := make(map[int]int, 0)
	// // 0 的前缀和是1
	// m[0] = 1
	// for i := 0 ; i < len(nums); i++ {
	// 	pre += nums[i]
	// 	if _,ok := m[pre-k];ok  {
	// 		count += m[pre-k]
	// 	}
	// 	m[pre] += 1
	// }
	// return count

	var pre int
	var count int
	m := make(map[int]int, 0)
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		fmt.Println("pre = ",pre)
		count += m[pre-k]
		fmt.Println("m[pre-k] = ",m[pre-k])
		m[pre]++
		fmt.Println("m[pre] = ",m[pre])
	}
	fmt.Println("m = ",m)
	return count
}

func main()  {
	nums := [...]int{2,2,3,4,5,6,7}
	count := subarraySum(nums[:],7)
	fmt.Println(count)
}