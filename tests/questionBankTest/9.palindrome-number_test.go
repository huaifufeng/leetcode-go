package questionBankTest

import (
	"leetcode-go/questionBank"
	"testing"
)

func TestPalindromeNumber(t *testing.T)  {
	reverseInt := questionBank.PalindromeNumber(1)
	if reverseInt != true {
		t.Log(reverseInt)
		t.Error("正数反转失败")
	}
}