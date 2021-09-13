package questionBank

import (
	"testing"
)

func TestNumberOfBoomerangs(t *testing.T) {

	points := [][]int{
		[]int{0, 0},
		[]int{1, 0},
		[]int{2, 0},
	}
	res := numberOfBoomerangs(points)
	if res != 2 {
		t.Error("回旋镖的数量算法测试1错误", res)
	}

	points = [][]int{
		[]int{1, 1},
		[]int{2, 2},
		[]int{3, 3},
	}
	res = numberOfBoomerangs(points)
	if res != 2 {
		t.Error("回旋镖的数量算法测试2错误", res)
	}

	points = [][]int{
		[]int{1, 1},
	}
	res = numberOfBoomerangs(points)
	if res != 0 {
		t.Error("回旋镖的数量算法测试3错误", res)
	}

	points = make([][]int, 0)
	for i := 0; i < 50; i++ {
		for j := 0; j < 10; j++ {
			points = append(points, []int{i, j})
		}
	}

	res = numberOfBoomerangs1(points)
	if res != 476784 {
		t.Error("回旋镖的数量算法1测试1错误", res)
	}
}
