//题目地址：https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/5/strings/33/
package 字符串

import "math"

func reverse(x int) int {
	max := math.MaxInt32 / 10
	min := int(math.Abs(math.MinInt32 / 10))
	reverser := 0

	for x != 0{
		num := x % 10

		if reverser > max || (reverser == max && num > 7) {
			 return 0
		}

		if reverser < min || (reverser == min && num > 8) {
			return 0
		}


		reverser = reverser * 10 + num
		x /= 10
	}

	return reverser
}