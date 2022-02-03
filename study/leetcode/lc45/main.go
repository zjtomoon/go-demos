package main

// 跳跃游戏II
func jump(nums []int) int {
	var step int
	var curDistance int
	var nextDistance int
	if len(nums) <= 1 {
		step = 0
	}
	for i := 0; i < len(nums) - 1; i++ { // 循环到 len(nums)-2
		nextDistance = max(nextDistance, i+nums[i]) //更新下一步覆盖的
		if i == curDistance { //初始位置 i = 0 curDistance = 0
			curDistance = nextDistance
			step++
		}
	}
	return step
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
