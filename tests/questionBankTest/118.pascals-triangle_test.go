package questionBankTest

import (
	"leetcode-go/questionBank"
	"testing"
)

func TestGenerate(t *testing.T) {
	res := questionBank.Generate(1)
	if res[0][0] != 1 {
		t.Error("方法错误")
	}
}