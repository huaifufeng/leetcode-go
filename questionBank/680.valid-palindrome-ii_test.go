package questionBank

import "testing"

func TestValidPalindrome(t *testing.T) {
	str := "abca"
	res := validPalindrome(str)

	if res != true {
		t.Error("回文字符串判断错误！")
	}
}
