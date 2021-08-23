package questionBank

import "testing"

func TestGetMaximumGenerated(t *testing.T) {
	n := 7
	res := getMaximumGenerated(n)
	if res != 3 {
		t.Error("获取生成数组中的最大值测试1错误！", res)
	}

	n = 2
	res = getMaximumGenerated(n)
	if res != 1 {
		t.Error("获取生成数组中的最大值测试2错误！", res)
	}

	n = 3
	res = getMaximumGenerated(n)
	if res != 2 {
		t.Error("获取生成数组中的最大值测试3错误！", res)
	}
}
