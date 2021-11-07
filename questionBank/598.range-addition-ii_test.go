package questionBank

import "testing"

func TestMaxCount(t *testing.T) {
	m := 3
	n := 3
	ops := [][]int{
		{2, 2},
		{3, 3},
	}

	res := maxCount(m, n, ops)
	if res != 4 {
		t.Error("范围求和 II算法测试1错误", res)
	}
}
