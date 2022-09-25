package Q26

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
