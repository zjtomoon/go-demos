package main

import (
	"encoding/json"
	"fmt"
	"sort"
)
//接下来如何移动left 和right呢， 如果nums[i] + nums[left] + nums[right] > 0 就说明 
//此时三数之和大了，因为数组是排序后了，所以right下标就应该向左移动，这样才能让三数之和小一些。

// 如果 nums[i] + nums[left] + nums[right] < 0 说明 此时 三数之和小了，left 就向右移动，才能让三数之和大一些，直到left与right相遇为止。
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i:= 0 ; i < len(nums); i++ {
		n1 := nums[i]
		if n1 > 0 {
			break
		}
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		l,r := i+1,len(nums)-1
		for l < r {
			n2,n3 := nums[l],nums[r]
			if n1 + n2 + n3 == 0 {
				res = append(res, []int{n1,n2,n3})
				for l < r && n2 == nums[l+1] {
					l++
				}
				for l < r && n3 == nums[r-1] {
					r--
				}
				l++
				r--
			}else if n1 + n2 + n3 < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}

// 二维数组去重
//基于 map 和 json 序列化实现了二维数组的去重
func differ(nums [][]int) [][]int {
	 var arr [][]int
	 had := make(map[string]struct{})
	for i := 0; i < len(nums); i++ {
		marshal,_ := json.Marshal(nums[i])
		had[string(marshal)] = struct{}{}
	}
	for key,_ := range had {
		var item []int
		json.Unmarshal([]byte(key),&item)
		arr = append(arr, item)
	}
	return arr
}

// 一维数组去重
// func differ(nums []int) []int {
// 	var arr []int
// 	had := make(map[interface{}]bool)
// 	for i := 0 ; i < len(nums) ; i++ {
// 		if _,ok := had[nums[i]];!ok {
// 			arr = append(arr, nums[i])
// 			had[nums[i]] = true
// 		}
// 	}
// 	return arr
// }

//先排序
func QuickSort(nums []int,start int,end int) []int {
	s,e := start,end
	temp := 0
	privot := nums[(start + end)/2]
	for s < e {
		for nums[s] < privot {
			s++
		}
		for nums[e] > privot {
			e--
		}
		if s >= e {
			break
		}
		//交换
		temp = nums[s]
		nums[s] = nums[e]
		nums[e] = temp

		// 表示左边排序完毕，
		if nums[s] == privot {
			e--
		}
		//表示右边排序完毕
		if nums[e] == privot {
			s++
		}

		if s == e {
			s++
			e--
		}
		if start < e {
			sort(nums,start,e)
		}
		if end > s {
			sort(nums,s,end)
		}
	}
	return nums
}


func main()  {
	a := []int{1,2,3,4,5,6,5,6,7,0,4,3}
	start,end:= 0,len(a)-1
	nums := sort(a,start,end)
	fmt.Println(nums)
	diff := differ(nums)
	fmt.Println(diff)
}