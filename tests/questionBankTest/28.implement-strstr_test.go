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

*/
package questionBankTest

import (
	"leetcode-go/questionBank"
	"testing"
)

func TestStrStr(t *testing.T) {
	res := questionBank.StrStr("hellp", "ll")
	if res != 2 {
		t.Log(res)
		t.Error("字符串查找失败")
	}
}