package questionBank

import "testing"

func TestIsPowerOfThree(t *testing.T) {
	n := 27
	res := isPowerOfThree(n)
	if !res {
		t.Error("3的幂算法测试1错误")
	}

	n = 0
	res = isPowerOfThree(n)
	if res {
		t.Error("3的幂算法测试2错误")
	}

	n = 9
	res = isPowerOfThree(n)
	if !res {
		t.Error("3的幂算法测3错误")
	}

	n = 45
	res = isPowerOfThree(n)
	if res {
		t.Error("3的幂算法测试4错误")
	}
}
