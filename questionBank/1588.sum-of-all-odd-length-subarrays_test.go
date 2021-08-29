package questionBank

import "testing"

func TestSumOddLengthSubarrays(t *testing.T) {
	arr := []int{1, 4, 2, 5, 3}
	res := sumOddLengthSubarrays(arr)
	if res != 58 {
		t.Error("所有奇数长度子数组的和算法测试1错误！", res)
	}

	arr = []int{1, 2}
	res = sumOddLengthSubarrays(arr)
	if res != 3 {
		t.Error("所有奇数长度子数组的和算法测试2错误！", res)
	}

	arr = []int{10, 11, 12}
	res = sumOddLengthSubarrays(arr)
	if res != 66 {
		t.Error("所有奇数长度子数组的和算法测试3错误！", res)
	}
}
