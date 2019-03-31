//题目地址：https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/1/array/27/
package 数组

func plusOne(digits []int) []int {
	length := len(digits)
	if length == 0 {
		return []int{1}
	}

	upVal := 1
	for i := length - 1; i >= 0; i-- {
		upVal+= digits[i]
		digits[i] = upVal %  10
		upVal /= 10
	}

	if upVal > 0 {
		digits = append([]int{upVal}, digits...)
	}

	return digits
}