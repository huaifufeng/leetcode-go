/**
  题目地址：https://leetcode-cn.com/problems/shortest-completing-word/

  解题：
    循环获取基准单词的字母字典，然后依次获取每个单词的字母字典，只要基准字典的每个字母在内部字母字典存在，并且个数不大于就符合，找最短的
*/
package questionBank

import (
	"strings"
)

func shortestCompletingWord(licensePlate string, words []string) string {
	chars := make(map[byte]int)
	licensePlate = strings.ToLower(licensePlate)
	for i := 0; i < len(licensePlate); i++ {
		if licensePlate[i] >= 'a' && licensePlate[i] <= 'z' {
			chars[licensePlate[i]]++
		}
	}
	if len(chars) == 0 {
		return ""
	}

	cur := ""
	for i := 0; i < len(words); i++ {
		wordChars := make(map[byte]int)
		if len(cur) > 0 && len(cur) <= len(words[i]) {
			continue
		}

		for j := 0; j < len(words[i]); j++ {
			wordChars[words[i][j]]++
		}

		flag := true
		for c, num := range chars {
			if wordChars[c] < num {
				flag = false
				break
			}
		}
		if flag {
			cur = words[i]
		}
	}

	return cur
}
