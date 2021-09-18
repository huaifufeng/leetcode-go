package questionBank

import "testing"

func TestCanWinNim(t *testing.T) {
	n := 4
	res := canWinNim(n)
	if res {
		t.Error("Nim 游戏算法测试1错误")
	}

	n = 1
	res = canWinNim(n)
	if !res {
		t.Error("Nim 游戏算法测试1错误")
	}

	n = 2
	res = canWinNim(n)
	if !res {
		t.Error("Nim 游戏算法测试1错误")
	}
}
