package questionBank

import "testing"

func TestRunningSum(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	res := runningSum(nums)

	if res[3] != 10 {
		t.Error("一维数组的动态和算法测试1错误！", res)
	}

	nums = []int{1, 1, 1, 1, 1}
	res = runningSum(nums)

	if res[4] != 5 {
		t.Error("一维数组的动态和算法测试2错误！", res)
	}

	nums = []int{3, 1, 2, 10, 1}
	res = runningSum(nums)

	if res[4] != 17 {
		t.Error("一维数组的动态和算法测试3错误！", res)
	}
}
