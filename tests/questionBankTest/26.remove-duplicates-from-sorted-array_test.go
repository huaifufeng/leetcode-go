package questionBankTest

import (
	"leetcode-go/questionBank"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	res := questionBank.RemoveDuplicates([]int{0,0,1,1,1,2,2,3,3,4})
	if res != 5 {
		t.Log(res)
		t.Error("方法验证失败")
	}
}