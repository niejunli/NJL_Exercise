package Q27

/*
	同样是快慢指针
	慢指针左侧为符合条件的结果集，快指针用来遍历数组
 */

func removeElement(nums []int, val int) int {
	if len(nums) == 0 || (len(nums) == 1 && nums[0] != val) {
		return len(nums)
	}
	if len(nums) == 1 && nums[0] == val {
		return 0
	}
	var l = 0
	for r := 0; r < len(nums); r++ {
		if nums[r] != val {
			nums[l] = nums[r]
			l++
		}
	}
	return l
}
