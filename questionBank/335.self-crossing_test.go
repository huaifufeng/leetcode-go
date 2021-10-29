package questionBank

import "testing"

func TestIsSelfCrossing(t *testing.T) {
	distance := []int{2, 1, 1, 2}
	res := isSelfCrossing(distance)
	if !res {
		t.Error("路径交叉算法测试1错误")
	}

	distance = []int{1, 2, 3, 4}
	res = isSelfCrossing(distance)
	if res {
		t.Error("路径交叉算法测试2错误")
	}

	distance = []int{1, 1, 1, 1}
	res = isSelfCrossing(distance)
	if !res {
		t.Error("路径交叉算法测试3错误")
	}

	distance = []int{1, 1, 2, 1, 1}
	res = isSelfCrossing(distance)
	if !res {
		t.Error("路径交叉算法测试4错误")
	}
}
