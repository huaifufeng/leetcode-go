/**
  题目地址：https://leetcode-cn.com/problems/reverse-integer/

  给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

  <pre>
     输入: 123
     输出: 321
  </pre>

  <pre>
     输入: -123
     输出: -321
  </pre>

  注意：
    假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

  解题： 时间复杂度 O(lgn)
   这里将数字反转，就是按照十进制，依次获取最后一位的值 pop，然后将当前值 reverseInt * 10 + pop。
   这里主要是要注意整数溢出的边界条件，也就是：
     1、正数时
       reverseInt > MAX / 10, pop > 0 报错
       reverseInt = MAX / 10, pop > 7 报错
     2、负数
       reverseInt < MIN / 10, pop > 0 报错
       reverseInt = MIN / 10, pop < -8 报错
*/
package questionBank

import (
	"math"
)

func reverse(x int) int {
	reverseInt := 0
	sign := 1
	if x < 0 {
		sign = -1
	}

	positiveNum := math.MaxInt32 / 10
	negativeNum := math.MinInt32 / 10
	pop := 0

	for{
		pop = x % 10
		x = x / 10
		reverseInt = reverseInt * 10 + pop
		if sign == 1 {
			if reverseInt > positiveNum && pop > 0 {
				return 0
			}

			if reverseInt == positiveNum && pop > 7 {
				return 0
			}
		} else {
			if reverseInt < negativeNum && pop < 0 {
				return 0
			}

			if reverseInt == negativeNum && pop < -8 {
				return 0
			}
		}

		if x > -10 && x < 10 {
			break
		}
	}

	if x != 0 {
		reverseInt = reverseInt * 10 + x
	}

	return reverseInt
}

func Reverse7(x int) int {
	return reverse(x)
}