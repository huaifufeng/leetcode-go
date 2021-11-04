package questionBank

import "testing"

func TestIsPerfectSquare(t *testing.T) {
	num := 16
	res := isPerfectSquare(num)
	if !res {
		t.Error("有效的完全平方数算法测试1错误")
	}

	num = 14
	res = isPerfectSquare(num)
	if res {
		t.Error("有效的完全平方数算法测试2错误")
	}

	num = 1
	res = isPerfectSquare(num)
	if !res {
		t.Error("有效的完全平方数算法测试3错误")
	}
}
