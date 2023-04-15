package Q42

// 双指针/单调栈 --- 接雨水

// 双指针
// 1. 对于某个下标i，能否在这个位置接上雨水，取决于i左右两侧的最大值(leftMax,rightMax)是否比height[i]大，只有当i左右两侧的最大值都比height[i]大时, 才能接i上的雨水，数量为min(leftMax, rightMax) - height[i];
// 限制当前位置接雨水的条件是其左右两侧最大值中的较小值min(leftMax, rightMax)
// 2. 对于左指针来说，leftMax是真实可信的, 因为leftMax的值是左指针一步一个脚印走出来的, 但是rightMax是不真实不可信的，因为i不知道从height[i]到height[j]之间是否有其他的数大于rightMax;
// 同样,对于右指针来说rightMax是真实可信的, leftMax值是不真实不可信的
//  对于左指针, 它右侧的真实的最大值 >= rightMax, 对于右指针, 它左侧的真实的最大值 >= leftMax
// 3. 对于一个位置来说影响它接水的应该是左右两侧最大值中的较小值, 也就是min(leftMax, rightMax)
// 左指针右侧的真实最大值会大于等于rightMax, 所以, 当出现leftMax < rightMax的时候, 左指针的位置是否能接雨水就已经确定了, 同样, 当出现leftMax >= rightMax的时候, 右指针的位置是否能接雨水就已经确定了

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func trap(height []int) int {
	left, right := 0, len(height)-1 // 双指针
	leftMax, rightMax, ans := 0, 0, 0
	for left < right {
		leftMax = Max(leftMax, height[left])
		rightMax = Max(rightMax, height[right])
		if leftMax < rightMax {
			ans += (leftMax - height[left])
			left += 1
		} else {
			ans += (rightMax - height[right])
			right -= 1
		}
	}
	return ans
}


// 单调栈
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func otherTrap(height []int) int {
	var ans int
	stack := []int{}
	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 { // 说明没有左边界
				break
			}
			left := stack[len(stack)-1]                     // 左边界
			curWidth := i - left - 1                        // i->右边界，left->左边界
			curHeight := Min(height[left], h) - height[top] // 高度取左右边界中的min，减去本身i的高度
			ans += (curWidth * curHeight)
		}
		stack = append(stack, i)
	}
	return ans
}
