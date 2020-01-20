package questionBankTest

import (
	"leetcode-go/questionBank"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	res := questionBank.MaxProfit([]int{})
	if res != 0 {
		t.Log(res)
		t.Error("搜索插入方法错误")
	}
}