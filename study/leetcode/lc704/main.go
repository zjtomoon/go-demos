package main


//二分法查找
func search(nums []int, target int) int {
    low,high := 0,len(nums)-1
    // if len(nums) == 0 {
    //     return -1
    // }
    // if len(nums) == 1 {
    //     if nums[0] == target {
    //         return 0
    //     } else {
    //         return -1
    //     }
    // }
    // if len(nums) == 2{
    //     if nums[0] == target {
    //         return 0
    //     } else if nums[1] == target {
    //         return 1
    //     } else {
    //         return -1
    //     }
    // }

    for low <= high { //必须是小于等于
        mid := (low + high)/2
        if nums[mid] < target {
            low = mid + 1
        } else if nums[mid] > target {
            high = mid -1
        } else {
            return mid
        }
    }
    return -1
}
