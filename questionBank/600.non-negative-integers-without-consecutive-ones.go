/**
  题目地址：https://leetcode-cn.com/problems/non-negative-integers-without-consecutive-ones/

  解题：
  1、暴力循环，计算每个值得合法情况，只要当前值除2的值不合法，或者 当前取余2为，除2的值得尾数是1  这两种情况不合法， 时间复杂度O(n)
  2、
*/
package questionBank

func findIntegers(n int) int {
	numMap := make(map[int]uint8)
	numMap[0] = 2 // 公式 2 * x + left    x是当前正确还是错误 正确1 错误 0  left 当前的余数 0或1
	total := 1
	for i := 1; i <= n; i++ {
		lastNum := i >> 1
		left := uint8(i % 2)
		delete(numMap, lastNum-1)
		//如果上个数值是非法的，当前数值非法
		//如果当前尾数是1， 上个数的尾数也是1，不合法
		if (numMap[lastNum] == 0 || numMap[lastNum] == 1) || (numMap[lastNum] == 3 && left == 1) {
			numMap[i] = left
			continue
		}

		numMap[i] = 2 + left
		total++
	}

	return total
}
