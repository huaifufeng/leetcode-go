package questionBank

import "testing"

func TestDistributeCandies(t *testing.T) {
	candyTye := []int{1, 1, 2, 2, 3, 3}
	res := distributeCandies(candyTye)
	if res != 3 {
		t.Error("分糖果算法测试1错误")
	}

	candyTye = []int{1, 1, 2, 3}
	res = distributeCandies(candyTye)
	if res != 2 {
		t.Error("分糖果算法测试2错误")
	}

	candyTye = []int{1, 1, 1, 1}
	res = distributeCandies(candyTye)
	if res != 1 {
		t.Error("分糖果算法测试3错误")
	}
}
