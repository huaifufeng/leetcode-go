package questionBank

import (
	"testing"
)

func TestStrStr(t *testing.T) {
	res := strStr("hellp", "ll")
	if res != 2 {
		t.Error("字符串查找失败")
	}
}
