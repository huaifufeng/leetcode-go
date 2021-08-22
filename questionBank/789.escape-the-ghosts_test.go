package questionBank

import "testing"

func TestEscapeGhosts(t *testing.T) {
	ghosts := [][]int{
		[]int{1, 0},
		[]int{0, 3},
	}
	target := []int{0, 1}
	res := escapeGhosts(ghosts, target)
	if res != true {
		t.Error("逃脱阻碍着算法1失败！")
	}

	ghosts = [][]int{
		[]int{1, 0},
	}
	target = []int{2, 0}
	res = escapeGhosts(ghosts, target)
	if res != false {
		t.Error("逃脱阻碍着算法2失败！")
	}

	ghosts = [][]int{
		[]int{2, 0},
	}
	target = []int{1, 0}
	res = escapeGhosts(ghosts, target)
	if res != false {
		t.Error("逃脱阻碍着算法3失败！")
	}

	ghosts = [][]int{
		[]int{5, 0},
		[]int{-10, -2},
		[]int{0, -5},
		[]int{-2, -2},
		[]int{-7, 1},
	}
	target = []int{7, 7}
	res = escapeGhosts(ghosts, target)
	if res != false {
		t.Error("逃脱阻碍着算法4失败！")
	}

	ghosts = [][]int{
		[]int{-1, 0},
		[]int{0, 1},
		[]int{-1, 0},
		[]int{0, 1},
		[]int{-1, 0},
	}
	target = []int{0, 0}
	res = escapeGhosts(ghosts, target)
	if res != true {
		t.Error("逃脱阻碍着算法5失败！")
	}
}
