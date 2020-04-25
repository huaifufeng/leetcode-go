/**
  题目地址：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

  给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

  <pre>
     输入: "abcabcbb"
     输出: 3
     解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
  </pre>

  <pre>
     输入: "pwwkew"
     输出: 3
     解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
          请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
  </pre>


  解题：
	1、暴力统计，循环整个字符串，从中找到不含有重复字符的字符串 时间复杂度O(n²)
	2、滑动窗口 两个变量 i j, i-j之间的字符是不重复的字符串，当遇到重复的字符串时，就比较i和重复字符串所在的位置，
       获取最大的值作为i的值，继续进行下去，直到比较完毕 时间复杂度O(n)
*/
package questionBank

func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	length := len(s)
	subStrMap := make(map[byte]int, length)
	i := 0
	j := 0

	//循环整个数组
	for i < length && j < length {
		//如果字符存在，说明要进行字符串位置i的一个互换操作
		if lastIndex, ok := subStrMap[s[j]]; ok {
			//只有在字符在i之后的位置时进行i值得重置操作，因为位于i之前的话，说明已经不再统计范围内部了，不需要进行处理了
			if lastIndex > i {
				currLength := j - i
				if currLength > maxLength {
					maxLength = currLength
				}

				i = lastIndex
			}
		}

		//更新字符所在的位置，永远是最新所在的位置
		subStrMap[s[j]] = j + 1
		j++
	}

	currLength := j - i
	if currLength > maxLength {
		maxLength = currLength
	}

	return maxLength
}

func LengthOfLongestSubstring(s string) int {
	return lengthOfLongestSubstring(s)
}
