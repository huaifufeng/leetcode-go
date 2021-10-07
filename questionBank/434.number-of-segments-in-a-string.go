/**
  题目地址：https://leetcode-cn.com/problems/number-of-segments-in-a-string/

  解题：
  1、循环字符串，比较字符和空字符，如果前一个字符不是空字符，当前是空字符，加一，考虑末尾是空字符和不是空字符的情况， O(n)
*/
package questionBank

func countSegments(s string) int {
	lastChar := ' '
	num := 0
	for _, v := range s {
		if v == ' ' && lastChar != ' ' {
			lastChar = v
			num++
			continue
		}
		lastChar = v
	}

	if lastChar != ' ' {
		num++
	}

	return num
}
