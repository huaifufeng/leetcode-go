/**
  题目地址：https://leetcode-cn.com/problems/implement-strstr/

  实现 strStr() 函数。
  给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

  <pre>
    输入: haystack = "hello", needle = "ll"
    输出: 2
  </pre>

  说明：
  当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
  对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。

  解题： 时间复杂度O(n) n为数组长度
	循环整个数组，比较当前元素和指定val，不相等的话，是有效的元素，将其移动到pos位置
*/
package questionBank

func strStr(haystack string, needle string) int {
	if len(needle) == 0  {
		return 0
	}

	strLen := len(needle)
	hayLen := len(haystack)
	for i := 0; i < hayLen; i++  {
		if haystack[i] == needle[0] {
			endPos := i+strLen
			if endPos > hayLen {
				return -1
			}
			substr := string(haystack[i:i+strLen])
			if substr == needle {
				return i
			} else if len(substr) < strLen {
				return -1
			}
		}
	}

	return -1
}

func StrStr(haystack string, needle string) int {
	return strStr(haystack, needle)
}