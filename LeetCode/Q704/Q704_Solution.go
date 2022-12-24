package Q704

/*
	左右指针 --- 二分查找
*/

func search(nums []int, target int) int {
    left, right := 0, len(nums)-1
    for left <= right {
        mid := (left + right) / 2
        if nums[mid] == target {
            return mid
        }
        if nums[mid] < target {
            left = mid + 1
        }
        if nums[mid] > target {
            right = mid - 1
        }
    }
    return -1
}