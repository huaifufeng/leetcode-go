package questionBank

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	res := Generate(1)
	if res[0][0] != 1 {
		t.Error("方法错误")
	}
}