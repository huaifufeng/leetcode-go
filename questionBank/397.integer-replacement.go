/**
  题目地址：https://leetcode-cn.com/problems/nteger-replacement/

  解题：
	1、递归法，如果是偶数加一，除2；奇数的话就分别取加1和减一中的小值。

*/
package questionBank

func integerReplacement(n int) int {
	if n == 1 {
		return 0
	}

	if n%2 == 0 {
		return 1 + integerReplacement(n/2)
	}

	return 2 + min(integerReplacement(n/2), integerReplacement(n/2+1))
}
