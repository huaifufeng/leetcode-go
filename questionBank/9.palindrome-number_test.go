package questionBank

import (
	"testing"
)

func TestPalindromeNumber(t *testing.T) {
	reverseInt := isPalindrome(1)
	if reverseInt != true {
		t.Log(reverseInt)
		t.Error("判断回文数组方法失败")
	}
}

func TestPalindromeNumber2(t *testing.T) {
	reverseInt := isPalindrome2(1)
	if reverseInt != true {
		t.Log(reverseInt)
		t.Error("判断回文数组方法2失败")
	}
}
