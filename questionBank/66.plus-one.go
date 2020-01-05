/**
  题目地址：https://leetcode-cn.com/problems/plus-one/

  给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
  最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
  你可以假设除了整数 0 之外，这个整数不会以零开头。

  <pre>
    输入: [1,2,3]
    输出: [1,2,4]
    解释: 输入数组表示数字 123。
  </pre>

  解题：
	1、从后向前对每个元素进行处理，加一之后，如果10的余数为0，说明上一位还要加1进位，否则不等于0就说明不需要进位直接返回；
    如果到最后都没有等于0的，说明数组向前多加了一位，所以最后单独处理一下
*/
package questionBank

func plusOne(digits []int) []int {
	digitLength := len(digits)

	for i := digitLength - 1; i >= 0; i-- {
		digits[i]++
		digits[i] %= 10
		if digits[i] != 0 {
			return digits
		}
	}

	digits = make([]int, digitLength + 1)
	digits[0] = 1

	return digits
}

func PlusOne(digits []int) []int {
	return plusOne(digits)
}