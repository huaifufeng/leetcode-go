package questionBank

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	res := searchInsert([]int{1, 3, 5, 6}, 0)
	if res != 0 {
		t.Log(res)
		t.Error("搜索插入方法错误")
	}
}
