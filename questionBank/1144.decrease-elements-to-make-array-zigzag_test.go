package questionBank

import "testing"

func TestMovesToMakeZigzag(t *testing.T) {
	nums := []int{2, 1, 2}
	res := movesToMakeZigzag(nums)

	t.Error(res)
	if res != 0 {
		t.Error("递减元素使数组呈锯齿状判断错误！")
	}
}
