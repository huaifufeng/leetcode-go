package questionBankTest

import (
	"leetcode-go/questionBank"
	"testing"
)

func TestValidParentheses(t *testing.T) {
	value := questionBank.ValidParentheses("{}")
	if value != true{
		t.Log(value)
		t.Error("方法验证失败")
	}
}