/**
  题目地址：https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/

  写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项（即 F(N)）。斐波那契数列的定义如下：
    F(0) = 0,   F(1) = 1
    F(N) = F(N - 1) + F(N - 2), 其中 N > 1.

  斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。
  答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

  <pre>
     输入：n = 2
     输出：1
  </pre>

  <pre>
     输入：n = 5
     输出：5
  </pre>

  PS:
    0 <= n <= 100


  解题：
    1、将递归算法转换为循环算法，计算的数量会少很多，时间复杂度 O(n)
*/
package offer

func fib(n int) int {
	if n < 2 {
		return n
	}

	left := 1
	right := 0
	cur := 0
	modNum := 1000000007
	for i := 2; i <= n; i++ {
		cur = (left + right) % modNum
		right = left
		left = cur
	}

	return cur
}
