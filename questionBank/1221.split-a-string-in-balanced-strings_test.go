package questionBank

import "testing"

func TestBalancedStringSplit(t *testing.T) {
	s := "RLRRLLRLRL"
	res := balancedStringSplit(s)
	if res != 4 {
		t.Error("分割平衡字符串算法测试1错误", res)
	}

	s = "RLLLLRRRLR"
	res = balancedStringSplit(s)
	if res != 3 {
		t.Error("分割平衡字符串算法测试2错误", res)
	}

	s = "LLLLRRRR"
	res = balancedStringSplit(s)
	if res != 1 {
		t.Error("分割平衡字符串算法测试3错误", res)
	}

	s = "RLRRRLLRLL"
	res = balancedStringSplit(s)
	if res != 2 {
		t.Error("分割平衡字符串算法测试4错误", res)
	}
}
