//题目地址：https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/5/strings/39/
package 字符串

import (
	"strconv"
)

func countAndSay(n int) string {
	str := ""

	for i := 0; i < n; i++ {
		if i == 0 {
			str = "1"
			continue
		}
		str = getStrNew(str)
	}

	return str
}

func getStrNew(str string) string {
	var currentByte byte
	num := 0
	strNew := ""

	for i := 0; i < len(str); i++ {
		if currentByte == 0 {
			currentByte = str[i]
			num++
			continue
		}

		if currentByte == str[i] {
			num++
		} else {
			strNew = strNew + strconv.Itoa(num) + string([]byte{currentByte})
			currentByte = str[i]
			num = 1
		}
	}

	if currentByte != 0 {
		strNew = strNew + strconv.Itoa(num) + string([]byte{currentByte})
	}

	return strNew
}

