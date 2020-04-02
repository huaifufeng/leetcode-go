package questionBank

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	res := MaxProfit([]int{7,6,4,3,1})
	if res != 0 {
		t.Log(res)
		t.Error("股票最佳收益方法失败")
	}
}