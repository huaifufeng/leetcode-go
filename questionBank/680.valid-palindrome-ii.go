/**
  题目地址：https://leetcode-cn.com/problems/valid-palindrome-ii/

  给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。

  <pre>
    输入: "abca"
    输出: True
    解释: 你可以删除c字符。
  </pre>

  解题： 时间复杂度O(n) n为数组长度
*/
package questionBank

import ()

func validPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] == s[j] {
			i++
			j--
			continue
		} else {
			flag1 := true
			flag2 := true

			for left, right := i+1, j; left < right; left, right = left+1, right-1 {
				if s[left] != s[right] {
					flag1 = false
					break
				}
			}

			for left, right := i, j-1; left < right; left, right = left+1, right-1 {
				if s[left] != s[right] {
					flag2 = false
					break
				}
			}

			return flag1 || flag2
		}
	}

	return true
}
