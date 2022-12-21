package Q1

/*
	两数之和
	2层for循环的方法，复杂度O(n2)
	利用hash表的话，只需要使用一次for循环
*/

func twoSum(nums []int, target int) []int {
	resMap := map[int]int{}
	for i, x := range nums {
		if j, ok := resMap[target-x]; ok {
			return []int{i, j}
		}
		resMap[x] = i
	}
	return nil
}
