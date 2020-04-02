package questionBank

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{0,0,1,1,1,2,2,3,3,4}
	res := RemoveElement(nums, 1)
	if res != 7 {
		t.Log(res)
		t.Error("方法验证失败")
	}
}