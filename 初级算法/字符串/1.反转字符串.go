//题目地址：https://leetcode-cn.com/explore/interview/card/top-interview-questions-easy/5/strings/32/
package 字符串

func reverseString(s []byte) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		s[i], s[length-1-i] = s[length-1-i], s[i]
	}
}
