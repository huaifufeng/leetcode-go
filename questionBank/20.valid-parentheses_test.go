package questionBank

import (
	"testing"
)

func TestValidParentheses(t *testing.T) {
	value := isValid("{}")
	if value != true {
		t.Error("括号匹配方法验证失败")
	}
}
