package questionBankTest

import (
	"leetcode-go/questionBank"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	res := questionBank.ClimbStairs(10)
	if res != 5 {
		t.Log(res)
		t.Error("爬楼算法错误")
	}
}