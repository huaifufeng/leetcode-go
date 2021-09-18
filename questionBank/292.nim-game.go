/**
  题目地址：https://leetcode-cn.com/problems/nim-game/


  解题：
	1、变量二维数组，分别建立行，列，方块的元素映射，只要存在大于1的值，说明重复非法， O(1) = O(81)
*/
package questionBank

func canWinNim(n int) bool {
	reminder := n % 4
	if reminder == 0 {
		return false
	}

	return true
}
