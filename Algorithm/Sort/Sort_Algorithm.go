package sort

// 选择排序 O(n*n)
// 遍历数组， 从中选择最小元素，将它与数组的第一个元素交换位置
// 继续从数组剩下的元素中选择出最小的元素，将它与数组的第二个元素交换位置
// 循环以上过程，直到将整个数组排序
func selectSort(nums []int64) []int64 {
	// 外层循环，用来标记此次要确定哪个位置的元素
	for i := 0; i < len(nums); i++ {
		min := i
		// 找到剩余元素中最小的元素
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		// 挪到第i个位置上，即确定好了第i个位置的元素
		tmp := nums[i]
		nums[i] = nums[min]
		nums[min] = tmp
	}
	return nums
}

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

// 选择排序是通过直接遍历剩下的元素，找到其中最小的元素来进行交换；冒泡排序是通过交换相邻的两个元素，每次将符合条件的元素进行移动来达到目的

// 快排是对冒泡排序的一种改进
func quickSort(list []int64, left, right int64) {
	var low = left
	var high = right
	// 边界判断，避免数组越界
	if low > high {
		return
	}
	// 每次要确定pivot元素应该所处的位置
	pivot := list[low] // 辅助额外空间，左侧数据要比它小，右侧数据要比它大
	for low < high {
		// 以下也相当于一个交换的过程
		for low < high && list[high] >= pivot { // 右侧找到第一个小于pivot的元素停止
			high--
		}
		list[low] = list[high]                 // 放到low位置
		for low < high && list[low] <= pivot { // low往右找到第一个大于pivot的元素，放到high位置
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
// 数组看成两部分：已排序部分和未排序部分。排序部分从第一个元素开始，该元素可以认为已经被排序。
// 遍历数组，每次将扫描到的元素，插入到有序部分的适当位置
func insertSort(nums []int64) []int64 {
	var temp int64
	// nums[i] 为即将插入的元素
	for i := 1; i < len(nums); i++ {
		temp = nums[i] // 额外空间
		// 插入到前面已经有序的部分当中
		for j := i - 1; j >= 0; j-- {
			if temp < nums[j] {
				// 向后移
				nums[j+1] = nums[j]
				nums[j] = temp
			} else {
				break
			}
		}
	}
	return nums
}

// 希尔排序，实质上是插入排序的优化(分组插入排序)
// 对于大规模的数组，插入排序很慢：只能交换相邻的元素位置，每次只能将未排序序列数量减1
// 希尔排序通过交换不相邻的元素位置，每次可以将未排序序列的减少数量变多
// 希尔排序使用插入排序对间隔d的序列进行排序，通过不断减小d。最后令d=1，就可以使得整个数组是有序的
func ShellSort(nums []int64) {
	// 数组长度
	n := len(nums)
	// 步长每次减半，直到步长为1
	for step := n / 2; step > 0; step /= 2 {
		// 开始插入排序，每一轮的步长为step(i指向未排序的部分)
		for i := step; i < n; i += step {
			for j := i - step; j >= 0; j -= step {
				if nums[j] > nums[j+step] {
					nums[j], nums[j+step] = nums[j+step], nums[j]
					continue
				}
				break
			}
		}
	}
}

// 归并排序 分治法的典型应用
func merge(a []int64, b []int64) []int64 {
	var r = make([]int64, len(a)+len(b))
	var i, j = 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			r[i+j] = a[i]
			i++
		} else {
			r[i+j] = b[j]
			j++
		}
	}
	for i < len(a) {
		r[i+j] = a[i]
		i++
	}
	for j < len(b) {
		r[i+j] = b[j]
		j++
	}
	return r
}
func MergeSort(items []int64) []int64 {
	if len(items) < 2 {
		return items
	}
	var middle = len(items) / 2
	var a = MergeSort(items[:middle])
	var b = MergeSort(items[middle:])
	return merge(a, b)
}

// 堆排序
func HeapSort(nums []int64) {
	length := len(nums)
	// 构建大顶堆
	for i := (length - 2) / 2; i >= 0; i-- {
		justDown(nums, i, length)
	}
	// 从最后一个元素开始交换
	end := length - 1
	for end != 0 {
		nums[0], nums[end] = nums[end], nums[0]
		justDown(nums, 0, end) // 每次交换后，最后一个元素即为最大值，然后调整顶堆结构
		end--
	}
}
func justDown(nums []int64, root, length int) {
	parent := root
	// 左孩子
	child := parent*2 + 1
	for child < length {
		// 比较左孩子和右孩子大小，指向较大的孩子
		if child+1 < length && nums[child+1] > nums[child] {
			child++
		}
		// 孩子和父节点是否交换
		if nums[child] > nums[parent] {
			nums[child], nums[parent] = nums[parent], nums[child]
			parent = child // 交换完后，指向下一层级的父节点
			child = parent*2 + 1
		} else {
			break
		}
	}
}
