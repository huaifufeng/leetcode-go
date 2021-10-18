/**
  题目地址：https://leetcode-cn.com/problems/number-complement/

  解题：
  1、依次除以2直到结果小于2，同时计算数值。 O(logn)
*/
package questionBank

func findComplement(num int) int {
	res := 0
	index := 1
	for num > 0 {
		left := num % 2
		num = num / 2
		if left == 0 {
			res = res + index
		}

		index = index << 1
	}

	return res
}
