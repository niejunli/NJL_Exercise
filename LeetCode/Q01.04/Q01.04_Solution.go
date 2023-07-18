package Q0104

// 回文排列

func canPermutePalindrome(s string) bool {
	if len(s) == 0 {
		return false
	}
	resMap := make(map[byte]int64, len(s))
	var flag = 0
	for i := 0; i < len(s); i++ {
		if _, ok := resMap[s[i]]; !ok {
			resMap[s[i]] = 1
		} else {
			resMap[s[i]] += 1
		}
	}
	for _, v := range resMap {
		if v%2 == 1 {
			if flag == 0 {
				flag++
			} else {
				return false
			}
		}
	}
	return true
}
