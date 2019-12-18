package questionBankTest

import (
	"leetcode-go/questionBank"
	"testing"
)

func TestReverseInteger(t *testing.T)  {
	reverseInt := questionBank.Reverse7(1534236469)
	if reverseInt != 10 {
		t.Log(reverseInt)
		t.Error("正数反转失败")
	}
}