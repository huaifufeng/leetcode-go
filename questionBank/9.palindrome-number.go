/**
  题目地址：https://leetcode-cn.com/problems/palindrome-number/

  判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

  <pre>
     输入: 121
     输出: true
  </pre>

  <pre>
     输入: -121
     输出: false
     解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
  </pre>

  进阶：
     你能不将整数转为字符串来解决这个问题吗？

  解题：
   按照规则，首先如果为负数，肯定是不满足条件的，首先判断一下，然后在按照下面之一的算法进行计算
   1、将数字转换为字符串，然后依次比较第一个字符和最后一个字符，直到字符串长度小于等于2的时候，说明比较完毕了
   2、进阶中要求不转换为字符串，那就这里用数字进行计算
*/
package questionBank

import (
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	intStr := strconv.Itoa(x)
	lenNow := len(intStr)
	//当字符串长度小于等于1的时候，说明里面已经处理完毕了 要不正确，要不失败
	for lenNow > 1 {
		lenNow = len(intStr)
		if intStr[0] == intStr[lenNow - 1] {
			if lenNow > 2 {
				intStr = intStr[1:lenNow-1]
			} else {
				return true
			}
		} else {
			return false
		}
	}

	return true
}

