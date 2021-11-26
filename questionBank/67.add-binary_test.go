package questionBank

import "testing"

func TestAddBinary(t *testing.T) {
	a := "11"
	b := "1"
	res := addBinary(a, b)
	if res != "100" {
		t.Error("二进制求和算法测试1错误", res)
	}

	a = "1010"
	b = "1011"
	res = addBinary(a, b)
	if res != "10101" {
		t.Error("二进制求和算法测试2错误", res)
	}
}
