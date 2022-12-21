package Q26

/*
	双指针 --- 删除有序数组中的重复项
	利用数组有序的特性。其中一个指针左侧维护满足条件的结果集，另一个指针用来遍历数组
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
