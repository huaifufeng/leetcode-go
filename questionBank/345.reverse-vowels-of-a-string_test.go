package questionBank

import "testing"

func TestReverseVowels(t *testing.T) {
	str := "hello"
	res := reverseVowels(str)
	if res != "holle" {
		t.Error("反转元音字符方法测试1错误", res)
	}

	str = "leetcode"
	res = reverseVowels(str)
	if res != "leotcede" {
		t.Error("反转元音字符方法测试2错误", res)
	}

	str = "aA"
	res = reverseVowels(str)
	if res != "Aa" {
		t.Error("反转元音字符方法测试3错误", res)
	}
}
