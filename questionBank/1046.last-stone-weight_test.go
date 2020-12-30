package questionBank

import "testing"

func TestLastStoneWeight(t *testing.T) {
	stones := []int{2, 7, 4, 1, 8, 1}

	res := lastStoneWeight(stones)
	if res != 1 {
		t.Error("最后石头重量出错")
	}
}
