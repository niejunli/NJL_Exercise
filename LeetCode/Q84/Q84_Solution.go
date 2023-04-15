package Q84

/*
单调栈 --- 柱状图中最大的矩形
*/

func largestRectangleArea(heights []int) int {
	var res int
	var stack = make([]int, 0) // 栈
	var nums = make([]int, 0)

	// 前后加个哨兵
	// 不加0的话，导致栈中会一直有有效元素没有弹干净，因此需要在首尾加0，使得原数列中的所有有效元素均被弹出，从而获取最大值。
	nums = append(nums, []int{0}...) // 头部加0是为了便于处理当栈里只有一个有效元素要弹出的时候计算面积
	nums = append(nums, heights...)
	nums = append(nums, []int{0}...) // 尾部的0是为了将栈中元素全部push出来进行遍历

	for i := 0; i < len(nums); i++ {
		// 栈顶元素a拿出来 和 数组中遍历到的元素b 相比较
		for len(stack) > 0 && nums[stack[len(stack)-1]] > nums[i] {
			// 以第[stack[len(stack)-1]根柱子为最矮柱子所能延伸的最大面积
			tmp := stack[len(stack)-1]      // 取栈顶元素
			stack = stack[0 : len(stack)-1] // 更新栈
			res = Max(res, nums[tmp]*(i-stack[len(stack)-1]-1))
		}
		stack = append(stack, i)
	}
	return res
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
