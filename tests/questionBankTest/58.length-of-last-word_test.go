package questionBankTest

import (
	"leetcode-go/questionBank"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	res := questionBank.LengthOfLastWord("Hello World")
	if res != 5 {
		t.Log(res)
		t.Error("最后单词长度算法错误")
	}
}