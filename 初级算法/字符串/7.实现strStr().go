//题目地址：https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/5/strings/38/
package 字符串

func strStr(haystack string, needle string) int {
	hayLen := len(haystack)
	needleLen := len(needle)
	ret := -1

	if needleLen == 0 {
		return 0
	}

	if hayLen < needleLen {
		return ret
	}

	for i := 0; i < hayLen; i++ {
		if hayLen-i < needleLen {
			return -1
		}

		j := 0
		if haystack[i] == needle[j] {
			flag := 1
			for ; j < needleLen; j++ {
				if haystack[i+j] != needle[j] {
					flag = 0
					break
				}
			}

			if flag == 1 {
				return i
			}
		}
	}

	return ret
}
