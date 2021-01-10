package questionBank

import (
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	nums := []int{0, 1, 2, 4, 5, 7}
	res := summaryRanges(nums)
	if !(len(res) == 3 && res[0] == "0->2" && res[1] == "4->5" && res[2] == "7") {
		t.Error("汇总区间算法错误1")
	}

	nums = []int{0, 2, 3, 4, 6, 8, 9}
	res = summaryRanges(nums)
	if !(len(res) == 4 && res[0] == "0" && res[1] == "2->4" && res[2] == "6" && res[3] == "8->9") {
		t.Error("汇总区间算法错误2")
	}

	nums = []int{}
	res = summaryRanges(nums)
	if len(res) != 0 {
		t.Error("汇总区间算法错误3")
	}

	nums = []int{-1}
	res = summaryRanges(nums)
	if !(len(res) == 1 && res[0] == "-1") {
		t.Error("汇总区间算法错误4")
	}

	nums = []int{0}
	res = summaryRanges(nums)
	if !(len(res) == 1 && res[0] == "0") {
		t.Error("汇总区间算法错误5")
	}
}
