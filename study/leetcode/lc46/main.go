package main

// var res [][]int
// func permute(nums []int) [][]int {
// 	res = [][]int{}
// 	backTrace()
// 	return res
// }

// func backTrace(nums []int, numsLen int, path []int) {
// 	if len(nums) == 0 {
// 		p := make([]int, len(path))
// 		copy(p,path)
// 		res = append(res, p)
// 	}	
// 	for i := 0; i < numsLen; i++ {
// 		cur := nums[i]
// 		path = append(path, cur)
// 		nums = append(nums[:i],nums[i+1:]...)
// 		backTrace(nums,len(nums),path)
// 		nums = append(nums[:i],append([]int{cur},nums[i:]...)...)
// 		path = path[:len(path)-1]
// 	}
// }

/* func permute(nums []int) [][]int {
//最终要返回的二维数组
	var res [][]int
	//已经用过的节点存储用的切片
	var path []int
	//将用过节点进行标记的哈希表
	visited :=make(map[int]bool)
	size :=len(nums)
	var backTrack func()
	backTrack = func() {
		//递归终止条件
		//也就是nums里的元素都用到了
		if len(path) ==size{
			//temp暂时存放path，path的长度肯定是nums的长度
			temp :=make([]int,size)
			copy(temp,path)
			res = append(res,temp)
			return
		}
		//从0开始所以不去等
		for i:=0;i<size;i++{
			//一个排列结果（path）里的一个元素只能使用一次
			//相当于查map里有没有这个元素，有就continue跳出
			if visited[nums[i]] {
				continue
			}
			//第一次出现就给他打个标记true
			visited[nums[i]] =true
			//将这个元素放入path路径中
			path = append(path,nums[i])
			//递归
			backTrack()
			//回溯
			visited[nums[i]] =false
			//就是吐出最后一个元素
			path = path[:len(path)-1]
		}

	}
	backTrack()
	return res
} */
