/**
  题目地址：https://leetcode-cn.com/problems/climbing-stairs/

  假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
  每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
  注意：给定 n 是一个正整数。

  <pre>
    输入： 2
    输出： 2
    解释： 有两种方法可以爬到楼顶。
      1.  1 阶 + 1 阶
      2.  2 阶
  </pre>

  解题： climbStairs(n) 到n层时的方法数量
	1、到n-1的方法数为 climbStairs(n-1)
	2、到n-2的方法数为 climbStairs(n-2)
    那么 climbStairs(n) = climbStairs(n-1) + climbStairs(n-2)
*/
package questionBank

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	//n-2的方法数量
	left := 1
	//n-1的方法数量
	right := 2
	//这里n最小是2
	res := 2
	for i := 3; i <= n; i++ {
		res = left + right
		//重新给前两个楼梯的方法数赋值
		left, right = right, res
	}

	return res
}

func ClimbStairs(n int) int {
	return climbStairs(n)
}