package questionBank

import "testing"

func TestFindPoisonedDuration(t *testing.T) {
	timeSeries := []int{1, 4}
	duration := 2
	res := findPoisonedDuration(timeSeries, duration)
	if res != 4 {
		t.Error("提莫攻击算法测试1错误", res)
	}

	timeSeries = []int{1, 2}
	duration = 2
	res = findPoisonedDuration(timeSeries, duration)
	if res != 3 {
		t.Error("提莫攻击算法测试2错误", res)
	}

	timeSeries = []int{1}
	duration = 2
	res = findPoisonedDuration(timeSeries, duration)
	if res != 2 {
		t.Error("提莫攻击算法测试3错误", res)
	}
}
