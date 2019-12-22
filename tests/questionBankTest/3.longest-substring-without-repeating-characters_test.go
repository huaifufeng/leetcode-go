package questionBankTest

import (
	"leetcode-go/questionBank"
	"testing"
)

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	value := questionBank.LengthOfLongestSubstring("abba")
	if value != 2{
		t.Log(value)
		t.Error("方法验证失败")
	}
}