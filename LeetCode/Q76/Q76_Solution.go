package Q76

/*
	双指针 --- 滑动窗口 --- 最小覆盖子串
*/

import (
	"math"
)

func minWindow(s string, t string) string {
	// need window相当于计数器，用来记录t中字符出现的次数和s窗口中相应字符出现的次数
	need, window := map[byte]int{}, map[byte]int{}
	// t中各字符出现次数
	for i := range t {
		need[t[i]]++
	}
	// left, right用来遍历s
	left, right := 0, 0
	// valid用来记录s窗口内，包含t中验证通过(出现次数)的字符个数
	valid := 0
	// 用来记录符合条件的字符串的起始位置以及长度
	index, length := 0, math.MaxInt32

	// 开始遍历字符串s
	for right < len(s) {
		// 滑动窗口的右限
		tempAdd := s[right]
		right++
		// 如果遍历的这个字符，符合条件(出现在t中)
		if _, ok := need[tempAdd]; ok {
			window[tempAdd]++
			if window[tempAdd] == need[tempAdd] {
				// 验证通过的字符 个数+1
				valid++
			}
		}
		// s窗口中的字符串已经包含了t，这个时候需要进行优化，寻求最短的字符串
		for valid == len(need) {
			// 更新结果
			if right-left < length {
				index = left
				length = right - left
			}
			// 滑动窗口的左限
			tempDel := s[left]
			left++
			if _, ok := need[tempDel]; ok {
				window[tempDel]--
				if window[tempDel] < need[tempDel] {
					valid--
				}
			}
		}
	}
	if length == math.MaxInt32 {
		return ""
	}
	return s[index : index+length]
}
