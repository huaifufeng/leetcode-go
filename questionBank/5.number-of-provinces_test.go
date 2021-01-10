package questionBank

import (
	"testing"
)

func TestFindCircleNum(t *testing.T) {
	isConnected := [][]int{[]int{1, 1, 0}, []int{1, 1, 0}, []int{0, 0, 1}}
	if findCircleNum(isConnected) != 2 {
		t.Error("省份数量计算错误1")
	}

	isConnected = [][]int{[]int{1, 0, 0}, []int{0, 1, 0}, []int{0, 0, 1}}
	if findCircleNum(isConnected) != 3 {
		t.Error("省份数量计算错误2")
	}

	isConnected = [][]int{[]int{1, 0, 0, 1}, []int{0, 1, 1, 0}, []int{0, 1, 1, 1}, []int{1, 0, 1, 1}}
	if findCircleNum(isConnected) != 1 {
		t.Error("省份数量计算错误3")
	}
}
