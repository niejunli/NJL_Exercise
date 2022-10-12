package sort

// 冒泡排序/起泡排序 O(n*n)
// 多次循环：每次从前往后把大元素往后调，每次确定一个最大元素，多次后达到排序序列
// 比如第一次确定len(nums)-1位置的元素，即最大的；第二次确定len(nums)-2位置的元素，即第二大的
func mpSort(nums []int64) []int64 {
	// 外层循环，用来标记此次要确定哪个位置的值，内循环每结束一轮，都会确定一个位置上的元素
	for j := len(nums) - 1; j >= 0; j-- {
		// 从头开始遍历，将剩下元素中最大的那个挪至预设位置
		for i := 0; i < j; i++ {
			if nums[i] > nums[i+1] {
				tmp := nums[i]
				nums[i] = nums[i+1]
				nums[i+1] = tmp
			}
		}
	}
	return nums
}

// 快排是对冒泡排序的一种改进
func quickSort(list []int64, left, right int64) {
	var low = left
	var high = right
	// 边界判断，避免数组越界
	if low > high {
		return
	}
	// 每次要确定pivot元素应该所处的位置
	pivot := list[low] // 额外空间，左侧数据要比它小，右侧数据要比它大
	for low < high {
		for low < high && pivot <= list[high] { // 右侧找到第一个小于pivot的元素停止
			high--
		}
		list[low] = list[high]                 // 放到low位置
		for low < high && pivot >= list[low] { // low往右找到第一个大于pivot的元素，放到high位置
			low++
		}
		list[high] = list[low]
	}
	list[low] = pivot // 确定了pivot的位置
	// 左右递归分治求解
	quickSort(list, left, low-1)
	quickSort(list, low+1, right)
}

// 插入排序
func insertSort(nums []int64) []int64 {
	var team int64
	// nums[i] 即将插入的元素
	for i := 1; i < len(nums); i++ {
		team = nums[i] // 额外空间
		// 插入到前面已经有序的部分当中
		for j := i - 1; j >= 0; j-- {
			if team < nums[j] {
				// 向后移
				nums[j+1] = nums[j]
				nums[j] = team
			} else {
				break
			}
		}
	}
	return nums
}
