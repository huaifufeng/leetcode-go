package questionBank

import (
	"testing"
)

func TestStrStr(t *testing.T) {
	res := StrStr("hellp", "ll")
	if res != 2 {
		t.Log(res)
		t.Error("字符串查找失败")
	}
}