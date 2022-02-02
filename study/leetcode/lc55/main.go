package main

// 跳跃游戏
func canJump(nums []int) bool  {
	var cover int // 表示当前位置能覆盖的长度
	if len(nums) == 1 {
		return true
	}
	for i := 0 ; i <= cover; i++ {
		cover = max(i + nums[i],cover)
		if cover >= len(nums)-1 {
			return true
		}
	}
	return false
}

func max(a,b int) int {
	if a > b {
		return a 
	}
	return b
}