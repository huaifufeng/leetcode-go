package questionBank

import (
	"testing"
)

//测试方法
func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	targets := TwoSum(nums, target)
	//返回值长度必须为2
	if len(targets) != 2 {
		t.Error("返回值长度不正确")
	}

	//返回值必须是指定的值
	if targets[0] != 0 || targets[1] != 1 {
		t.Error("返回值不正确")
	}
}
