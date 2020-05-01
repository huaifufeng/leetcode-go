package questionBank

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	res := removeElement(nums, 1)
	if res != 7 {
		t.Error("移除重复元素方法验证失败")
	}
}
