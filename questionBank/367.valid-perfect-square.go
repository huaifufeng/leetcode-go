/**
 题目地址：https://leetcode-cn.com/problems/valid-perfect-square/

解题： 时间复杂度O(n) n为数组长度

*/
package questionBank

func isPerfectSquare(num int) bool {
	sqrter := 1

	for {
		multi := sqrter * sqrter
		if multi == num {
			return true
		}

		if multi > num {
			return false
		}

		if num-sqrter*sqrter < 2*sqrter+1 {
			return false
		}
		sqrter++
	}
}
