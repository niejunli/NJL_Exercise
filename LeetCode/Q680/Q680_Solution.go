package Q680

/*
双指针 --- 验证回文串2⃣️
*/

func validPalindrome(s string) bool {
	return helper(s, "left") || helper(s, "right")
}

func helper(s, deleteSide string) bool {
	var left, right = 0, len(s) - 1
	flag := true // 标识是否已经删除了字符
	for left < right {
		if s[left] == s[right] {
			left += 1
			right -= 1
		} else { // 在不满足回文的条件时
			if flag { // 还未删除字符
				if deleteSide == "left" {
					left += 1
					flag = false
				} else if deleteSide == "right" {
					right -= 1
					flag = false
				}
			} else { // 已删除了一次，不允许再删除
				return false
			}
		}
	}
	return true
}
