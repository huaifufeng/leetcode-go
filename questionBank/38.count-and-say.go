/**
  题目地址：https://leetcode-cn.com/problems/count-and-say/

  解题：
  1、依次处理每个数字对应的数字字符串，
*/
package questionBank

import (
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	str := "1"
	for n > 1 {
		builder := strings.Builder{}
		lastStr := ' '
		num := 0
		for _, val := range str {
			if val != lastStr {
				if num > 0 && lastStr != ' ' {
					nStr := strconv.Itoa(num)
					builder.WriteString(nStr)
					builder.WriteRune(lastStr)
				}
				lastStr = val
				num = 1
			} else {
				num++
			}
		}
		if num > 0 && lastStr != ' ' {
			nStr := strconv.Itoa(num)
			builder.WriteString(nStr)
			builder.WriteRune(lastStr)
		}

		str = builder.String()
		n--
	}

	return str
}
