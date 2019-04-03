//题目地址：https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/5/strings/37/
package 字符串

import (
	"math"
	"strings"
)

func myAtoi(str string) int {
	str = strings.Trim(str, " ")
	length := len(str)

	if length == 0 || (!(str[0] >= '0' && str[0] <= '9') && str[0] != '+' && str[0] != '-') {
		return 0
	}

	str = strings.ToLower(str)
	strNew := []byte(str)

	ret := 0
	flag := 1
	i := 0

	if str[0] == '-' {
		flag = -1
		i++
	}

	if str[0] == '+' {
		i++
	}

	max := math.MaxInt32 / 10
	min := math.MinInt32 / 10

	for ; i<length; i++ {
		if !(str[i] >= '0' && str[i] <= '9') {
			break
		}

		newVal := int(strNew[i] - '0')

		if ret > max || (ret == max && newVal > 7) {
			return math.MaxInt32
		}

		if ret < min || (ret == min && newVal > 8) {
			return math.MinInt32
		}

		ret = ret * 10 + flag * newVal
	}

	return ret
}