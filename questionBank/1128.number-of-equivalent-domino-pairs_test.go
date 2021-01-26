package questionBank

import (
	"fmt"
	"testing"
)

func TestNumEquivDominoPairs(t *testing.T) {
	dominoes := [][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}
	res := numEquivDominoPairs(dominoes)

	if res != 1 {
		fmt.Println(res)
		t.Error("等价塔罗牌数量计算错误1！")
	}

	dominoes = [][]int{{1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}}
	res = numEquivDominoPairs(dominoes)

	if res != 3 {
		fmt.Println(res)
		t.Error("等价塔罗牌数量计算错误2！")
	}
}
