//题目地址：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
package 题库

import (
	"math"
)

//滑动窗口优化
func lengthOfLongestSubstring(s string) int {
	length := len(s)
	maxLength := 0
	hasMap := make(map[uint8]int)

	i := 0

	for j := 0; j < length; j++ {
		if _, ok := hasMap[s[j]]; ok {
			i = int(math.Max(float64(hasMap[s[j]]), float64(i)))
		}
		maxLength = int(math.Max(float64(maxLength), float64(j - i + 1)))
		hasMap[s[j]] = j + 1
	}

	return maxLength
}

//滑动窗口
func lengthOfLongestSubstring2(s string) int {
	maxLength := 0
	length := len(s)
	hasMap := make(map[uint8]int)

	i, j := 0, 0

	for i < length && j < length {
		if _, ok := hasMap[s[j]]; !ok {
			hasMap[s[j]] = j
			j++
			maxLength = int(math.Max(float64(maxLength), float64(j - i)))
		} else {
			delete(hasMap, s[i])
			i++
		}
	}

	return maxLength
}

//循环查找
func lengthOfLongestSubstring1(s string) int {
	sLength := len(s)
	if sLength == 0 {
		return 0
	}

	maxLength := 0
	theLength := 0
	hasMap := make(map[rune]int, sLength)
	hasSlice := make([]rune, sLength)

	for _, val := range s {
		if value, ok := hasMap[val]; ok {
			theLength = len(hasMap)
			if theLength > maxLength {
				maxLength = theLength
			}

			hasMap = make(map[rune]int, sLength)
			hasSlice = hasSlice[value+1:]
			for _, hasVal := range hasSlice {
				if hasVal == 0 {
					break
				}

				hasMap[hasVal] = len(hasMap)
			}

			hasMap[val] = len(hasMap)
			hasSlice[hasMap[val]] = val
		} else {
			hasMap[val] = len(hasMap)
			hasSlice[hasMap[val]] = val
		}
	}

	theLength = len(hasMap)
	if theLength > maxLength {
		maxLength = theLength
	}

	return maxLength
}
