package questionBank

import "testing"

func TestAllPathsSourceTarget(t *testing.T) {
	graph := [][]int{
		[]int{1, 2},
		[]int{3},
		[]int{3},
		[]int{},
	}
	res := allPathsSourceTarget(graph)
	if len(res) != 2 {
		t.Error("所有可能的路径算法测试1错误！", res)
	}

	graph = [][]int{
		[]int{4, 3, 1},
		[]int{3, 2, 4},
		[]int{3},
		[]int{4},
		[]int{},
	}
	res = allPathsSourceTarget(graph)
	if len(res) != 5 {
		t.Error("所有可能的路径算法测试2错误！", res)
	}

	graph = [][]int{
		[]int{1},
		[]int{},
	}
	res = allPathsSourceTarget(graph)
	if len(res) != 1 {
		t.Error("所有可能的路径算法测试3错误！", res)
	}

	graph = [][]int{
		[]int{1, 2, 3},
		[]int{2},
		[]int{3},
		[]int{},
	}
	res = allPathsSourceTarget(graph)
	if len(res) != 3 {
		t.Error("所有可能的路径算法测试4错误！", res)
	}

	graph = [][]int{
		[]int{1, 3},
		[]int{2},
		[]int{3},
		[]int{},
	}
	res = allPathsSourceTarget(graph)
	if len(res) != 2 {
		t.Error("所有可能的路径算法测试5错误！", res)
	}

	graph = [][]int{
		[]int{4, 3, 1},
		[]int{3, 2, 4},
		[]int{},
		[]int{4},
		[]int{},
	}
	res = allPathsSourceTarget(graph)
	if len(res) != 2 {
		t.Error("所有可能的路径算法测试6错误！", res)
	}
}
