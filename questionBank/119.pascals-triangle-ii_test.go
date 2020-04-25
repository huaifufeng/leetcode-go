package questionBank

import (
	"testing"
)

func TestGetRow(t *testing.T) {
	res := GetRow(3)
	if res[2] != 3 {
		t.Error("方法错误")
	}
}
