package main


// pre[i] == pre[i-1] + nums[i]
// pre[i+1] == pre[i] + nums[i+1]

// pre[i-1] + nums[i] + (pre[n]-pre[i]) = total
//满足条件时
// 2*pre[i-1] + nums[i] = total

func privotIndex(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var total int
	//求数组总和
	for _,v := range nums {
		total += v
	}
	var pre int
	for i := 0 ; i < len(nums); i++ {
		// pre += nums[i]
		if pre *2 + nums[i] == total {
			return i
		}
		//先判断后求和
		pre += nums[i]
	}
	return -1
}