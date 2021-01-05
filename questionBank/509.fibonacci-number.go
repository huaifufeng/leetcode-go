/**
  题目地址：https://leetcode-cn.com/problems/fibonacci-number/

  斐波那契数，通常用 F(n) 表示，形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：
  	F(0) = 0，F(1) = 1
	F(n) = F(n - 1) + F(n - 2)，其中 n > 1
  给你 n ，请计算 F(n) 。

  <pre>
    输入：2
    输出：1
    解释：F(2) = F(1) + F(0) = 1 + 0 = 1
  </pre>

  解题：
    1、特殊处理0和1之后，其他的元素都可以使用递归进行处理
*/
package questionBank

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}