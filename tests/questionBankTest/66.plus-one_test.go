package questionBankTest

import (
	"leetcode-go/questionBank"
	"testing"
)

func TestPlusOne(t *testing.T) {
	res := questionBank.PlusOne([]int{9})
	if res[2] != 4 {
		t.Log(res)
		t.Error("数组整数加1出错")
	}
}