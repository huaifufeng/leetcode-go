package questionBank

import "testing"

func TestOriginalDigits(t *testing.T) {
	s := "owoztneoer"
	res := originalDigits(s)
	if res != "012" {
		t.Error("从英文中重建数字算法测试1错误", res)
	}

	s = "fviefuro"
	res = originalDigits(s)
	if res != "45" {
		t.Error("从英文中重建数字算法测试2错误", res)
	}
}
