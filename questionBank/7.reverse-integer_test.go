package questionBank

import (
	"testing"
)

func TestReverseInteger(t *testing.T) {
	reverseInt := Reverse7(1534236469)
	if reverseInt != 0 {
		t.Log(reverseInt)
		t.Error("正数反转失败")
	}
}