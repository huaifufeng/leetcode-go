package questionBank

import (
	"testing"
)

func TestValidParentheses(t *testing.T) {
	value := ValidParentheses("{}")
	if value != true{
		t.Log(value)
		t.Error("方法验证失败")
	}
}