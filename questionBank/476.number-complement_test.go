package questionBank

import "testing"

func TestFindComplement(t *testing.T) {
	num := 1
	res := findComplement(num)
	if res != 0 {
		t.Error("数字的补数算法测试1错误")
	}

	num = 5
	res = findComplement(num)
	if res != 2 {
		t.Error("数字的补数算法测试2错误")
	}
}
