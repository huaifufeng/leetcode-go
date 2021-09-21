/**
  题目地址：https://leetcode-cn.com/problems/nim-game/


  解题：
	1、归纳总结，当是4的倍数的时候，一定会输，否则就赢 O(1)
*/
package questionBank

func canWinNim(n int) bool {
	reminder := n % 4
	if reminder == 0 {
		return false
	}

	return true
}
