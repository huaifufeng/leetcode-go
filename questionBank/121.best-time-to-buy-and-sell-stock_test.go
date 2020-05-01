package questionBank

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	res := maxProfit([]int{7, 6, 4, 3, 1})
	if res != 0 {
		t.Error("股票最佳收益方法失败")
	}
}

func TestMaxProfit2(t *testing.T) {
	res := maxProfit2([]int{7, 6, 4, 3, 1})
	if res != 0 {
		t.Error("股票最佳收益2方法失败")
	}
}
