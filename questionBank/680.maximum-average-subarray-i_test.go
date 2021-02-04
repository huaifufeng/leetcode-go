package questionBank

import "testing"

func TestFindMaxAverage(t *testing.T) {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	res := findMaxAverage(nums, k)

	if res != 12.75 {
		t.Error("子字符串最大平均数I错误！")
	}
}
