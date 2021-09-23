/**
  题目地址：https://leetcode-cn.com/problems/power-of-three/

  解题：
	1.3的幂次方，一直除以3，直到结果是1 O(logn)
    2.计算许可范围内的最大3次方值，n如果是这个值得可除数，就是3的幂次方 O(1)
*/
package questionBank

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}

	for n%3 == 0 {
		n = n / 3
	}

	return n == 1
}
