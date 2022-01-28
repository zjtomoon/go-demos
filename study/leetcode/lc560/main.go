package main

// 给你一个整数数组 nums 和一个整数 k ，请你统计并返回该数组中和为 k 的连续子数组的个数。


//pre[i] == pre[i-1] + nums[i]

// [j..i]子数组和为k
// pre[i] - pre[j-1] == k
 
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


}