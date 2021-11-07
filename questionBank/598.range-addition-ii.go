/**
  题目地址：https://leetcode-cn.com/problems/range-addition-ii/

  解题：
  1、
*/
package questionBank

func maxCount(m int, n int, ops [][]int) int {
	line := m
	high := n

	for _, op := range ops {
		if op[0] < line {
			line = op[0]
		}

		if op[1] < high {
			high = op[1]
		}
	}

	return line * high
}
