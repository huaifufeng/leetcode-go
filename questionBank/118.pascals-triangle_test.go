package questionBank

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	res := generate(1)
	if res[0][0] != 1 {
		t.Error("杨辉三角方法错误")
	}
}
