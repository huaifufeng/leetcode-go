//题目地址：https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/5/strings/36/
package 字符串

import (
	"strings"
)

func isPalindrome(s string) bool {
	length := len(s)
	if length == 0 {
		return true
	}

	s = strings.ToLower(s)
	index := 0
	sSlice := make([]uint8, length)
	for i := 0; i < length; i++ {
		if s[i] >= '0' && s[i] <= '9' {
			sSlice[index] = s[i]
			index++
			continue
		}

		if s[i] >= 'a' && s[i] <= 'z' {
			sSlice[index] = s[i]
			index++
			continue
		}
	}

	sSlice = sSlice[:index]
	i := 0
	j := len(sSlice) - 1
	for i < j {
		if sSlice[i] != sSlice[j] {
			return false
		}
		i++
		j--
	}

	return true
}
