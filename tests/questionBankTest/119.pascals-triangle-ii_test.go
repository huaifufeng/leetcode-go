package questionBankTest

import (
	"leetcode-go/questionBank"
	"testing"
)

func TestGetRow(t *testing.T) {
	res := questionBank.GetRow(3)
	if res[2] != 3 {
		t.Error("方法错误")
	}
}