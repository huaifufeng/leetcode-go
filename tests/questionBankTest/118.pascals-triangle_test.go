package questionBankTest

import (
	"leetcode-go/questionBank"
	"testing"
)

func TestGenerate(t *testing.T) {
	res := questionBank.Generate(0)
	if res[0][0] != 2 {
		t.Error("方法错误")
	}
}