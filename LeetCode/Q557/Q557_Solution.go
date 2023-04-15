package Q557

import "strings"

/*
字符串 --- 反转字符串中的单词 III
*/

// 使用额外空间
func reverseWords1(s string) string {
	length := len(s)
	ret := []byte{}
	for i := 0; i < length; {
		start := i
		for i < length && s[i] != ' ' {
			i++
		}
		p := start
		for temp := 0; temp < (i - p); temp++ {
			ret = append(ret, s[i-temp-1])
		}
		for i < length && s[i] == ' ' {
			i++
			ret = append(ret, ' ')
		}
	}
	return string(ret)
}

// 原地翻转
func reverseWords2(s string) string {
	s1 := strings.Split(s, " ")
	for k, _ := range s1 {
		s1[k] = string(reverse([]byte(s1[k])))
	}
	return strings.Join(s1, " ")
}

func reverse(s []byte) []byte {
	first, last := 0, len(s)-1
	for last > first {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
	return s
}
