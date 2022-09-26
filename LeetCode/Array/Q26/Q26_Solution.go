package Q26

/*
	快慢指针
	利用数组有序的特性。同样是慢指针左侧维护满足条件的结果集，快指针用来遍历数组
*/

func removeDuplicates(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	var l = 0
	for r := 1; r < len(nums); r++ {
		if nums[r] != nums[l] {
			l++
			nums[l] = nums[r]
		}
	}
	return l + 1
}
