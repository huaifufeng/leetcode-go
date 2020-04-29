/**
  题目地址：https://leetcode-cn.com/problems/longest-common-prefix/

  编写一个函数来查找字符串数组中的最长公共前缀。
  如果不存在公共前缀，返回空字符串 ""。
  <pre>
     输入: ["flower","flow","flight"]
     输出: "fl"
  </pre>

  <pre>
     输入: ["dog","racecar","car"]
     输出: ""
     解释: 输入不存在公共前缀。
  </pre>

  说明：
     所有输入只包含小写字母 a-z 。

  解题： 时间复杂度O(n)
	1.比较字符串中最小字符串长度的字符，看看是不是每个字符串都是一样的，一样就继续下去，直到最短的字符串匹配完毕；
*/
package questionBank

func longestCommonPrefix(strs []string) string {
	strLength := len(strs)
	if strLength == 0 {
		return ""
	}

	if strLength == 1 {
		return strs[0]
	}

	length := len(strs[0])
	for i := 0; i < length; i++ {
		cur := strs[0][i]
		for j := 0; j < strLength; j++ {
			if i == len(strs[j]) || strs[j][i] != cur {
				return strs[0][0:i]
			}
		}
	}

	return strs[0]
}
