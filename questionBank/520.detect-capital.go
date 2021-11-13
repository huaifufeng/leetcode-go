/**
  题目地址：https://leetcode-cn.com/problems/detect-capital/

  解题：
  1、循环这个字符串，O(n)
*/
package questionBank

import "unicode"

func detectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}

	needChar := 0
	if word[0] >= 'a' && word[0] <= 'z' {
		needChar = 1
	}

	for i := 1; i < len(word); i++ {
		if i == 1 && needChar == 0 {
			if word[i] >= 'a' && word[i] <= 'z' {
				needChar = 1
			} else {
				needChar = 2
			}

			continue
		}

		if needChar == 1 && (word[i] >= 'a' && word[i] <= 'z') {
			continue
		}

		if needChar == 2 && (word[i] >= 'A' && word[i] <= 'Z') {
			continue
		}

		return false
	}

	return true
}

func detectCapitalUse1(word string) bool {
	// 若第 1 个字母为小写，则需额外判断第 2 个字母是否为小写
	if len(word) >= 2 && unicode.IsLower(rune(word[0])) && unicode.IsUpper(rune(word[1])) {
		return false
	}

	// 无论第 1 个字母是否大写，其他字母必须与第 2 个字母的大小写相同
	for i := 2; i < len(word); i++ {
		if unicode.IsLower(rune(word[i])) != unicode.IsLower(rune(word[1])) {
			return false
		}
	}
	return true
}
